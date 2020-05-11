package token

import (
	"testing"
)


func TestList_PeekAndNext(t *testing.T) {
	list := NewList()
	tokens := []*Token{
		NewToken(Int, "int"),
		NewToken(Add, "+"),
		NewToken(IntLiteral, "10"),
		NewToken(Id, "age"),
	}
	list.Push(tokens...)
	tmp := list.PeekAndNext()
	t.Log(tmp)

}

func TestList_String(t *testing.T) {
	list := NewList()
	tokens := []*Token{
		NewToken(Int, "int"),
		NewToken(Add, "+"),
		NewToken(IntLiteral, "10"),
		NewToken(Id, "age"),
	}
	list.Push(tokens...)
	t.Log(list)
}