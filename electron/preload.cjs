const { contextBridge, ipcRenderer } = require('electron')

const bootstrap = ipcRenderer.sendSync('interview-agent:bootstrap')

contextBridge.exposeInMainWorld('interviewAgent', {
  apiBaseUrl: bootstrap.apiBaseUrl,
  apiToken: bootstrap.apiToken,
  platform: bootstrap.platform,
})
