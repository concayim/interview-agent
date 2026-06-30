const { contextBridge, ipcRenderer } = require('electron')

function runtimeArg(name) {
  const prefix = `--${name}=`
  return process.argv.find((value) => value.startsWith(prefix))?.slice(prefix.length)
}

function readRuntimeConfig() {
  try {
    const config = ipcRenderer.sendSync('interview-agent-runtime-config')
    if (config?.apiBaseUrl && config?.apiToken) return config
  } catch { /* fall back to launch arguments */ }
  return {
    apiBaseUrl: runtimeArg('interview-agent-api-base-url') || process.env.INTERVIEW_AGENT_API_BASE_URL || 'http://127.0.0.1:46831/api/v1',
    apiToken: runtimeArg('interview-agent-token') || process.env.INTERVIEW_AGENT_TOKEN || '',
  }
}

const runtimeConfig = readRuntimeConfig()

contextBridge.exposeInMainWorld('interviewAgent', {
  ...runtimeConfig,
  getRuntimeConfig: readRuntimeConfig,
  platform: process.platform,
})
