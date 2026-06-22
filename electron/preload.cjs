const { contextBridge } = require('electron')

contextBridge.exposeInMainWorld('interviewAgent', {
  apiBaseUrl: 'http://127.0.0.1:46831/api/v1',
  apiToken: process.env.INTERVIEW_AGENT_TOKEN || '',
  platform: process.platform,
})
