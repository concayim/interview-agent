# Changelog

本项目遵循 [Keep a Changelog](https://keepachangelog.com/zh-CN/1.1.0/) 的记录方式。

## [Unreleased]

### Planned

- 持久化面试历史与题库管理。
- 为扫描版 PDF 增加可选 OCR。
- GitHub 远端地址确认后建立自动化发布流水线。

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
