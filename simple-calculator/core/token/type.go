package token

type Type int

const (
	Int              Type = iota
	Assignment            // '='
	Add                   // '+'
	Sub                   // '-'
	Mul                   // '*'
	Div                   // '/'
	Semi                  // semicolon ';'
	IntLiteral            // 0-9
	Id                    // 标识符
	LeftParenthesis       // '('
	RightParenthesis      // ')'
	UnKnown               // 未知
)

func (t *Type) String() string {
	switch *t {
	case Int:
		return "Int"
	case Assignment:
		return "Assignment"
	case Add, Sub, Mul, Div:
		return "operator"
	case Semi:
		return "EOF"
	case IntLiteral:
		return "IntLiteral"
	case Id:
		return "Id"
	case LeftParenthesis, RightParenthesis:
		return "Parentheses"
	default:
		return "UnKnown"
	}
}
