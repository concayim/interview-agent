const { app, BrowserWindow } = require('electron')
const crypto = require('node:crypto')
const net = require('node:net')
const path = require('node:path')

const WEBSOCKET_GUID = '258EAFA5-E914-47DA-95CA-C5AB0DC85B11'

async function run() {
  const sockets = new Set()
  const server = net.createServer((socket) => {
    sockets.add(socket)
    socket.once('close', () => sockets.delete(socket))
    socket.once('data', (data) => {
      const key = data.toString().match(/^Sec-WebSocket-Key:\s*(.+)$/im)?.[1]?.trim()
      if (!key) return socket.destroy(new Error('Missing Sec-WebSocket-Key'))
      const accept = crypto.createHash('sha1').update(key + WEBSOCKET_GUID).digest('base64')
      socket.write([
        'HTTP/1.1 101 Switching Protocols',
        'Upgrade: websocket',
        'Connection: Upgrade',
        `Sec-WebSocket-Accept: ${accept}`,
        '',
        '',
      ].join('\r\n'))
    })
  })
  await new Promise((resolve, reject) => {
    server.once('error', reject)
    server.listen(0, '127.0.0.1', resolve)
  })

  const port = server.address().port
  let window
  try {
    window = new BrowserWindow({ show: false, webPreferences: { sandbox: true } })
    await window.loadFile(path.join(__dirname, '..', 'dist', 'index.html'))
    await window.webContents.executeJavaScript(`new Promise((resolve, reject) => {
      const socket = new WebSocket('ws://127.0.0.1:${port}/csp-test')
      const timer = setTimeout(() => reject(new Error('WebSocket open timeout')), 2000)
      socket.onopen = () => { clearTimeout(timer); socket.close(); resolve(true) }
      socket.onerror = () => { clearTimeout(timer); reject(new Error('WebSocket connection failed')) }
    })`)
  } catch (error) {
    throw new Error(`Local WebSocket port ${port} was not reachable from Electron: ${error.message}`)
  } finally {
    window?.destroy()
    for (const socket of sockets) socket.destroy()
    await new Promise((resolve) => server.close(resolve))
  }
}

app.whenReady()
  .then(run)
  .then(() => app.quit())
  .catch((error) => {
    console.error(error.message)
    app.exit(1)
  })
