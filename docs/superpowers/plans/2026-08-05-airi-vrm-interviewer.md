# AIRI-inspired VRM Interviewer Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Execute inline with strict TDD checkpoints. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add an animated AI interviewer avatar that participates in the existing real-time text, ASR, and TTS interview flow.

**Architecture:** A React-owned VRM stage renders a local, redistributable model with Three.js and `@pixiv/three-vrm`. A small pure state mapper translates interview, speech-recognition, and TTS lifecycle state into avatar modes; the existing Go question/answer service remains the only conversation authority.

**Tech Stack:** React 19, TypeScript, Three.js, `@pixiv/three-vrm`, Vitest, Electron.

## Global Constraints

- Keep provider credentials and conversation state in the existing boundaries.
- Do not embed AIRI in an iframe or introduce a second chat runtime.
- Bundle the VRM model locally so the desktop app works without a model CDN.
- TTS and avatar failures must not block text interviews.
- Preserve all existing dirty-worktree changes.

---

### Task 1: Avatar State Contract

**Files:**
- Create: `src/avatar-state.ts`
- Create: `src/avatar-state.test.ts`

**Interfaces:**
- Produces: `deriveAvatarMode({ speechMode, speaking, submitting }): AvatarMode`

- [x] Write failing tests for listening, speaking, thinking, and idle priority.
- [x] Run the focused test and verify failure because the mapper does not exist.
- [x] Implement the minimal pure mapper and verify the test passes.

### Task 2: VRM Stage and Voice Lifecycle

**Files:**
- Create: `src/AvatarStage.tsx`
- Modify: `src/tts.ts`
- Modify: `src/App.tsx`
- Modify: `src/styles.css`
- Modify: `package.json`
- Create: `public/models/interviewer.vrm`
- Create: `public/models/LICENSE.txt`

**Interfaces:**
- Consumes: `AvatarMode` and TTS playback callbacks.
- Produces: a responsive WebGL stage with idle movement, gaze, expressions, and speech mouth animation.

- [x] Install the VRM rendering dependencies and a locally bundled licensed sample model.
- [x] Add TTS start/end callbacks without changing synthesis behavior.
- [x] Render and animate the avatar from the pure mode contract.
- [x] Place the stage beside the active conversation without nesting tool cards.
- [x] Run the focused test and renderer build.

### Task 3: Documentation and Visual Verification

**Files:**
- Modify: `README.md`
- Modify: `docs/requirements.md`
- Modify: `CHANGELOG.md`

- [x] Document the AIRI-inspired architecture and local VRM asset boundary.
- [x] Run the complete test suite and `git diff --check`.
- [x] Start the local app and verify desktop and narrow-screen canvas pixels, framing, animation, and non-overlap.
