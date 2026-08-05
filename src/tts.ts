import { api } from './api'

type VoicePlayerEvents = {
  onPlaybackChange?: (playing: boolean) => void
}

export class InterviewVoicePlayer {
  private audio?: HTMLAudioElement
  private abort?: AbortController
  private objectURL?: string

  constructor(private readonly events: VoicePlayerEvents = {}) {}

  async speak(text: string) {
    this.stop()
    const abort = new AbortController()
    this.abort = abort
    const blob = await api.synthesizeSpeech(text, abort.signal)
    if (abort.signal.aborted) return
    const objectURL = URL.createObjectURL(blob)
    this.objectURL = objectURL
    const audio = new Audio(objectURL)
    this.audio = audio
    const release = () => {
      this.events.onPlaybackChange?.(false)
      if (this.objectURL === objectURL) {
        URL.revokeObjectURL(objectURL)
        this.objectURL = undefined
      }
    }
    audio.addEventListener('ended', release, { once: true })
    audio.addEventListener('error', release, { once: true })
    this.events.onPlaybackChange?.(true)
    try { await audio.play() }
    catch (error) { release(); throw error }
  }

  stop() {
    this.abort?.abort()
    this.abort = undefined
    this.audio?.pause()
    this.audio = undefined
    this.events.onPlaybackChange?.(false)
    if (this.objectURL) {
      URL.revokeObjectURL(this.objectURL)
      this.objectURL = undefined
    }
  }
}
