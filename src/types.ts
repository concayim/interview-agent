export type Resume = {
  id: string
  fileName: string
  contentType: string
  preview: string
  keywords: string[]
  characters: number
  createdAt: string
}

export type Question = {
  id: string
  promptId: string
  stage: 'introduction' | 'technical' | 'follow_up'
  leadIn?: string
  language: Language
  difficulty: Difficulty
  prompt: string
  tags: string[]
  followUp: boolean
  round: number
  followUpTotal: number
}

export type Evaluation = {
  score: number
  summary: string
  strengths: string[]
  improvements: string[]
  source: 'local' | 'llm'
}

export type Session = {
  id: string
  candidateName: string
  language: Language
  difficulty: Difficulty
  industry: string
  domainSkillId: string
  domainSkillName: string
  interviewerSkillId: string
  interviewerName: string
  interviewerOpening: string
  includeFoundation: boolean
  status: 'active' | 'completed'
  phase: 'introduction' | 'technical' | 'completed'
  current: number
  total: number
  followUpRound: number
  followUpTotal: number
  currentQuestion?: Question
  startedAt: string
}

export type AnswerRecord = {
  question: Question & { standardAnswer: string; keyPoints: string[] }
  answer: string
  elapsedSeconds: number
  evaluation: Evaluation
  followUps: {
    round: number
    prompt: string
    answer: string
    elapsedSeconds: number
    evaluation: Evaluation
  }[]
  averageScore: number
}

export type Report = {
  sessionId: string
  candidateName: string
  language: Language
  difficulty: Difficulty
  domainSkillName: string
  interviewerName: string
  score: number
  answered: number
  durationSeconds: number
  highlights: string[]
  focusAreas: string[]
  introduction?: {
    prompt: string
    answer: string
    elapsedSeconds: number
  }
  answers: AnswerRecord[]
  startedAt: string
  completedAt: string
}

export type AnswerResult = {
  evaluation: Evaluation
  questionCompleted: boolean
  completed: boolean
  nextQuestion?: Question
  current: number
  total: number
  followUpRound: number
  followUpTotal: number
  phase: 'introduction' | 'technical' | 'completed'
  report?: Report
}

export type ModelConfig = {
  baseUrl: string
  model: string
  enabled: boolean
  hasApiKey: boolean
}

export type Skill = {
  id: string
  kind: 'interviewer' | 'domain'
  name: string
  shortLabel: string
  description: string
  accent: string
  avatar?: string
  openingLine?: string
  evaluationFocus?: string[]
  feedbackTone?: string
  followUpRounds?: number
  industry?: string
  industryName?: string
  domain?: string
  language?: string
  knowledgeBaseId?: string
  topics?: string[]
}

export type SkillCatalog = {
  interviewers: Skill[]
  domains: Skill[]
  industries: { id: string; name: string }[]
}

export type KnowledgeBase = {
  id: string
  name: string
  description: string
  industry: string
  industryName: string
  language: string
  accent: string
  topics: string[]
  itemCount: number
  issuedCount: number
}

export type QAItem = {
  id: string
  baseId: string
  question: string
  answer: string
  keyPoints: string[]
  tags: string[]
  language: string
  difficulty: string
  source: 'built-in' | 'interview' | 'manual'
  issuedCount: number
  createdAt: string
  updatedAt: string
}

export type LearningResource = {
  id: string
  title: string
  url: string
  kind: 'article' | 'video' | 'course' | 'docs'
  source: string
  authority: string
  summary: string
  domainSkillIds: string[]
  publishedAt?: string
  selected: boolean
  live: boolean
}

export type LearningResult = {
  resources: LearningResource[]
  refreshedAt?: string
  warnings?: string[]
}

export type Language = string
export type Difficulty = 'easy' | 'medium' | 'hard' | 'mixed'
