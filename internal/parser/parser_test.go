package parser

import (
	"archive/zip"
	"os"
	"path/filepath"
	"testing"
)

func TestParseDOCXAndExtractKeywords(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "resume.docx")
	file, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	archive := zip.NewWriter(file)
	doc, err := archive.Create("word/document.xml")
	if err != nil {
		t.Fatal(err)
	}
	_, _ = doc.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?><w:document xmlns:w="urn:test"><w:body><w:p><w:r><w:t>资深 Golang 工程师</w:t></w:r></w:p><w:p><w:r><w:t>负责 Docker、Kubernetes、Redis 与高并发微服务。</w:t></w:r></w:p></w:body></w:document>`))
	if err := archive.Close(); err != nil {
		t.Fatal(err)
	}
	if err := file.Close(); err != nil {
		t.Fatal(err)
	}

	resume, err := Parse(path, "resume.docx", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	if err != nil {
		t.Fatal(err)
	}
	if resume.Characters == 0 {
		t.Fatal("expected extracted text")
	}
	want := map[string]bool{"Go": true, "Docker": true, "Kubernetes": true, "Redis": true, "高并发": true, "微服务": true}
	for _, keyword := range resume.Keywords {
		delete(want, keyword)
	}
	if len(want) != 0 {
		t.Fatalf("missing keywords: %#v; got %#v", want, resume.Keywords)
	}
}

func TestParseRejectsUnsupportedFile(t *testing.T) {
	_, err := Parse("ignored", "resume.txt", "text/plain")
	if err == nil {
		t.Fatal("expected unsupported extension error")
	}
}

func TestCleanTextCollapsesWhitespace(t *testing.T) {
	got := cleanText(" first  line \r\n\r\n\r\n second\tline ")
	if got != "first line\n\nsecond line" {
		t.Fatalf("unexpected text %q", got)
	}
}
