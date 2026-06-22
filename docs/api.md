# Interview Copilot API

版本：`v1`

本地地址：`http://127.0.0.1:46831/api/v1`

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

## 面试

### `POST /interviews`

请求：

```json
{
  "candidateName": "小林",
  "resumeId": "resume-a2f04c...",
  "language": "golang",
  "difficulty": "mixed",
  "questionCount": 5
}
```

枚举：

- `language`: `golang | java | python | cpp`
- `difficulty`: `easy | medium | hard | mixed`
- `questionCount`: `1..10`，若超过当前筛选结果则返回实际可用数量

响应 `201`：

```json
{
  "id": "session-17ac...",
  "candidateName": "小林",
  "language": "golang",
  "difficulty": "mixed",
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

最后一题的响应中 `completed` 为 `true`，并附带完整 `report`。`source` 为 `llm` 或 `local`。

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
