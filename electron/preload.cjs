const { contextBridge } = require('electron')

function runtimeArg(name) {
  const prefix = `--${name}=`
  return process.argv.find((value) => value.startsWith(prefix))?.slice(prefix.length)
}

contextBridge.exposeInMainWorld('interviewAgent', {
  apiBaseUrl: runtimeArg('interview-agent-api-base-url') || process.env.INTERVIEW_AGENT_API_BASE_URL || 'http://127.0.0.1:46831/api/v1',
  apiToken: runtimeArg('interview-agent-token') || process.env.INTERVIEW_AGENT_TOKEN || '',
  platform: process.platform,
})
