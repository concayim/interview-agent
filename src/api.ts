import type { AnswerResult, Difficulty, KnowledgeBase, LearningResult, ModelConfig, QAItem, Report, Resume, Session, SkillCatalog } from './types'

function runtimeConfig() {
  const config = window.interviewAgent?.getRuntimeConfig?.() ?? window.interviewAgent
  return {
    baseUrl: (config?.apiBaseUrl || 'http://127.0.0.1:46831/api/v1').replace(/\/$/, ''),
    token: config?.apiToken || '',
  }
}

async function request<T>(path: string, init: RequestInit = {}): Promise<T> {
  const { baseUrl, token } = runtimeConfig()
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
  skills: () => request<SkillCatalog>('/skills'),
  startInterview: (input: { candidateName: string; resumeId?: string; domainSkillId: string; interviewerSkillId: string; includeFoundation: boolean; videoEnabled: boolean; speechLanguage: string; difficulty: Difficulty; questionCount: number }) =>
    request<Session>('/interviews', { method: 'POST', body: JSON.stringify(input) }),
  answer: (sessionId: string, answer: string, elapsedSeconds: number) =>
    request<AnswerResult>(`/interviews/${sessionId}/answers`, { method: 'POST', body: JSON.stringify({ answer, elapsedSeconds }) }),
  transcribeSpeech: (audio: Blob, language: string) => {
    const body = new FormData()
    body.append('audio', audio, `answer-${Date.now()}.webm`)
    body.append('language', language.startsWith('en') ? 'en' : 'zh')
    return request<{ text: string }>('/speech/transcriptions', { method: 'POST', body })
  },
  report: (sessionId: string) => request<Report>(`/interviews/${sessionId}/report`),
  getModelConfig: () => request<ModelConfig>('/config/model'),
  saveModelConfig: (input: { apiKey: string; baseUrl: string; model: string; speechModel: string; enabled: boolean; clearApiKey?: boolean }) =>
    request<ModelConfig>('/config/model', { method: 'PUT', body: JSON.stringify(input) }),
  testModelConfig: () => request<{ ok: boolean; message: string }>('/config/model/test', { method: 'POST', body: '{}' }),
  knowledgeBases: () => request<{ bases: KnowledgeBase[] }>('/knowledge/bases'),
  searchKnowledge: (baseId: string, query = '', limit = 30) => request<{ items: QAItem[] }>(`/knowledge/bases/${encodeURIComponent(baseId)}/qa?query=${encodeURIComponent(query)}&limit=${limit}`),
  addKnowledge: (baseId: string, input: { question: string; answer: string; keyPoints?: string[]; tags?: string[]; difficulty?: string }) => request<QAItem>(`/knowledge/bases/${encodeURIComponent(baseId)}/qa`, { method: 'POST', body: JSON.stringify(input) }),
  learningResources: (domainSkillId = 'all', kind = 'all', selectedOnly = false) => request<LearningResult>(`/learning/resources?domainSkillId=${encodeURIComponent(domainSkillId)}&kind=${encodeURIComponent(kind)}&selectedOnly=${selectedOnly}`),
  refreshLearning: () => request<LearningResult>('/learning/refresh', { method: 'POST', body: '{}' }),
  selectLearning: (id: string, selected: boolean) => request<{ id: string; selected: boolean }>(`/learning/resources/${encodeURIComponent(id)}/selection`, { method: 'PUT', body: JSON.stringify({ selected }) }),
}
