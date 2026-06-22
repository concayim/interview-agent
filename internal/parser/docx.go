package parser

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

func parseDOCX(path string) (string, error) {
	zr, err := zip.OpenReader(path)
	if err != nil {
		return "", fmt.Errorf("打开 DOCX 失败: %w", err)
	}
	defer zr.Close()
	for _, file := range zr.File {
		if file.Name != "word/document.xml" {
			continue
		}
		r, err := file.Open()
		if err != nil {
			return "", err
		}
		defer r.Close()
		return decodeWordXML(r)
	}
	return "", errorsNewDOCX()
}

func decodeWordXML(r io.Reader) (string, error) {
	decoder := xml.NewDecoder(r)
	var out bytes.Buffer
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", fmt.Errorf("解析 DOCX XML 失败: %w", err)
		}
		switch value := token.(type) {
		case xml.StartElement:
			if value.Name.Local == "tab" {
				out.WriteByte('\t')
			}
			if value.Name.Local == "br" {
				out.WriteByte('\n')
			}
		case xml.CharData:
			out.Write(value)
		case xml.EndElement:
			if value.Name.Local == "p" {
				out.WriteByte('\n')
			}
		}
	}
	return out.String(), nil
}

func errorsNewDOCX() error {
	return fmt.Errorf("无效的 DOCX：缺少 %s", strings.TrimSpace("word/document.xml"))
}
