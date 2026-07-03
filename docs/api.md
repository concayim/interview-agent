# Interview Copilot API

版本：`v1`（应用 `v0.2.0`）

本地地址：默认 `http://127.0.0.1:46831/api/v1`。Electron 启动时如果端口已被占用，会自动选择可用端口，并通过主进程运行时配置把实际地址与令牌注入前端。

所有响应均为 UTF-8 JSON。除健康检查、Skill 目录和 `OPTIONS` 外，Electron 模式下请求需要：

```http
X-Interview-Agent-Token: <本次启动生成的随机令牌>
```

错误格式：

```json
{ "error": "可读的错误信息" }
```

## 健康检查

### `GET /health`

响应 `200`：

```json
{
  "status": "ok",
  "time": "2026-06-22T12:00:00Z",
  "modelConfigured": false
}
```

## 简历

### `POST /resumes`

`Content-Type: multipart/form-data`，字段名为 `file`。支持 `.pdf`、`.docx`，最大 10 MB。

响应 `201`：

```json
{
  "id": "resume-a2f04c...",
  "fileName": "resume.pdf",
  "contentType": "application/pdf",
  "preview": "候选人简历内容摘要…",
  "keywords": ["Go", "Redis", "Docker", "微服务"],
  "characters": 2831,
  "createdAt": "2026-06-22T12:00:00+08:00"
}
```

可能状态：`400` 文件过大或缺少文件；`415` 格式不支持；`422` 文件损坏或没有可读文本。

## Skill 目录

### `GET /skills`

返回动态面试官 Skill、领域 Skill 和行业目录。Skill 的私有提示词不会返回给前端。

该接口为公开只读接口，不要求 `X-Interview-Agent-Token`，用于保证准备页在运行时端口或令牌刷新后仍能稳定渲染 Skill 目录。

```json
{
  "interviewers": [
    {
      "id": "vera-challenger",
      "kind": "interviewer",
      "name": "Vera · 压力挑战官",
      "shortLabel": "Vera",
      "description": "节奏直接、标准严格…",
      "accent": "#ff8194",
      "openingLine": "我会比较直接…",
      "evaluationFocus": ["事实准确性", "边界条件", "抗压表达", "反例意识"],
      "feedbackTone": "直接、严格、不刻薄"
    }
  ],
  "domains": [
    {
      "id": "computer-golang",
      "kind": "domain",
      "name": "Golang 工程师",
      "industry": "computer",
      "industryName": "计算机",
      "language": "golang",
      "knowledgeBaseId": "computer-golang",
      "topics": ["Go", "GMP", "并发"]
    }
  ],
  "industries": [{ "id": "computer", "name": "计算机" }]
}
```

## 面试

### `POST /interviews`

请求：

```json
{
  "candidateName": "小林",
  "resumeId": "resume-a2f04c...",
  "domainSkillId": "computer-golang",
  "interviewerSkillId": "atlas-architect",
  "includeFoundation": true,
  "videoEnabled": true,
  "speechLanguage": "zh-CN",
  "difficulty": "mixed",
  "questionCount": 5
}
```

枚举：

- `domainSkillId`: 当前为 `computer-golang | computer-java | computer-python | computer-cpp`
- `interviewerSkillId`: `echo-coach | atlas-architect | vera-challenger | socrates-guide`
- `includeFoundation`: 是否混入计算机基础公共库
- `videoEnabled`: 是否开启视频面试
- `speechLanguage`: 语音输入语言，支持 `zh-CN | en-US`
- `difficulty`: `easy | medium | hard | mixed`
- `questionCount`: `1..10`，若超过当前筛选结果则返回实际可用数量

已配置模型时，服务端会优先按领域 Skill、难度、简历关键词和基础库偏好动态生成本场题目；生成失败时回退内置题库。

响应 `201`：

```json
{
  "id": "session-17ac...",
  "candidateName": "小林",
  "language": "golang",
  "difficulty": "mixed",
  "industry": "computer",
  "domainSkillId": "computer-golang",
  "domainSkillName": "Golang 工程师",
  "interviewerSkillId": "atlas-architect",
  "interviewerName": "Atlas · 架构面试官",
  "interviewerOpening": "我会追问设计背后的约束和取舍…",
  "includeFoundation": true,
  "videoEnabled": true,
  "speechLanguage": "zh-CN",
  "questionSource": "model",
  "status": "active",
  "current": 0,
  "total": 5,
  "currentQuestion": {
    "id": "go-01",
    "language": "golang",
    "difficulty": "easy",
    "prompt": "Go 的 goroutine 和操作系统线程有什么区别？",
    "tags": ["Go", "高并发", "调度器"]
  },
  "startedAt": "2026-06-22T12:00:00+08:00"
}
```

注意：进行中的面试不会返回 `standardAnswer` 与 `keyPoints`。
`questionSource` 为 `model` 或 `built-in`，用于标识本场题目来源。

### `GET /interviews/{id}`

返回当前面试状态，结构同创建响应。不存在返回 `404`。

### `POST /interviews/{id}/answers`

请求：

```json
{
  "answer": "goroutine 由 Go runtime 调度…",
  "elapsedSeconds": 95
}
```

未完成时响应 `200`：

```json
{
  "accepted": true,
  "intent": "answer",
  "evaluation": {
    "score": 82,
    "summary": "核心机制基本正确。",
    "strengths": ["说明了 GMP 与轻量调度"],
    "improvements": ["补充动态栈和工作窃取"],
    "source": "llm"
  },
  "completed": false,
  "nextQuestion": { "id": "go-02", "language": "golang", "difficulty": "easy", "prompt": "…", "tags": ["Go"] },
  "current": 1,
  "total": 5
}
```

`accepted=true` 表示本次输入已被当作正式回答或跳过请求并推进题目。`intent` 可能为 `answer` 或 `skip`；跳过时会生成 `0` 分本地评价并进入下一题。
回答链路服务端超时为 20 秒；模型意图识别、模型评价或本地评分会在该窗口内完成，模型不可用时自动降级。

如果用户是在求提示、要求解释或重复题目，响应不会推进当前题：

```json
{
  "accepted": false,
  "intent": "hint",
  "assistantReply": "可以按这个顺序组织：先定义…",
  "completed": false,
  "current": 0,
  "total": 5
}
```

这类 `intent` 可能为 `hint | clarify | repeat | off_topic | smalltalk`，不会返回 `evaluation`。其中 `off_topic` 与 `smalltalk` 由大模型意图识别处理，用于理解与当前题不直接相关的输入，并把对话拉回面试。最后一题的正式回答响应中 `completed` 为 `true`，并附带完整 `report`。`source` 为 `llm` 或 `local`。

### `POST /speech/transcriptions`

`Content-Type: multipart/form-data`，字段名为 `audio`，可选字段 `language` 为 `zh`、`en`、`zh-CN` 或 `en-US`。该接口使用当前模型配置中的 `Base URL`、`API Key` 与 `Speech Model` 调用 OpenAI-compatible `/audio/transcriptions`，用于 Electron 不支持 Web Speech API 时的录音转写兜底。单段语音最大 12 MB。

响应 `200`：

```json
{ "text": "转写后的回答文本" }
```

需要在模型设置中填写 `Speech Model`；未填写时返回 `400`，上游模型不可用或不支持音频时返回 `502`，错误信息会优先提取服务商返回的 `error.message` 或 `message`。

### `POST /speech/transcriptions/stream`

请求体同 `/speech/transcriptions`。服务端会向上游 `/audio/transcriptions` 追加 `stream=true`，并以 `text/event-stream` 返回统一的流式事件。前端语音兜底默认使用该接口，把增量文本写入回答框。

事件格式：

```text
event: delta
data: {"text":"增量文字"}

event: done
data: {"text":"完整文本，可为空"}

event: error
data: {"error":"可读的错误信息"}
```

`delta` 是前端写入回答框的主要增量；`done` 只表示结束，`text` 可能为空。服务端会兼容 `delta`、`choices[].delta.content`、`text`、`transcript` 等常见上游字段，并支持 LF / CRLF 事件分隔。若上游不支持 SSE 但返回普通 JSON `{ "text": "..." }`，服务端会发送一次 `delta` 后再发送 `done`。若上游因不支持 `stream=true` 返回 400 / 404 / 422，服务端会自动重试普通转写并以同样的 SSE 格式返回最终文本。

### `GET /interviews/{id}/report`

仅完成后可访问；未完成返回 `409`。

响应 `200` 的关键结构：

```json
{
  "sessionId": "session-17ac...",
  "candidateName": "小林",
  "language": "golang",
  "difficulty": "mixed",
  "score": 78,
  "answered": 5,
  "durationSeconds": 1032,
  "highlights": ["覆盖了 goroutine 和 GMP"],
  "focusAreas": ["补充并发取消语义"],
  "answers": [
    {
      "question": {
        "id": "go-01",
        "prompt": "…",
        "standardAnswer": "…",
        "keyPoints": ["goroutine", "gmp"],
        "tags": ["Go", "高并发"]
      },
      "answer": "…",
      "elapsedSeconds": 95,
      "evaluation": { "score": 82, "summary": "…", "strengths": [], "improvements": [], "source": "llm" }
    }
  ],
  "startedAt": "2026-06-22T12:00:00+08:00",
  "completedAt": "2026-06-22T12:17:12+08:00"
}
```

## 模型配置

### `GET /config/model`

API Key 永远不会返回。

```json
{
  "baseUrl": "https://api.example.com/v1",
  "model": "model-name",
  "speechModel": "whisper-1",
  "enabled": true,
  "hasApiKey": true
}
```

### `PUT /config/model`

```json
{
  "apiKey": "sk-...",
  "baseUrl": "https://api.example.com/v1",
  "model": "model-name",
  "speechModel": "whisper-1",
  "enabled": true,
  "clearApiKey": false
}
```

- `apiKey` 留空且 `clearApiKey=false` 时保留已有 Key。
- `clearApiKey=true` 时清除 Key。
- 启用模型时必须已有或提供 API Key，并填写 `model`。
- `speechModel` 仅用于 `/speech/transcriptions` 的录音兜底转写；不填写时不影响出题、意图识别和点评，但录音转写会返回配置提示。

响应结构同 `GET`。

### `POST /config/model/test`

使用已保存配置发起一次最小模型请求。

成功 `200`：

```json
{ "ok": true, "message": "模型连接成功" }
```

未配置或上游失败返回 `502`。

## 知识库

### `GET /knowledge/bases`

返回公共库与各语言库，以及 QA 数和累计出题次数。

```json
{
  "bases": [
    {
      "id": "computer-foundation",
      "name": "计算机基础公共库",
      "language": "foundation",
      "topics": ["数据结构", "操作系统", "计算机网络"],
      "itemCount": 6,
      "issuedCount": 3
    }
  ]
}
```

### `GET /knowledge/bases/{id}/qa?query=TCP&limit=30`

按问题、答案和标签检索指定知识库。`query` 为空时按出题次数与更新时间返回。

```json
{
  "items": [
    {
      "id": "qa-foundation-02",
      "baseId": "computer-foundation",
      "question": "TCP 为什么需要三次握手，而不是两次？",
      "answer": "三次握手让双方确认…",
      "keyPoints": ["双向/bidirectional", "序列号/sequence number"],
      "tags": ["计算机网络", "TCP", "基础"],
      "difficulty": "easy",
      "source": "built-in",
      "issuedCount": 1
    }
  ]
}
```

### `POST /knowledge/bases/{id}/qa`

手动新增或按问题内容更新 QA：

```json
{
  "question": "什么是时间复杂度？",
  "answer": "描述输入规模增长时算法耗时的渐近趋势。",
  "keyPoints": ["渐近", "输入规模"],
  "tags": ["算法"],
  "difficulty": "easy"
}
```

成功返回 `201` 和完整 QA。创建面试时，服务会自动对实际选中的 QA 执行 upsert 并增加 `issuedCount`。

## 学习中心

### `GET /learning/resources`

查询参数：

- `domainSkillId`: `all` 或领域 Skill ID；
- `kind`: `all | article | video | course | docs`；
- `selectedOnly`: `true | false`。

```json
{
  "resources": [
    {
      "id": "docs-go-tour",
      "title": "A Tour of Go",
      "url": "https://go.dev/tour/",
      "kind": "docs",
      "source": "Go Documentation",
      "authority": "Go Team",
      "summary": "Go 官方交互式语言导览…",
      "domainSkillIds": ["computer-golang"],
      "selected": true,
      "live": false
    }
  ],
  "refreshedAt": "2026-06-23T00:35:42+08:00",
  "warnings": []
}
```

### `POST /learning/refresh`

并发抓取五个权威 Feed 的标题、链接、短摘要与发布时间。单源失败写入 `warnings`，其余结果仍返回 `200`。请求超时 20 秒。

### `PUT /learning/resources/{id}/selection`

```json
{ "selected": true }
```

将资源加入或移出本地学习清单。响应：`{ "id": "...", "selected": true }`。

## 状态码

| 状态码 | 含义 |
| --- | --- |
| `200` | 请求成功 |
| `201` | 资源创建成功 |
| `204` | CORS 预检成功 |
| `400` | 参数或请求体错误 |
| `401` | 本地访问令牌无效 |
| `404` | 资源不存在 |
| `409` | 当前状态不允许操作 |
| `415` | 附件格式不支持 |
| `422` | 附件无法读取 |
| `502` | 大模型连接或调用失败 |
| `500` | 服务内部错误 |
