package file_parser

import "testing"

func Test_parse_file(t *testing.T) {
	parser := NewParser()
	res, err := parser.Parse("../test-data/test")
	if err != nil {
		t.Error(err)
	}
	if len(res) != 3 {
		t.Error("parse file error")
	}
}
