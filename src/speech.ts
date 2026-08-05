export type RealtimeSpeechSession = {
  stop: () => void
}

type RuntimeConfig = { baseUrl: string; token: string }

function runtimeConfig(): RuntimeConfig {
  const config = window.interviewAgent?.getRuntimeConfig?.() ?? window.interviewAgent
  return {
    baseUrl: (config?.apiBaseUrl || 'http://127.0.0.1:46831/api/v1').replace(/\/$/, ''),
    token: config?.apiToken || '',
  }
}

export async function startRealtimeSpeech(
  language: string,
  onText: (text: string, final: boolean) => void,
  onError: (error: Error) => void,
  signal?: AbortSignal,
): Promise<RealtimeSpeechSession> {
  const stream = await navigator.mediaDevices.getUserMedia({
    audio: { channelCount: 1, echoCancellation: true, noiseSuppression: true, autoGainControl: true },
  })
  if (signal?.aborted) {
    stream.getTracks().forEach((track) => track.stop())
    throw new DOMException('实时语音连接已取消', 'AbortError')
  }
  const audioContext = new AudioContext()
  const source = audioContext.createMediaStreamSource(stream)
  const processor = audioContext.createScriptProcessor(8192, 1, 1)
  const silentOutput = audioContext.createGain()
  silentOutput.gain.value = 0

  const { baseUrl, token } = runtimeConfig()
  const socketURL = new URL(`${baseUrl.replace(/^http/, 'ws')}/speech/realtime`)
  socketURL.searchParams.set('language', language)
  if (token) socketURL.searchParams.set('token', token)
  const socket = new WebSocket(socketURL)
  socket.binaryType = 'arraybuffer'

  let active = true
  let recording = false
  let captureCleaned = false

  const stopCapture = () => {
    if (captureCleaned) return
    captureCleaned = true
    recording = false
    processor.disconnect()
    source.disconnect()
    silentOutput.disconnect()
    stream.getTracks().forEach((track) => track.stop())
    void audioContext.close()
  }

  const closeWithError = (error: Error) => {
    if (!active) return
    active = false
    stopCapture()
    socket.close()
    onError(error)
  }

  processor.onaudioprocess = (event) => {
    if (!recording || socket.readyState !== WebSocket.OPEN) return
    const pcm = floatToPCM16(resampleTo16k(event.inputBuffer.getChannelData(0), audioContext.sampleRate))
    if (pcm.byteLength) socket.send(pcm)
  }

  await new Promise<void>((resolve, reject) => {
    let ready = false
    const rejectAndClean = (error: Error) => {
      if (!active) return
      active = false
      stopCapture()
      socket.close()
      reject(error)
    }
    const abort = () => rejectAndClean(new DOMException('实时语音连接已取消', 'AbortError'))
    signal?.addEventListener('abort', abort, { once: true })
    socket.onerror = () => rejectAndClean(new Error('无法连接本地实时语音服务'))
    socket.onclose = () => {
      if (!active) return
      if (ready) closeWithError(new Error('实时语音连接已关闭'))
      else rejectAndClean(new Error('无法连接本地实时语音服务'))
    }
    socket.onmessage = (event) => {
      let message: { type?: string; text?: string; error?: string }
      try {
        message = JSON.parse(String(event.data))
      } catch {
        rejectAndClean(new Error('实时语音服务返回了无法解析的数据'))
        return
      }
      if (message.type === 'error') {
        rejectAndClean(new Error(message.error || '实时语音识别失败'))
        return
      }
      if (message.type !== 'ready') return
      ready = true
      signal?.removeEventListener('abort', abort)
      socket.onerror = () => closeWithError(new Error('实时语音连接异常中断'))
      socket.onmessage = (nextEvent) => {
        try {
          const next = JSON.parse(String(nextEvent.data)) as { type?: string; text?: string; error?: string }
          if (next.type === 'error') { closeWithError(new Error(next.error || '实时语音识别失败')); return }
          if ((next.type === 'partial' || next.type === 'final') && typeof next.text === 'string') {
            onText(next.text, next.type === 'final')
            if (next.type === 'final') {
              active = false
              stopCapture()
              socket.close()
            }
          }
        } catch {
          closeWithError(new Error('实时语音服务返回了无法解析的数据'))
        }
      }
      recording = true
      source.connect(processor)
      processor.connect(silentOutput)
      silentOutput.connect(audioContext.destination)
      resolve()
    }
  })

  return {
    stop: () => {
      if (!active) return
      stopCapture()
      if (socket.readyState === WebSocket.OPEN) socket.send(JSON.stringify({ type: 'stop' }))
      else socket.close()
    },
  }
}

function resampleTo16k(input: Float32Array, inputRate: number) {
  if (inputRate === 16000) return input
  const ratio = inputRate / 16000
  const output = new Float32Array(Math.floor(input.length / ratio))
  for (let index = 0; index < output.length; index += 1) {
    const position = index * ratio
    const left = Math.floor(position)
    const right = Math.min(left + 1, input.length - 1)
    const weight = position - left
    output[index] = input[left] * (1 - weight) + input[right] * weight
  }
  return output
}

function floatToPCM16(input: Float32Array) {
  const output = new ArrayBuffer(input.length * 2)
  const view = new DataView(output)
  input.forEach((sample, index) => {
    const normalized = Math.max(-1, Math.min(1, sample))
    view.setInt16(index * 2, normalized < 0 ? normalized * 0x8000 : normalized * 0x7fff, true)
  })
  return output
}
