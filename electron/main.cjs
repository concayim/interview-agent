const { app, BrowserWindow, ipcMain, shell } = require('electron')
const { spawn } = require('node:child_process')
const { randomBytes } = require('node:crypto')
const net = require('node:net')
const path = require('node:path')

const TOKEN = randomBytes(24).toString('hex')
let backend
let port

function backendCommand(backendPort) {
  if (!app.isPackaged) {
    return { command: 'go', args: ['run', './cmd/server', '-port', String(backendPort)], cwd: path.join(__dirname, '..') }
  }
  const binary = process.platform === 'win32' ? 'interview-agent-server.exe' : 'interview-agent-server'
  return { command: path.join(process.resourcesPath, 'bin', binary), args: ['-port', String(backendPort)], cwd: process.resourcesPath }
}

function findAvailablePort() {
  return new Promise((resolve, reject) => {
    const probe = net.createServer()
    probe.unref()
    probe.once('error', reject)
    probe.listen(0, '127.0.0.1', () => {
      const address = probe.address()
      const availablePort = typeof address === 'object' && address ? address.port : 0
      probe.close((error) => error ? reject(error) : resolve(availablePort))
    })
  })
}

function startBackend(backendPort) {
  const target = backendCommand(backendPort)
  backend = spawn(target.command, target.args, {
    cwd: target.cwd,
    env: {
      ...process.env,
      INTERVIEW_AGENT_TOKEN: TOKEN,
      INTERVIEW_AGENT_DATA_DIR: path.join(app.getPath('userData'), 'data'),
    },
    stdio: ['ignore', 'pipe', 'pipe'],
  })
  backend.stdout.on('data', (chunk) => console.log(`[backend] ${chunk.toString().trimEnd()}`))
  backend.stderr.on('data', (chunk) => console.error(`[backend] ${chunk.toString().trimEnd()}`))
  backend.on('exit', (code) => { if (code && !app.isQuitting) console.error(`Backend exited with code ${code}`) })
}

async function waitForBackend(backendPort) {
  const deadline = Date.now() + 30000
  while (Date.now() < deadline) {
    try {
      const response = await fetch(`http://127.0.0.1:${backendPort}/api/v1/health`)
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
    port = await findAvailablePort()
    ipcMain.on('interview-agent:bootstrap', (event) => {
      event.returnValue = {
        apiBaseUrl: `http://127.0.0.1:${port}/api/v1`,
        apiToken: TOKEN,
        platform: process.platform,
      }
    })
    startBackend(port)
    await waitForBackend(port)
    createWindow()
  } catch (error) {
    console.error(error)
    app.quit()
  }
  app.on('activate', () => { if (BrowserWindow.getAllWindows().length === 0) createWindow() })
})

app.on('before-quit', () => {
  app.isQuitting = true
  if (backend && !backend.killed) backend.kill('SIGTERM')
})

app.on('window-all-closed', () => { if (process.platform !== 'darwin') app.quit() })
