#!/usr/bin/env python3
"""Generate Interview Copilot domain manifests from the Runoob catalog."""

from __future__ import annotations

import argparse
import json
from pathlib import Path


ALIASES = {"go": "golang"}
SHORT_LABELS = {"assembly": "ASM", "javascript": "JS", "typescript": "TS", "powershell": "PS", "csharp": "C#"}
DOMAINS = {
    "assembly": "systems", "c": "systems", "cpp": "systems", "rust": "systems", "zig": "systems", "verilog": "systems",
    "csharp": "backend", "go": "backend", "java": "backend", "php": "backend", "python": "backend", "ruby": "backend",
    "javascript": "frontend", "typescript": "frontend",
    "dart": "mobile", "kotlin": "mobile", "swift": "mobile",
    "julia": "data", "r": "data", "sql": "data",
    "lua": "language", "perl": "language", "powershell": "language", "scala": "language", "vbscript": "language",
}
TOPICS = {
    "assembly": ["指令集", "寄存器", "内存模型"], "c": ["指针", "内存管理", "编译链接"],
    "cpp": ["C++", "RAII", "现代 C++"], "rust": ["所有权", "生命周期", "并发安全"],
    "zig": ["内存管理", "编译期", "系统编程"], "verilog": ["RTL", "时序逻辑", "仿真"],
    "csharp": ["C#", ".NET", "异步编程"], "go": ["Go", "并发", "工程实践"],
    "java": ["Java", "JVM", "并发"], "php": ["PHP", "Web 开发", "类型系统"],
    "python": ["Python", "语言机制", "异步编程"], "ruby": ["Ruby", "对象模型", "元编程"],
    "javascript": ["JavaScript", "事件循环", "异步编程"], "typescript": ["TypeScript", "类型系统", "泛型"],
    "dart": ["Dart", "异步编程", "Flutter"], "kotlin": ["Kotlin", "协程", "空安全"],
    "swift": ["Swift", "值语义", "并发"], "julia": ["Julia", "多重派发", "科学计算"],
    "r": ["R", "数据分析", "统计计算"], "sql": ["SQL", "查询优化", "事务"],
    "lua": ["Lua", "协程", "元表"], "perl": ["Perl", "正则表达式", "文本处理"],
    "powershell": ["PowerShell", "管道", "自动化"], "scala": ["Scala", "函数式编程", "类型系统"],
    "vbscript": ["VBScript", "脚本自动化", "COM"],
}
ACCENTS = ("#4f9cf9", "#20b486", "#e09f3e", "#d85b73", "#7768d8", "#168aad", "#b56576", "#5f8d4e")


def main() -> None:
    skill_root = Path(__file__).resolve().parents[1]
    repo_root = skill_root.parents[2]
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--catalog", type=Path, default=skill_root / "references/catalog.json")
    parser.add_argument("--output", type=Path, default=repo_root / "internal/skills/manifests/domain/computer")
    parser.add_argument("--force", action="store_true", help="replace existing hand-maintained manifests")
    args = parser.parse_args()

    catalog = json.loads(args.catalog.read_text(encoding="utf-8"))
    args.output.mkdir(parents=True, exist_ok=True)
    written = skipped = 0
    for index, item in enumerate(catalog["languages"]):
        source_slug = item["slug"]
        language = ALIASES.get(source_slug, source_slug)
        path = args.output / f"{language}.json"
        if path.exists() and not args.force:
            skipped += 1
            continue
        name = item["name"]
        topics = TOPICS.get(source_slug, [name, "语言基础", "工程实践"])
        manifest = {
            "id": f"computer-{language}",
            "kind": "domain",
            "name": f"{name} 工程师",
            "shortLabel": SHORT_LABELS.get(source_slug, name[:4]),
            "description": f"覆盖 {name} 语言基础、核心机制与工程实践，参考菜鸟教程 {item['chapters']} 个章节。",
            "industry": "computer",
            "industryName": "计算机",
            "domain": DOMAINS.get(source_slug, "language"),
            "language": language,
            "accent": ACCENTS[index % len(ACCENTS)],
            "knowledgeBaseId": f"computer-{language}",
            "prompt": f"围绕 {name} 的语言机制、常见陷阱和工程实践出题；可参考已同步的菜鸟教程内容，重点考查理解而非语法背诵。",
            "topics": topics,
        }
        path.write_text(json.dumps(manifest, ensure_ascii=False, indent=2) + "\n", encoding="utf-8")
        written += 1
    print(f"wrote {written} manifests, preserved {skipped} existing manifests in {args.output}")


if __name__ == "__main__":
    main()
