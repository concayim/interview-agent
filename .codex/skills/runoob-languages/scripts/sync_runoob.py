#!/usr/bin/env python3
"""Synchronize Runoob programming-language tutorials into Markdown references."""

from __future__ import annotations

import argparse
import concurrent.futures
import hashlib
import html
import json
import os
import re
import tempfile
import time
from dataclasses import dataclass
from datetime import datetime, timezone
from html.parser import HTMLParser
from pathlib import Path
from typing import Iterable
from urllib.error import HTTPError, URLError
from urllib.parse import urljoin, urlparse, urlunparse
from urllib.request import Request, urlopen

USER_AGENT = "runoob-languages-skill/1.0 (+local educational sync)"
TIMEOUT = 25


@dataclass(frozen=True)
class Language:
    slug: str
    name: str
    url: str
    aliases: tuple[str, ...] = ()
    extra_prefixes: tuple[str, ...] = ()


LANGUAGES = (
    Language("python", "Python", "https://www.runoob.com/python3/python3-tutorial.html", ("python3", "py")),
    Language("r", "R", "https://www.runoob.com/r/r-tutorial.html"),
    Language("julia", "Julia", "https://www.runoob.com/julia/julia-tutorial.html"),
    Language("javascript", "JavaScript", "https://www.runoob.com/js/js-tutorial.html", ("js",)),
    Language("typescript", "TypeScript", "https://www.runoob.com/typescript/ts-tutorial.html", ("ts",)),
    Language("php", "PHP", "https://www.runoob.com/php/php-tutorial.html"),
    Language("java", "Java", "https://www.runoob.com/java/java-tutorial.html"),
    Language("go", "Go", "https://www.runoob.com/go/go-tutorial.html", ("golang",)),
    Language("rust", "Rust", "https://www.runoob.com/rust/rust-tutorial.html"),
    Language("csharp", "C#", "https://www.runoob.com/csharp/csharp-tutorial.html", ("c#", "cs", "dotnet")),
    Language("vbscript", "VBScript", "https://www.runoob.com/vbscript/vbscript-tutorial.html", ("vbs",)),
    Language("swift", "Swift", "https://www.runoob.com/swift/swift-tutorial.html"),
    Language("kotlin", "Kotlin", "https://www.runoob.com/kotlin/kotlin-tutorial.html"),
    Language("c", "C", "https://www.runoob.com/c/c-tutorial.html", extra_prefixes=("/cprogramming/",)),
    Language("cpp", "C++", "https://www.runoob.com/cplusplus/cpp-tutorial.html", ("c++", "cplusplus")),
    Language("zig", "Zig", "https://www.runoob.com/zig/zig-tutorial.html"),
    Language("scala", "Scala", "https://www.runoob.com/scala/scala-tutorial.html"),
    Language("ruby", "Ruby", "https://www.runoob.com/ruby/ruby-tutorial.html"),
    Language("perl", "Perl", "https://www.runoob.com/perl/perl-tutorial.html"),
    Language("lua", "Lua", "https://www.runoob.com/lua/lua-tutorial.html"),
    Language("dart", "Dart", "https://www.runoob.com/dart/dart-tutorial.html"),
    Language("assembly", "Assembly", "https://www.runoob.com/assembly/assembly-tutorial.html", ("asm", "汇编")),
    Language("verilog", "Verilog", "https://www.runoob.com/w3cnote/verilog-tutorial.html"),
    Language("powershell", "PowerShell", "https://www.runoob.com/powershell/powershell-tutorial.html", ("pwsh",)),
    Language("sql", "SQL", "https://www.runoob.com/sql/sql-tutorial.html"),
)


class TutorialParser(HTMLParser):
    def __init__(self) -> None:
        super().__init__(convert_charrefs=True)
        self.nav_depth = 0
        self.article_depth = 0
        self.skip_depth = 0
        self.links: list[tuple[str, str]] = []
        self.current_href = ""
        self.current_link_text: list[str] = []
        self.title = ""
        self.title_depth = 0
        self.parts: list[str] = []
        self.in_pre = 0

    def handle_starttag(self, tag: str, attrs: list[tuple[str, str | None]]) -> None:
        values = dict(attrs)
        classes = set((values.get("class") or "").split())
        if tag == "div" and values.get("id") == "leftcolumn":
            self.nav_depth = 1
        elif self.nav_depth and tag == "div":
            self.nav_depth += 1
        if tag == "div" and "article-intro" in classes:
            self.article_depth = 1
        elif self.article_depth and tag == "div":
            self.article_depth += 1
        if self.article_depth:
            if tag in {"script", "style", "form", "noscript"}:
                self.skip_depth += 1
            if self.skip_depth:
                return
            if tag == "h1":
                self.title_depth = 1
            elif self.title_depth:
                self.title_depth += 1
            if tag in {"h1", "h2", "h3", "h4"}:
                self.parts.append("\n\n" + "#" * min(int(tag[1]) + 1, 6) + " ")
            elif tag in {"p", "div", "ul", "ol", "table", "blockquote"}:
                self.parts.append("\n\n")
            elif tag == "li":
                self.parts.append("\n- ")
            elif tag == "br":
                self.parts.append("\n")
            elif tag == "pre":
                self.in_pre += 1
                self.parts.append("\n\n```\n")
            elif tag in {"code", "kbd"} and not self.in_pre:
                self.parts.append("`")
        if self.nav_depth and tag == "a" and values.get("href"):
            self.current_href = values["href"] or ""
            self.current_link_text = []

    def handle_endtag(self, tag: str) -> None:
        if self.article_depth and not self.skip_depth:
            if tag == "pre" and self.in_pre:
                self.in_pre -= 1
                self.parts.append("\n```\n")
            elif tag in {"code", "kbd"} and not self.in_pre:
                self.parts.append("`")
        if tag in {"script", "style", "form", "noscript"} and self.skip_depth:
            self.skip_depth -= 1
        if self.title_depth:
            self.title_depth -= 1
        if self.current_href and tag == "a":
            text = clean_inline(" ".join(self.current_link_text))
            self.links.append((self.current_href, text))
            self.current_href = ""
            self.current_link_text = []
        if self.nav_depth and tag == "div":
            self.nav_depth -= 1
        if self.article_depth and tag == "div":
            self.article_depth -= 1

    def handle_data(self, data: str) -> None:
        if self.current_href:
            self.current_link_text.append(data)
        if self.article_depth and not self.skip_depth:
            value = data if self.in_pre else re.sub(r"\s+", " ", data)
            if self.title_depth:
                self.title += value
            self.parts.append(value)


def clean_inline(value: str) -> str:
    return re.sub(r"\s+", " ", html.unescape(value)).strip()


def canonical_url(value: str) -> str:
    parsed = urlparse(value)
    scheme = "https"
    host = parsed.netloc.lower() or "www.runoob.com"
    if host == "runoob.com":
        host = "www.runoob.com"
    path = re.sub(r"/{2,}", "/", parsed.path)
    return urlunparse((scheme, host, path, "", "", ""))


def fetch(url: str, attempts: int = 3) -> str:
    request = Request(url, headers={"User-Agent": USER_AGENT, "Accept": "text/html,application/xhtml+xml"})
    for attempt in range(attempts):
        try:
            with urlopen(request, timeout=TIMEOUT) as response:
                content_type = response.headers.get_content_type()
                if content_type not in {"text/html", "application/xhtml+xml"}:
                    raise RuntimeError(f"unexpected content type {content_type}")
                return response.read().decode(response.headers.get_content_charset() or "utf-8", errors="replace")
        except (HTTPError, URLError, TimeoutError) as exc:
            if attempt + 1 == attempts:
                raise RuntimeError(f"fetch {url}: {exc}") from exc
            time.sleep(0.5 * (2**attempt))
    raise AssertionError("unreachable")


def tutorial_prefix(language: Language) -> str:
    parts = [part for part in urlparse(language.url).path.split("/") if part]
    return "/" + parts[0] + "/" if parts else "/"


def discover(language: Language) -> list[tuple[str, str]]:
    parser = TutorialParser()
    parser.feed(fetch(language.url))
    prefix = tutorial_prefix(language)
    allowed_prefixes = (prefix, *language.extra_prefixes)
    result: list[tuple[str, str]] = []
    seen: set[str] = set()
    for href, label in [(language.url, language.name + " 教程"), *parser.links]:
        url = canonical_url(urljoin(language.url, href))
        parsed = urlparse(url)
        if parsed.netloc != "www.runoob.com" or not parsed.path.startswith(allowed_prefixes) or not parsed.path.endswith((".html", "/")):
            continue
        if url in seen:
            continue
        seen.add(url)
        result.append((url, label or parsed.path.rsplit("/", 1)[-1]))
    return result


def normalize_markdown(parts: Iterable[str]) -> str:
    value = "".join(parts).replace("\xa0", " ")
    value = re.sub(r"[ \t]+\n", "\n", value)
    value = re.sub(r"\n[ \t]+", "\n", value)
    value = re.sub(r"\n{3,}", "\n\n", value)
    value = re.sub(r"(?m)^-\s*\n+", "- ", value)
    value = re.sub(r" {2,}", " ", value)
    return value.strip()


def fetch_chapter(item: tuple[str, str]) -> tuple[str, str, str]:
    url, nav_title = item
    parser = TutorialParser()
    parser.feed(fetch(url))
    title = clean_inline(parser.title) or clean_inline(nav_title)
    body = normalize_markdown(parser.parts)
    if not body:
        raise RuntimeError(f"no article content at {url}")
    return title, url, body


def write_atomic(path: Path, content: str) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    fd, temp_name = tempfile.mkstemp(prefix=path.name + ".", dir=path.parent)
    try:
        with os.fdopen(fd, "w", encoding="utf-8") as output:
            output.write(content)
        os.replace(temp_name, path)
    finally:
        if os.path.exists(temp_name):
            os.unlink(temp_name)


def sync_language(language: Language, output_dir: Path, jobs: int, max_pages: int | None) -> dict[str, object]:
    chapters = discover(language)
    if max_pages is not None:
        chapters = chapters[:max_pages]
    with concurrent.futures.ThreadPoolExecutor(max_workers=jobs) as executor:
        results = list(executor.map(fetch_chapter, chapters))
    sections = [f"# {language.name} - 菜鸟教程\n\nTutorial: {language.url}\n"]
    for title, url, body in results:
        sections.append(f"\n---\n\n## {title}\n\nSource: {url}\n\n{body}\n")
    content = "".join(sections)
    write_atomic(output_dir / f"{language.slug}.md", content)
    return {
        "slug": language.slug,
        "name": language.name,
        "aliases": list(language.aliases),
        "tutorial_url": language.url,
        "file": f"{language.slug}.md",
        "chapters": len(results),
        "sha256": hashlib.sha256(content.encode()).hexdigest(),
    }


def select_languages(values: list[str]) -> list[Language]:
    if not values:
        return list(LANGUAGES)
    lookup = {key.lower(): item for item in LANGUAGES for key in (item.slug, item.name, *item.aliases)}
    selected: list[Language] = []
    for value in values:
        item = lookup.get(value.lower())
        if not item:
            raise SystemExit(f"unknown language: {value}")
        if item not in selected:
            selected.append(item)
    return selected


def main() -> None:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--language", action="append", default=[], help="language slug/name; repeatable")
    parser.add_argument("--jobs", type=int, default=3, help="parallel requests per language (default: 3)")
    parser.add_argument("--max-pages", type=int, help="limit chapters per language for testing")
    parser.add_argument("--output", type=Path, default=Path(__file__).resolve().parents[1] / "references")
    parser.add_argument("--list", action="store_true", help="list supported languages")
    args = parser.parse_args()
    if args.list:
        for language in LANGUAGES:
            print(f"{language.slug:12} {language.name:12} {language.url}")
        return
    if not 1 <= args.jobs <= 8:
        parser.error("--jobs must be between 1 and 8")
    selected = select_languages(args.language)
    existing: dict[str, dict[str, object]] = {}
    catalog_path = args.output / "catalog.json"
    if catalog_path.exists():
        data = json.loads(catalog_path.read_text(encoding="utf-8"))
        existing = {item["slug"]: item for item in data.get("languages", [])}
    for language in selected:
        print(f"sync {language.name} ...", flush=True)
        existing[language.slug] = sync_language(language, args.output, args.jobs, args.max_pages)
        print(f"  {existing[language.slug]['chapters']} chapters", flush=True)
    catalog = {
        "source": "https://www.runoob.com/",
        "synced_at": datetime.now(timezone.utc).isoformat(),
        "languages": sorted(existing.values(), key=lambda item: str(item["slug"])),
    }
    write_atomic(catalog_path, json.dumps(catalog, ensure_ascii=False, indent=2) + "\n")
    print(f"wrote {catalog_path}")


if __name__ == "__main__":
    main()
