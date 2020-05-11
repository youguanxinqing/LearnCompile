package core

import "testing"

func TestParser_parse(t *testing.T) {
	parser := NewParser()
	root := parser.parse("int a=1 + 1; a = a + 10; 10")
	t.Log(root)
}