package ast

import (
	"interpreter/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{token.LET, "let"},
				Name: &Identifier{
					Token: token.Token{token.IDENT, "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{token.IDENT, "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}
	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. Got=%q", program.String())
	}
}
