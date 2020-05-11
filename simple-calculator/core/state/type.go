package state

type Type int

const (
	Init Type = iota
	Identifier
	IntLiteral
	Int1  // int[1-3]  int 关键字判定过程
	Int2
	Int3
)

