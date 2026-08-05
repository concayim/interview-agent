import { describe, expect, it } from 'vitest'
import { deriveAvatarMode } from './avatar-state'

describe('deriveAvatarMode', () => {
  it('shows listening while realtime speech capture is active', () => {
    expect(deriveAvatarMode({ speechMode: 'recording', speaking: false, submitting: false })).toBe('listening')
  })

  it('keeps listening above playback to avoid a speaking avatar during capture', () => {
    expect(deriveAvatarMode({ speechMode: 'recording', speaking: true, submitting: false })).toBe('listening')
  })

  it('shows speaking only while interviewer audio is playing', () => {
    expect(deriveAvatarMode({ speechMode: 'idle', speaking: true, submitting: false })).toBe('speaking')
  })

  it('shows thinking while an answer is being evaluated', () => {
    expect(deriveAvatarMode({ speechMode: 'idle', speaking: false, submitting: true })).toBe('thinking')
  })

  it('returns to idle when no realtime activity is present', () => {
    expect(deriveAvatarMode({ speechMode: 'idle', speaking: false, submitting: false })).toBe('idle')
  })
})
