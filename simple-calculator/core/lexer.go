package core

import (
	state "LearnCompile/simple-calculator/core/state"
	"LearnCompile/simple-calculator/core/token"
	"fmt"
	"io"
	"strings"
)

type Lexer struct {
	tokens    *token.List
	state     state.Type
	tokenType token.Type
	tokenText []byte
}

func NewLexer() *Lexer {
	return &Lexer{
		tokens:    token.NewList(),
		state:     state.Init,
		tokenType: token.UnKnown,
		tokenText: make([]byte, 0),
	}
}

func (l *Lexer) Tokens() *token.List {
	return l.tokens
}

func isNumber(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isOperator(ch byte) bool {
	switch ch {
	case '+', '-', '*', '/', '=':
		return true
	default:
		return false
	}
}

func isSemi(ch byte) bool {
	return ch == ';'
}

func isAlpha(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func (l *Lexer) tokTextAppend(ch byte) {
	l.tokenText = append(l.tokenText, ch)
}

func (l *Lexer) tokenize(script string) {
	stream := strings.NewReader(script)
	for {
		ch, err := stream.ReadByte()
		if err != nil || err == io.EOF {
			l.initToken(' ') // 收尾
			return
		} else if err != nil {
			fmt.Printf("tokenize error, reason: %s\n", err)
		}

		switch l.state {
		case state.Init:
			l.initToken(ch)
		case state.IntLiteral:
			if isNumber(ch) {
				l.tokTextAppend(ch)
			} else {
				l.initToken(ch)
			}
		case state.Identifier:
			if isNumber(ch) || isAlpha(ch) {
				l.tokTextAppend(ch)
			} else {
				l.initToken(ch)
			}
		case state.Int1:
			if isAlpha(ch) || isNumber(ch) {
				if ch == 'n' {
					l.state = state.Int2
				} else {
					l.state = state.Identifier
				}
				l.tokTextAppend(ch)
			} else {
				l.initToken(ch)
			}
		case state.Int2:
			if isAlpha(ch) || isNumber(ch) {
				if ch == 't' {
					l.state = state.Int3
				} else {
					l.state = state.Identifier
				}
				l.tokTextAppend(ch)
			} else {
				l.initToken(ch)
			}
		case state.Int3:
			if isAlpha(ch) && isNumber(ch) {
				l.state = state.Identifier
				l.tokTextAppend(ch)
			} else {
				l.tokenType = token.Int
				l.initToken(ch)
			}
		}
	}
}

func (l *Lexer) initToken(ch byte) {
	if len(l.tokenText) > 0 {
		// 结束上一个 token
		tok := token.NewToken(l.tokenType, string(l.tokenText))
		l.tokens.Push(tok)
		// 开始下一个 token
		l.state = state.Init
		l.tokenText = make([]byte, 0)
	}

	// token 首个字符的处理方式
	if isNumber(ch) {
		l.tokenType = token.IntLiteral
		l.tokTextAppend(ch)
		l.state = state.IntLiteral
	} else if isAlpha(ch) {
		if ch == 'i' {
			l.state = state.Int1
		} else {
			l.state = state.Identifier
		}
		l.tokenType = token.Id
		l.tokTextAppend(ch)
	} else if isOperator(ch) {
		switch ch {
		case '+':
			l.tokenType = token.Add
		case '-':
			l.tokenType = token.Sub
		case '*':
			l.tokenType = token.Mul
		case '/':
			l.tokenType = token.Div
		case '=':
			l.tokenType = token.Assignment
		}
		l.state = state.Init
		l.tokTextAppend(ch)
	} else if isSemi(ch) {
		l.tokenType = token.Semi
		l.state = state.Init
		l.tokTextAppend(ch)
	}
}
