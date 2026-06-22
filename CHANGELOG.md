# Changelog

本项目遵循 [Keep a Changelog](https://keepachangelog.com/zh-CN/1.1.0/) 的记录方式。

## [Unreleased]

### Planned

- 持久化面试历史与题库管理。
- 为扫描版 PDF 增加可选 OCR。
- 建立 GitHub Actions 跨平台构建与自动发布流水线。

## [0.2.0] - 2026-06-23

### Added

- 新增 Echo、Atlas、Vera、Socrates 四个面试官风格 Skill，并将风格提示、评价重点和反馈语气接入 Eino Agent。
- 新增可扩展行业/领域 Skill Manifest；当前提供计算机基础、Golang、Java、Python、C++。
- 新增 6 道计算机基础公共题，可选择混入语言专项面试。
- 新增 `computer-foundation` 与四个语言独立知识库，支持 QA 持久化、检索、手动新增和出题次数统计。
- 面试创建时自动把实际出题记录到对应知识库。
- 新增学习中心，实时聚合 Go、Java、Spring、Python、C++ 权威 Feed。
- 新增官方文档、课程和视频入口，以及本地学习清单。
- 新增 Skill、知识库和学习资源 REST API。

### Changed

- 准备页从固定语言选项升级为动态行业/领域 Skill 选择。
- 面试页和复盘页展示实际面试官 Skill 与领域 Skill。
- 版本升级至 `0.2.0`。

### Security

- Feed 条目只接受与权威来源同域的 HTTPS 链接，页面内容始终按纯文本展示。

## [0.1.0] - 2026-06-22

### Added

- 初始化 Go、Eino、Electron、React 工程。
- 新增 PDF / DOCX 简历解析、关键词提取与 10 MB 上传限制。
- 新增 Golang、Java、Python、C++ 共 24 道标准化题目。
- 新增简历相关性排序、面试状态机、计时和进度管理。
- 新增 Eino OpenAI-compatible 模型评价与本地评分降级。
- 新增逐题反馈、综合报告和标准答案复盘。
- 新增模型设置、本机凭据持久化和连接测试。
- 新增 Electron 本地服务托管、上下文隔离和随机访问令牌。
- 新增中文需求文档、接口文档、README 与测试用例。

### Security

- Go API 仅监听回环地址。
- Electron 每次启动生成随机 API 访问令牌。
- 简历临时文件解析后删除，模型 API Key 文件权限设为 `0600`。
