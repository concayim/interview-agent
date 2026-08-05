# Doubao Duplex Voice Interview Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or execute inline with strict TDD checkpoints.

**Goal:** Keep real-time Doubao ASR for candidate answers and add Doubao TTS playback for interviewer speech.

**Architecture:** The Go backend owns separate ASR and TTS credentials. A new HTTP synthesis endpoint calls Doubao's unidirectional streaming TTS API, decodes the NDJSON Base64 audio chunks, and returns the completed MP3 to React. The interview screen owns one cancellable audio player and automatically speaks opening, questions, and assistant replies.

**Tech Stack:** Go `net/http`, React/TypeScript, HTML Audio, Doubao Speech APIs.

## Global Constraints

- Do not expose provider credentials to the renderer.
- Preserve the existing real-time ASR path.
- TTS failures must not block text interviews.
- Work in the existing dirty workspace because the uncommitted ASR implementation is a required dependency.

---

### Task 1: TTS Provider Boundary

**Files:**
- Create: `internal/api/speech_tts.go`
- Create: `internal/api/speech_tts_test.go`

- [ ] Write failing tests for request JSON, old/new console authentication, NDJSON audio decoding, and upstream errors.
- [ ] Run focused tests and verify failure because the TTS boundary does not exist.
- [ ] Implement the minimal provider client and `POST /api/v1/speech/synthesis` handler.
- [ ] Run focused tests and verify they pass.

### Task 2: Independent TTS Configuration

**Files:**
- Modify: `internal/config/store.go`
- Modify: `internal/api/server.go`
- Modify: `src/types.ts`
- Modify: `src/api.ts`
- Modify: `src/App.tsx`

- [ ] Write failing config preservation tests.
- [ ] Add TTS API Key, App ID, Resource ID, speaker, and enabled fields without reusing ASR resource configuration.
- [ ] Add settings controls and preserve stored secrets on blank updates.
- [ ] Run Go and TypeScript checks.

### Task 3: Automatic Interview Playback

**Files:**
- Create: `src/tts.ts`
- Modify: `src/App.tsx`
- Modify: `src/styles.css`

- [ ] Implement a single cancellable player that fetches synthesized MP3.
- [ ] Speak opening/current question, assistant replies, and next questions once per content ID.
- [ ] Add mute/unmute and replay icon controls with tooltips.
- [ ] Ensure ASR capture stops TTS playback to prevent echo transcription.
- [ ] Build the renderer.

### Task 4: Documentation and Verification

**Files:**
- Modify: `README.md`
- Modify: `docs/api.md`
- Modify: `docs/requirements.md`
- Modify: `CHANGELOG.md`

- [ ] Document separate ASR/TTS credentials and endpoint behavior.
- [ ] Run `npm test` and `git diff --check`.
- [ ] Start the desktop development environment and verify health.
