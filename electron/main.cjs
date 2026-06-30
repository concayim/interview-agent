const { app, BrowserWindow, ipcMain, shell } = require('electron')
const { spawn, spawnSync } = require('node:child_process')
const { randomBytes } = require('node:crypto')
const fs = require('node:fs')
const net = require('node:net')
const os = require('node:os')
const path = require('node:path')

const DEFAULT_PORT = 46831
const TOKEN = randomBytes(24).toString('hex')
let port = DEFAULT_PORT
let backend
let runtimeConfig = {
  apiBaseUrl: `http://127.0.0.1:${DEFAULT_PORT}/api/v1`,
  apiToken: TOKEN,
}

ipcMain.on('interview-agent-runtime-config', (event) => {
  event.returnValue = runtimeConfig
})

function checkPort(candidate) {
  return new Promise((resolve) => {
    const server = net.createServer()
    server.once('error', () => resolve(false))
    server.once('listening', () => server.close(() => resolve(true)))
    server.listen(candidate, '127.0.0.1')
  })
}

async function findAvailablePort(preferred) {
  for (let candidate = preferred; candidate < preferred + 40; candidate += 1) {
    if (await checkPort(candidate)) return candidate
  }
  return await new Promise((resolve, reject) => {
    const server = net.createServer()
    server.once('error', reject)
    server.listen(0, '127.0.0.1', () => {
      const address = server.address()
      server.close(() => resolve(address.port))
    })
  })
}

function backendCommand() {
  if (!app.isPackaged) {
    const cwd = path.join(__dirname, '..')
    const binary = path.join(os.tmpdir(), `interview-agent-server-dev-${process.pid}${process.platform === 'win32' ? '.exe' : ''}`)
    const build = spawnSync('go', ['build', '-o', binary, './cmd/server'], { cwd, stdio: 'inherit' })
    if (build.status !== 0) throw new Error('后端服务构建失败')
    return { command: binary, args: ['-port', String(port)], cwd, cleanup: () => fs.rmSync(binary, { force: true }) }
  }
  const binary = process.platform === 'win32' ? 'interview-agent-server.exe' : 'interview-agent-server'
  return { command: path.join(process.resourcesPath, 'bin', binary), args: ['-port', String(port)], cwd: process.resourcesPath }
}

async function startBackend() {
  port = await findAvailablePort(DEFAULT_PORT)
  const apiBaseUrl = `http://127.0.0.1:${port}/api/v1`
  runtimeConfig = { apiBaseUrl, apiToken: TOKEN }
  const target = backendCommand()
  process.env.INTERVIEW_AGENT_TOKEN = TOKEN
  process.env.INTERVIEW_AGENT_API_BASE_URL = apiBaseUrl
  backend = spawn(target.command, target.args, {
    cwd: target.cwd,
    env: {
      ...process.env,
      INTERVIEW_AGENT_TOKEN: TOKEN,
      INTERVIEW_AGENT_API_BASE_URL: apiBaseUrl,
      INTERVIEW_AGENT_DATA_DIR: path.join(app.getPath('userData'), 'data'),
    },
    stdio: ['ignore', 'pipe', 'pipe'],
    detached: process.platform !== 'win32',
  })
  backend.cleanup = target.cleanup
  backend.stdout.on('data', (chunk) => console.log(`[backend] ${chunk.toString().trimEnd()}`))
  backend.stderr.on('data', (chunk) => console.error(`[backend] ${chunk.toString().trimEnd()}`))
  backend.on('exit', (code) => {
    backend.cleanup?.()
    if (code && !app.isQuitting) console.error(`Backend exited with code ${code}`)
  })
}

async function waitForBackend() {
  const deadline = Date.now() + 30000
  while (Date.now() < deadline) {
    try {
      const response = await fetch(`http://127.0.0.1:${port}/api/v1/health`)
      if (response.ok) return
    } catch { /* backend is still starting */ }
    await new Promise((resolve) => setTimeout(resolve, 250))
  }
  throw new Error('后端服务启动超时')
}

function createWindow() {
  const window = new BrowserWindow({
    width: 1420,
    height: 920,
    minWidth: 1080,
    minHeight: 720,
    show: false,
    backgroundColor: '#0b1220',
    titleBarStyle: process.platform === 'darwin' ? 'hiddenInset' : 'default',
    webPreferences: {
      preload: path.join(__dirname, 'preload.cjs'),
      contextIsolation: true,
      nodeIntegration: false,
      sandbox: true,
      additionalArguments: [
        `--interview-agent-api-base-url=${runtimeConfig.apiBaseUrl}`,
        `--interview-agent-token=${runtimeConfig.apiToken}`,
      ],
    },
  })
  window.webContents.setWindowOpenHandler(({ url }) => {
    if (url.startsWith('https://')) shell.openExternal(url)
    return { action: 'deny' }
  })
  window.once('ready-to-show', () => window.show())
  if (process.env.VITE_DEV_SERVER_URL) window.loadURL(process.env.VITE_DEV_SERVER_URL)
  else window.loadFile(path.join(__dirname, '..', 'dist', 'index.html'))
}

app.whenReady().then(async () => {
  try {
    await startBackend()
    await waitForBackend()
    createWindow()
  } catch (error) {
    console.error(error)
    app.quit()
  }
  app.on('activate', () => { if (BrowserWindow.getAllWindows().length === 0) createWindow() })
})

app.on('before-quit', () => {
  app.isQuitting = true
  if (backend && !backend.killed) {
    if (process.platform !== 'win32') process.kill(-backend.pid, 'SIGTERM')
    else backend.kill('SIGTERM')
  }
})

app.on('window-all-closed', () => { if (process.platform !== 'darwin') app.quit() })
