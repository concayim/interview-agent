export function isEnglishInterview(language: string) {
  return language.toLowerCase().startsWith('en')
}

export function buildInterviewOpeningSpeech(language: string, candidateName: string, chineseOpening: string, firstQuestion?: string) {
  if (isEnglishInterview(language)) {
    const englishName = /[\u3400-\u9fff]/.test(candidateName) ? 'Candidate' : candidateName
    return [`Hello, ${englishName}. Welcome to your technical interview.`, firstQuestion ? `First question: ${firstQuestion}` : ''].filter(Boolean).join(' ')
  }
  return `你好，${candidateName}。${chineseOpening}${firstQuestion ? `。第一题，${firstQuestion}` : ''}`
}

export function buildNextQuestionSpeech(language: string, summary: string, nextQuestion: string) {
  if (isEnglishInterview(language)) return `${summary} Next question: ${nextQuestion}`
  return `${summary}。下一题，${nextQuestion}`
}

export function buildFollowUpSpeech(language: string, summary: string, followUpQuestion: string) {
  if (isEnglishInterview(language)) return `${summary} Follow-up: ${followUpQuestion}`
  return `${summary}继续追问：${followUpQuestion}`
}

export function buildCompletionSpeech(language: string, summary: string) {
  if (isEnglishInterview(language)) return `${summary} The interview is complete, and your detailed review is ready.`
  return `${summary}。本场面试已经完成，完整复盘已为你整理好。`
}
