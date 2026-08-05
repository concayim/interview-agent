---
name: runoob-languages
description: Synchronize, search, and use the Chinese programming-language tutorials published on runoob.com (菜鸟教程). Use for programming syntax, concepts, examples, language comparisons, study notes, interview preparation, or when the user asks to pull, refresh, cite, or answer from Runoob language tutorials. Covers Python, R, Julia, JavaScript, TypeScript, PHP, Java, Go, Rust, C#, VBScript, Swift, Kotlin, C, C++, Zig, Scala, Ruby, Perl, Lua, Dart, Assembly, Verilog, PowerShell, and SQL.
---

# Runoob Languages

Use the locally synchronized Runoob corpus for programming-language questions. Keep answers concise, distinguish tutorial claims from current language behavior, and include source URLs when the user asks for citations.

## Find content

1. Read `references/catalog.json` to resolve a language name or alias to its reference file.
2. Search only the relevant file first:

   ```bash
   rg -n -i "search terms" references/<language>.md
   ```

3. Read the matching section and its `Source:` line. Expand to related languages only for comparisons.
4. If a reference file is missing or stale, run the synchronization command below before answering.

## Synchronize

Run from this skill directory:

```bash
python3 scripts/sync_runoob.py
```

Useful options:

```bash
python3 scripts/sync_runoob.py --language rust
python3 scripts/sync_runoob.py --language c --language cpp --jobs 4
python3 scripts/sync_runoob.py --list
```

The script discovers each tutorial's chapter links from its navigation, fetches only `www.runoob.com` HTML pages within that tutorial path, strips navigation/comments/ads, and writes one Markdown file per language plus `references/catalog.json`. It is idempotent: a completed sync replaces generated reference files atomically.

To expose synchronized languages in the Interview Copilot frontend, run:

```bash
python3 scripts/sync_app_skills.py
```

This generates missing domain manifests under `internal/skills/manifests/domain/computer` and preserves hand-maintained manifests. Pass `--force` only when intentionally replacing them.

Respect `robots.txt`, use the default conservative concurrency, and do not crawl unrelated notes, comments, search results, or interactive runners. Do not edit generated reference files manually; update the script or resynchronize.

## Answering rules

- Prefer the synchronized corpus when the request explicitly names Runoob or 菜鸟教程.
- Treat it as an introductory secondary source. For version-sensitive, security-sensitive, or disputed behavior, verify against official language documentation and state the difference.
- Cite the exact chapter URL stored beside the section, not only the Runoob home page.
- Do not imply the corpus is current beyond the `synced_at` value in `catalog.json`.
- When code from the tutorial is outdated, provide a corrected modern example and briefly identify the compatibility issue.
