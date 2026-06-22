package parser

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"interview-agent/internal/domain"
)

var ErrNoText = errors.New("附件中没有可读取的文本；扫描版 PDF 暂不支持，请先进行 OCR")

func Parse(path, originalName, contentType string) (domain.Resume, error) {
	ext := strings.ToLower(filepath.Ext(originalName))
	var text string
	var err error
	switch ext {
	case ".docx":
		text, err = parseDOCX(path)
	case ".pdf":
		text, err = parsePDF(path)
	default:
		return domain.Resume{}, fmt.Errorf("不支持的附件格式 %q，仅支持 .docx 和 .pdf", ext)
	}
	if err != nil {
		return domain.Resume{}, err
	}
	text = cleanText(text)
	if len([]rune(text)) < 20 {
		return domain.Resume{}, ErrNoText
	}
	preview := []rune(text)
	if len(preview) > 360 {
		preview = preview[:360]
	}
	return domain.Resume{
		FileName:    originalName,
		ContentType: contentType,
		Text:        text,
		Preview:     string(preview),
		Keywords:    ExtractKeywords(text, 18),
		Characters:  len([]rune(text)),
	}, nil
}

func cleanText(text string) string {
	lines := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
	cleaned := make([]string, 0, len(lines))
	lastBlank := false
	for _, line := range lines {
		line = strings.Join(strings.Fields(line), " ")
		if line == "" {
			if !lastBlank {
				cleaned = append(cleaned, "")
			}
			lastBlank = true
			continue
		}
		lastBlank = false
		cleaned = append(cleaned, line)
	}
	return strings.TrimSpace(strings.Join(cleaned, "\n"))
}
