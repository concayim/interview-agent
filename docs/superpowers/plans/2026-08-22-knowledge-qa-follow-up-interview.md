# Knowledge QA Follow-up Interview Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Make interviews ask knowledge-base QA questions, score cumulative candidate answers against the stored standard answer, and keep asking interviewer-style follow-ups until the score reaches 75.

**Architecture:** The interview service selects QA records from the chosen domain knowledge base before any generated or built-in fallback. The existing evaluator returns the semantic score and a follow-up in one model call; the service owns the 75-point progression rule and keeps prior attempts as context. The existing answer endpoint remains backward compatible and adds follow-up state fields.

**Tech Stack:** Go 1.24, Eino OpenAI-compatible ChatModel, React 19, TypeScript, Vitest.

## Global Constraints

- A score below 75 keeps the current question active.
- A score of 75 or above advances to the next question.
- Follow-ups use the selected interviewer Skill's prompt and feedback tone.
- Model failure falls back to local key-point scoring and local style-aware follow-ups.
- Each answer request remains within the existing 20-second API timeout.

---

### Task 1: Lock the interview progression contract

**Files:**
- Modify: `internal/interview/service_test.go`

**Interfaces:**
- Consumes: `Service.Start`, `Service.Answer`, `domain.Evaluation.Score`.
- Produces: regression coverage for knowledge selection and the 74/75 progression boundary.

- [x] **Step 1: Write failing tests** for knowledge-base-first selection, low-score follow-up, cumulative retry context, and threshold advancement.
- [x] **Step 2: Run `go test ./internal/interview`** and verify failures are caused by missing follow-up behavior.

### Task 2: Implement knowledge QA selection and adaptive follow-ups

**Files:**
- Modify: `internal/domain/types.go`
- Modify: `internal/agent/evaluator.go`
- Modify: `internal/interview/service.go`

**Interfaces:**
- Consumes: `knowledge.Store.Search`, knowledge QA standard answers, interviewer Skill prompt/tone.
- Produces: `Evaluation.FollowUpQuestion`, `AnswerResult.RequiresFollowUp`, and a follow-up `NextQuestion` while `Current` remains unchanged.

- [x] **Step 1: Extend evaluation and session state** with follow-up text and accumulated answer context.
- [x] **Step 2: Select the chosen domain knowledge base first**, including one foundation QA when requested, then fall back to existing generation/built-in paths.
- [x] **Step 3: Update the model prompt** to score semantic coverage against the standard answer and produce one focused, style-aware follow-up below 75.
- [x] **Step 4: Enforce progression in the service**, retaining and combining attempts until the threshold is met.
- [x] **Step 5: Run `go test ./internal/interview ./internal/agent ./internal/knowledge`** and verify green.

### Task 3: Present follow-up state in the interview UI

**Files:**
- Modify: `src/types.ts`
- Modify: `src/App.tsx`

**Interfaces:**
- Consumes: `requiresFollowUp`, `nextQuestion`, and the current evaluation.
- Produces: visible low-score feedback, unchanged progress, a new follow-up prompt, and TTS playback.

- [x] **Step 1: Extend the TypeScript response contract** with `requiresFollowUp`.
- [x] **Step 2: Keep the current progress on follow-up**, replace the active prompt, reset timing, and speak the interviewer follow-up.
- [x] **Step 3: Run frontend unit tests and TypeScript build**.

### Task 4: Document, verify, and ship

**Files:**
- Modify: `README.md`
- Modify: `docs/requirements.md`
- Modify: `docs/api.md`
- Modify: `CHANGELOG.md`

**Interfaces:**
- Produces: user requirements, API contract, setup behavior, and release history matching implementation.

- [x] **Step 1: Document knowledge-first questions, semantic scoring, the 75-point threshold, and response fields.**
- [x] **Step 2: Run `npm test` and `git diff --check`.**
- [x] **Step 3: Review only task-related diffs, commit them, and push `codex/skill-knowledge-learning`.**
