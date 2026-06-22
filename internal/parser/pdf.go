package parser

import (
	"fmt"
	"io"

	pdf "github.com/ledongthuc/pdf"
)

func parsePDF(path string) (string, error) {
	file, reader, err := pdf.Open(path)
	if err != nil {
		return "", fmt.Errorf("打开 PDF 失败: %w", err)
	}
	defer file.Close()
	plain, err := reader.GetPlainText()
	if err != nil {
		return "", fmt.Errorf("读取 PDF 文本失败: %w", err)
	}
	content, err := io.ReadAll(plain)
	if err != nil {
		return "", fmt.Errorf("读取 PDF 文本失败: %w", err)
	}
	return string(content), nil
}
