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
  language: Language
  difficulty: Difficulty
  prompt: string
  tags: string[]
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
  status: 'active' | 'completed'
  current: number
  total: number
  currentQuestion?: Question
  startedAt: string
}

export type AnswerRecord = {
  question: Question & { standardAnswer: string; keyPoints: string[] }
  answer: string
  elapsedSeconds: number
  evaluation: Evaluation
}

export type Report = {
  sessionId: string
  candidateName: string
  language: Language
  difficulty: Difficulty
  score: number
  answered: number
  durationSeconds: number
  highlights: string[]
  focusAreas: string[]
  answers: AnswerRecord[]
  startedAt: string
  completedAt: string
}

export type AnswerResult = {
  evaluation: Evaluation
  completed: boolean
  nextQuestion?: Question
  current: number
  total: number
  report?: Report
}

export type ModelConfig = {
  baseUrl: string
  model: string
  enabled: boolean
  hasApiKey: boolean
}

export type Language = 'golang' | 'java' | 'python' | 'cpp'
export type Difficulty = 'easy' | 'medium' | 'hard' | 'mixed'
