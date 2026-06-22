const { app, BrowserWindow, shell } = require('electron')
const { spawn } = require('node:child_process')
const { randomBytes } = require('node:crypto')
const path = require('node:path')

const PORT = 46831
const TOKEN = randomBytes(24).toString('hex')
let backend

function backendCommand() {
  if (!app.isPackaged) {
    return { command: 'go', args: ['run', './cmd/server', '-port', String(PORT)], cwd: path.join(__dirname, '..') }
  }
  const binary = process.platform === 'win32' ? 'interview-agent-server.exe' : 'interview-agent-server'
  return { command: path.join(process.resourcesPath, 'bin', binary), args: ['-port', String(PORT)], cwd: process.resourcesPath }
}

function startBackend() {
  const target = backendCommand()
  process.env.INTERVIEW_AGENT_TOKEN = TOKEN
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

async function waitForBackend() {
  const deadline = Date.now() + 30000
  while (Date.now() < deadline) {
    try {
      const response = await fetch(`http://127.0.0.1:${PORT}/api/v1/health`)
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
  startBackend()
  try {
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
  if (backend && !backend.killed) backend.kill('SIGTERM')
})

app.on('window-all-closed', () => { if (process.platform !== 'darwin') app.quit() })
