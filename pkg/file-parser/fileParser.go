package file_parser

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type FileParser interface {
	Parse(filePath string) ([]string, error)
}

type Parser struct {
}

func (p *Parser) Parse(filePath string) ([]string, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		return nil, errors.New("file not exist")
	}
	bytesContent, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("read file error")
	}

	r := bufio.NewReader(bytesContent)
	var chunks []string
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}
		chunks = append(chunks, string(line))
	}

	return chunks, nil
}

func NewParser() *Parser {
	return &Parser{}
}

func isTxtFile(filePath string) bool {
	if strings.HasSuffix(filePath, ".txt") {
		return true
	}
	return false
}

func isJsonFile(filePath string) bool {
	if strings.HasSuffix(filePath, ".json") {
		return true
	}
	return false
}

func isYamlFile(filePath string) bool {
	if strings.HasSuffix(filePath, ".yaml") {
		return true
	}
	return false
}
