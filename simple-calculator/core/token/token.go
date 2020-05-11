package token

import (
	"fmt"
	"strings"
)

type Token struct {
	_type Type
	_text string
}

func (t *Token) String() string {
	return fmt.Sprintf("%s:%s", t._type.String(), t._text)
}

func NewToken(_type Type, _text string) *Token {
	return &Token{
		_type: _type,
		_text: _text,
	}
}

func (t *Token) Type() Type {
	return t._type
}

func (t *Token) Text() string {
	return t._text
}

type List struct {
	list []*Token
	cur int
	len int
}

func NewList() *List {
	return &List{
		list: make([]*Token, 0),
		cur: 0,
		len: 0,
	}
}

func (l *List) Push(token ...*Token) {
	l.list = append(l.list, token...)
	l.len += len(token)
}

// Peek implements 读当前位置的 token, 只读不消耗
func (l *List) Peek() *Token {
	if l.cur < l.len {
		return l.list[l.cur]
	}
	return nil
}

// Next implements 跳至下一个 token, 并返回该 token
func (l *List) Next() *Token {
	if l.cur < l.len {
		l.cur++
	}

	if l.cur < l.len {
		return l.list[l.cur]
	}
	return nil
}

func (l *List) Prev() *Token {
	if l.cur > 0 {
		l.cur--
		return l.list[l.cur]
	}
	return nil
}

type Pos int  // 快照
func (l *List) Position() Pos { return Pos(l.cur) }
func (l *List) SetPosition(p Pos) {  l.cur = int(p) }

// PeekAndNext implements 读当前 token, 跳到下个 token
func (l *List) PeekAndNext() *Token {
	token := l.Peek()
	l.Next()
	return token
}

func (l *List) ToSlice() []*Token {
	return l.list
}

func (l *List) String() string {
	str := new(strings.Builder)
	str.Write([]byte{'[', ' '})
	for l.Peek() != nil {
		str.WriteString(l.Peek().String() + " ")
		l.Next()
	}
	
	str.WriteByte(']')
	return str.String()
}

