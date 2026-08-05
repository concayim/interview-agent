export type AvatarMode = 'idle' | 'listening' | 'thinking' | 'speaking'

type AvatarActivity = {
  speechMode: 'idle' | 'connecting' | 'recording' | 'finalizing'
  speaking: boolean
  submitting: boolean
}

export function deriveAvatarMode({ speechMode, speaking, submitting }: AvatarActivity): AvatarMode {
  if (speechMode !== 'idle') return 'listening'
  if (speaking) return 'speaking'
  if (submitting) return 'thinking'
  return 'idle'
}
