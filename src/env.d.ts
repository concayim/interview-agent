/// <reference types="vite/client" />

interface Window {
  interviewAgent?: {
    apiBaseUrl: string
    apiToken: string
    getRuntimeConfig?: () => { apiBaseUrl: string; apiToken: string }
    platform: string
  }
}
