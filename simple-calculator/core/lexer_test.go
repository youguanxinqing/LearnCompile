package core

import "testing"

func TestLexer_tokenize(t *testing.T) {
	code := "int age=10;age = age*10 +16; age / int ;"
	lexer := NewLexer()

	lexer.tokenize(code)
	t.Log(lexer.tokens)
}

