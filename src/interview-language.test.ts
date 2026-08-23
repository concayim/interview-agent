import { describe, expect, it } from 'vitest'
import { buildCompletionSpeech, buildFollowUpSpeech, buildInterviewOpeningSpeech, buildNextQuestionSpeech } from './interview-language'

describe('English interview speech copy', () => {
  it('builds an English opening without Chinese fragments', () => {
    const text = buildInterviewOpeningSpeech('en-US', 'Alex', '不用追求背诵感。', 'Explain ownership in Rust.')

    expect(text).toBe('Hello, Alex. Welcome to your technical interview. First question: Explain ownership in Rust.')
    expect(text).not.toMatch(/[\u3400-\u9fff]/)
  })

  it('defensively replaces a Chinese default candidate name', () => {
    expect(buildInterviewOpeningSpeech('en-US', '候选人', '中文开场'))
      .toBe('Hello, Candidate. Welcome to your technical interview.')
  })

  it('builds English transition and completion speech', () => {
    expect(buildNextQuestionSpeech('en-US', 'Good coverage.', 'How does borrowing work?'))
      .toBe('Good coverage. Next question: How does borrowing work?')
    expect(buildCompletionSpeech('en-US', 'Strong answer.'))
      .toBe('Strong answer. The interview is complete, and your detailed review is ready.')
  })

  it('announces a follow-up without calling it the next question', () => {
    expect(buildFollowUpSpeech('zh-CN', '还需补充调度细节。', 'G、M、P 如何协作？'))
      .toBe('还需补充调度细节。继续追问：G、M、P 如何协作？')
    expect(buildFollowUpSpeech('en-US', 'Please add the scheduler details.', 'How do G, M, and P cooperate?'))
      .toBe('Please add the scheduler details. Follow-up: How do G, M, and P cooperate?')
  })
})
