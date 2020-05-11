package ast

type Type int

const (
	Program Type = iota

	Statement
	IntDeclare
	Assignment

	Additive
	Multiplicative
	Primary

	Identifier
	IntLiteral
)

func (t *Type) String() string {
	switch *t {
	case Program:
		return "Program"
	case Statement:
		return "Statement"
	case IntDeclare:
		return "IntDeclare"
	case Assignment:
		return "Assignment"
	case Additive:
		return "Additive"
	case Multiplicative:
		return "Multiplicative"
	case Primary:
		return "Primary"
	case Identifier:
		return "Identifier"
	case IntLiteral:
		return "IntLiteral"
	default:
		return "Unknown"
	}
}
