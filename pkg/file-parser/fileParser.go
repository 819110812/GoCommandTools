package file_parser

import (
	"bufio"
	"errors"
	"os"
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
