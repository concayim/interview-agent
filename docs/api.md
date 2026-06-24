# Interview Copilot API

版本：`v1`（应用 `v0.3.0`）

独立启动后端的默认地址：`http://127.0.0.1:46831/api/v1`。Electron 模式会为每次启动动态分配空闲端口，并通过 preload 向渲染进程提供本次 API 地址。

所有响应均为 UTF-8 JSON。除健康检查和 `OPTIONS` 外，Electron 模式下请求需要：

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
      "feedbackTone": "直接、严格、不刻薄",
      "followUpRounds": 3
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
  "difficulty": "mixed",
  "questionCount": 5
}
```

枚举：

- `domainSkillId`: 当前为 `computer-golang | computer-java | computer-python | computer-cpp`
- `interviewerSkillId`: `echo-coach | atlas-architect | vera-challenger | socrates-guide`
- `includeFoundation`: 是否混入计算机基础公共库
- `difficulty`: `easy | medium | hard | mixed`
- `questionCount`: 自定义主问题数量，接口范围 `1..20`；实际数量不超过当前难度与知识库的可用题数

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
  "status": "active",
  "phase": "introduction",
  "current": 0,
  "total": 5,
  "followUpRound": 0,
  "followUpTotal": 3,
  "currentQuestion": {
    "id": "introduction",
    "promptId": "session-17ac...-introduction",
    "stage": "introduction",
    "language": "golang",
    "difficulty": "mixed",
    "prompt": "小林，正式开始前，请先用 1–2 分钟做个自我介绍……",
    "tags": ["自我介绍"],
    "followUp": false,
    "round": 0,
    "followUpTotal": 0
  },
  "startedAt": "2026-06-22T12:00:00+08:00"
}
```

注意：进行中的面试不会返回 `standardAnswer` 与 `keyPoints`。

### `GET /interviews/{id}`

返回当前面试状态，结构同创建响应。不存在返回 `404`。

### `POST /interviews/{id}/answers`

请求：

```json
{
  "answer": "我主要负责 Go 后端和 Redis 高并发项目，在团队里承担核心服务设计与稳定性治理。",
  "elapsedSeconds": 76
}
```

首次提交用于完成自我介绍。该阶段不参与技术评分，`current` 保持 `0`；响应会返回按简历与介绍关键词排序后的第一个技术问题：

```json
{
  "evaluation": {
    "score": 0,
    "summary": "自我介绍已记录，接下来的问题会结合你的简历与刚才提到的经历。",
    "strengths": ["完成面试开场"],
    "improvements": [],
    "source": "local"
  },
  "questionCompleted": true,
  "completed": false,
  "phase": "technical",
  "nextQuestion": {
    "id": "foundation-04",
    "promptId": "foundation-04",
    "stage": "technical",
    "leadIn": "谢谢你的介绍。我看到你的简历里提到了「Redis」，我们就从这段经历展开。",
    "language": "foundation",
    "difficulty": "medium",
    "prompt": "什么是缓存穿透、击穿和雪崩？分别如何治理？",
    "tags": ["缓存", "Redis", "高并发"],
    "followUp": false,
    "round": 0,
    "followUpTotal": 3
  },
  "current": 0,
  "total": 5,
  "followUpRound": 0,
  "followUpTotal": 3
}
```

后续使用同一接口提交技术回答：

```json
{
  "answer": "goroutine 由 Go runtime 调度…",
  "elapsedSeconds": 95
}
```

提交主回答后，通常返回第 1 轮追问：

```json
{
  "evaluation": {
    "score": 82,
    "summary": "核心机制基本正确。",
    "strengths": ["说明了 GMP 与轻量调度"],
    "improvements": ["补充动态栈和工作窃取"],
    "source": "llm"
  },
  "questionCompleted": false,
  "completed": false,
  "phase": "technical",
  "nextQuestion": {
    "id": "go-01",
    "promptId": "go-01-followup-1",
    "stage": "follow_up",
    "language": "golang",
    "difficulty": "easy",
    "prompt": "追问 1/3：如果把 goroutine 放进真实系统，你会如何定义约束和容量？",
    "tags": ["Go"],
    "followUp": true,
    "round": 1,
    "followUpTotal": 3
  },
  "current": 0,
  "total": 5,
  "followUpRound": 1,
  "followUpTotal": 3
}
```

完成一个主问题的最后一轮追问时，`questionCompleted=true`，`current` 加一，`nextQuestion` 变为下一主问题。最后一个主问题完成时 `completed=true`、`phase=completed`，并附带完整 `report`。`source` 为 `llm` 或 `local`。

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
  "introduction": {
    "prompt": "小林，正式开始前，请先用 1–2 分钟做个自我介绍……",
    "answer": "我主要负责 Go 后端和 Redis 高并发项目……",
    "elapsedSeconds": 76
  },
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
      "evaluation": { "score": 82, "summary": "…", "strengths": [], "improvements": [], "source": "llm" },
      "followUps": [
        {
          "round": 1,
          "prompt": "如果放进真实系统，你会如何定义约束和容量？",
          "answer": "我会先估算并发量…",
          "elapsedSeconds": 61,
          "evaluation": { "score": 86, "summary": "…", "strengths": [], "improvements": [], "source": "llm" }
        }
      ],
      "averageScore": 84
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
  "enabled": true,
  "clearApiKey": false
}
```

- `apiKey` 留空且 `clearApiKey=false` 时保留已有 Key。
- `clearApiKey=true` 时清除 Key。
- 启用模型时必须已有或提供 API Key，并填写 Model。

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
