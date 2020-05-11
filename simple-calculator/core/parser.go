package core

import (
	"LearnCompile/simple-calculator/core/ast"
	"LearnCompile/simple-calculator/core/token"
	"fmt"
)

type Parser struct{}

func NewParser() *Parser {
	return new(Parser)
}

func (p *Parser) parse(script string) (root *ast.Node) {
	// 词法解析
	lexer := NewLexer()
	lexer.tokenize(script)
	// 语法解析
	root = ast.NewNode(ast.Program, "")
	var (
		child *ast.Node
	)
	tokens := lexer.Tokens()
	for tokens.Peek() != nil {
		child = p.declareSmt(tokens)
		if child != nil {
			root.AddChild(child)
			continue
		}

		child = p.assignSmt(tokens)
		if child != nil {
			root.AddChild(child)
			continue
		}

		child = p.statement(tokens)
		if child != nil {
			root.AddChild(child)
			continue
		}

		child = p.addictionExpr(tokens)
		if child != nil {
			root.AddChild(child)
			continue
		}
	}
	return
}

func (p *Parser) declareSmt(tokens *token.List) (node *ast.Node) {
	tok := tokens.Peek()
	if tok == nil {
		return
	}

	// int declare
	if tok.Type() == token.Int {
		tok = tokens.Next()
		if tok.Type() == token.Id {
			node = ast.NewNode(ast.IntDeclare, tok.Text())
		} else {
			fmt.Printf("%s\n", "excepting identifier bebind declare smt")
			return
		}
		tok = tokens.Next()
		if tok != nil && tok.Type() == token.Assignment {
			// 声明并初始化变量
			tokens.Next()
			child := p.addictionExpr(tokens)

			tok = tokens.Peek()
			if child != nil && tok != nil && tok.Type() == token.Semi {
				// 消耗分号(';')
				tokens.Next()
				node.AddChild(child)
			} else {
				node = nil // 回溯
				fmt.Printf("%s\n", "excepting ';' at end of smt")
			}
		} else if tok != nil && tok.Type() == token.Semi {
			tokens.Next() // 消耗分号(';')
		} else {
			node = nil // 回溯
			fmt.Printf("%s\n", "excepting ';' at end of smt")
		}
	}
	return
}

func (p *Parser) assignSmt(tokens *token.List) (node *ast.Node) {
	tok1 := tokens.Peek()
	if tok1 == nil {
		return
	}

	if tok1.Type() == token.Id {
		tok2 := tokens.Next()
		if tok2 != nil && tok2.Type() == token.Assignment {
			node = ast.NewNode(ast.Assignment, tok1.Text())
			tokens.Next()
			child := p.addictionExpr(tokens)
			if child == nil {
				fmt.Printf("%s\n", "excepting expression after '='")
			}
			tok2 = tokens.Peek()
			if tok2 != nil && tok2.Type() == token.Semi {
				node.AddChild(child)
				tokens.Next()
			} else {
				node = nil // 回溯
				fmt.Printf("%s\n", "excepting ';' at end of smt")
			}
		} else {
			tokens.Prev()
		}
	}
	return
}

func (p *Parser) statement(tokens *token.List) (node *ast.Node) {
	tok := tokens.Peek()
	if tok == nil {
		return
	}

	screen := tokens.Position() // 设置快照
	child := p.addictionExpr(tokens)
	tok = tokens.Peek()
	if tok != nil && tok.Type() == token.Semi {
		node = ast.NewNode(ast.Statement, "")
		if child != nil {
			node.AddChild(child)
		}
		tokens.Next()
		return
	}
	tokens.SetPosition(screen) // 还原
	return
}

func (p *Parser) addictionExpr(tokens *token.List) (node *ast.Node) {
	node = p.multiplicationExpr(tokens)
	child1 := node

	tok := tokens.Peek()
	if tok == nil {
		return
	}
	if tok.Type() == token.Add || tok.Type() == token.Sub {
		node = ast.NewNode(ast.Additive, tok.Text())
		tokens.Next()
		child2 := p.addictionExpr(tokens)
		if child2 != nil {
			node.AddChild(child1)
			node.AddChild(child2)
		} else {
			panic(fmt.Sprintf("%s '%s'", "excepting IntLiteral, Id or expression behind", tok.Text()))
		}
	}
	return
}

func (p *Parser) multiplicationExpr(tokens *token.List) (node *ast.Node) {
	child1 := p.primary(tokens)
	node = child1

	tok := tokens.Peek()
	if tok == nil {
		return
	}
	if tok.Type() == token.Mul || tok.Type() == token.Div {
		node = ast.NewNode(ast.Multiplicative, tok.Text())
		tokens.Next()
		child2 := p.multiplicationExpr(tokens)
		if child2 != nil {
			node.AddChild(child1)
			node.AddChild(child2)
		} else {
			panic(fmt.Sprintf("%s '%s'", "excepting IntLiteral or Id behind", tok.Text()))
		}
	}
	return
}

var HasLeftParenthesis = false

func (p *Parser) primary(tokens *token.List) (node *ast.Node) {
	tok := tokens.Peek()
	if tok == nil {
		return
	}

	switch tok.Type() {
	case token.IntLiteral:
		node = ast.NewNode(ast.IntLiteral, tok.Text())
	case token.Id:
		node = ast.NewNode(ast.Identifier, tok.Text())
	case token.LeftParenthesis:
		HasLeftParenthesis = true
		tokens.Next() // 消耗左括号
		node = p.addictionExpr(tokens)
		if node == nil {
			panic(fmt.Sprintf("%s", "excepted expression within parenthese"))
		}
		tok = tokens.Peek()
		if tok == nil || tok.Type() != token.RightParenthesis {
			panic(fmt.Sprintf("%s", "excepted right parenthesis"))
		}
	case token.RightParenthesis:
		if HasLeftParenthesis {
			HasLeftParenthesis = false
		} else {
			panic(fmt.Sprintf("%s", "excepted left parenthesis"))
		}
	}

	tokens.Next()
	return
}
