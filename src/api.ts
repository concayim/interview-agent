import type { AnswerResult, Difficulty, Language, ModelConfig, Report, Resume, Session } from './types'

const baseUrl = window.interviewAgent?.apiBaseUrl ?? 'http://127.0.0.1:46831/api/v1'
const token = window.interviewAgent?.apiToken ?? ''

async function request<T>(path: string, init: RequestInit = {}): Promise<T> {
  const headers = new Headers(init.headers)
  if (!(init.body instanceof FormData)) headers.set('Content-Type', 'application/json')
  if (token) headers.set('X-Interview-Agent-Token', token)
  const response = await fetch(`${baseUrl}${path}`, { ...init, headers })
  const payload = await response.json().catch(() => ({}))
  if (!response.ok) throw new Error(payload.error || `请求失败（${response.status}）`)
  return payload as T
}

export const api = {
  health: () => request<{ status: string; modelConfigured: boolean }>('/health'),
  uploadResume: (file: File) => {
    const body = new FormData()
    body.append('file', file)
    return request<Resume>('/resumes', { method: 'POST', body })
  },
  startInterview: (input: { candidateName: string; resumeId?: string; language: Language; difficulty: Difficulty; questionCount: number }) =>
    request<Session>('/interviews', { method: 'POST', body: JSON.stringify(input) }),
  answer: (sessionId: string, answer: string, elapsedSeconds: number) =>
    request<AnswerResult>(`/interviews/${sessionId}/answers`, { method: 'POST', body: JSON.stringify({ answer, elapsedSeconds }) }),
  report: (sessionId: string) => request<Report>(`/interviews/${sessionId}/report`),
  getModelConfig: () => request<ModelConfig>('/config/model'),
  saveModelConfig: (input: { apiKey: string; baseUrl: string; model: string; enabled: boolean; clearApiKey?: boolean }) =>
    request<ModelConfig>('/config/model', { method: 'PUT', body: JSON.stringify(input) }),
  testModelConfig: () => request<{ ok: boolean; message: string }>('/config/model/test', { method: 'POST', body: '{}' }),
}
