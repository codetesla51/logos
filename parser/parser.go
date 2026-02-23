package parser

import (
	"fmt"
	"strings"

	"github.com/codetesla51/golexer/golexer"
)

const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

type Node interface {
	TokenLiteral() string
	String() string
}
type Statement interface {
	Node
	statmentNode()
}
type Expression interface {
	Node
	expressionNode()
}
type Program struct {
	Statements []Statement
}
type Identifier struct {
	Token golexer.Token
	Value string
}
type IntegerLiteral struct {
	Token golexer.Token
	Value int64
}
type InfixExpression struct {
	Token    golexer.Token
	Left     Expression
	Operator string
	Right    Expression
}
type LetStatement struct {
	Token golexer.Token
	Name  *Identifier
	Value Expression
}
type ReturnStatement struct {
	Token       golexer.Token
	ReturnValue Expression
}
type ExpressionStatement struct {
	Token      golexer.Token
	Expression Expression
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}
func (p *Program) String() string {
	var out strings.Builder
	fmt.Println("all statements")
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}
func (i *Identifier) expressionNode()      {}
func (i *Identifier) statmentNode()        {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	return fmt.Sprintf("(%s %s %s)", ie.Left.String(), ie.Operator, ie.Right.String())
}

func (ls *LetStatement) statmentNode()        {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var out strings.Builder
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

func (rs *ReturnStatement) statmentNode()        {}
func (rs *ReturnStatement) expressionNode()      {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out strings.Builder
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}
func (es *ExpressionStatement) statmentNode()        {}
func (es *ExpressionStatement) expressionNode()      {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type Parser struct {
	lexer  *golexer.Lexer
	errors []string

	curToken  golexer.Token
	peekToken golexer.Token
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}
func (p *Parser) curTokenIs(t golexer.TokenType) bool {
	return p.curToken.Type == t
}
func (p *Parser) peekTokenIs(t golexer.TokenType) bool {
	return p.peekToken.Type == t
}
func (p *Parser) expectPeek(t golexer.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}
func (p *Parser) peekError(t golexer.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
func (p *Parser) Errors() []string {
	return p.errors
}

func NewParser(lexer *golexer.Lexer) *Parser {
	p := &Parser{
		lexer:  lexer,
		errors: []string{},
	}
	p.nextToken()
	p.nextToken()
	return p
}
func (p *Parser) Parse() *Program {
	program := &Program{}
	program.Statements = []Statement{}
	for !p.curTokenIs(golexer.EOF) {
		stmt := p.parseStatment()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}
func (p *Parser) parseStatment() Statement {
	switch p.curToken.Type {
	case golexer.LET:
		return p.parseLteStatment()
	case golexer.RETURN:
		return p.parseReturnStatment()
	default:
		return p.parseExpressionStatment()
	}

}
