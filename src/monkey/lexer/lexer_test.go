package lexer

import (
	"testing"

	"token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	let add = fn(x, y) {
		x + y;
	  };
	  
	  let result = add(five, ten);
	  !-/*5;
	  5 < 10 > 5;
	  
	  if (5 < 10) {
		  return true;
	  } else {
		  return false;
	  }
	  
	  10 == 10;
	  10 != 9;
	  `

	  tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	  } {
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
	  }

	  l := New(input)

	  for i, tt := range tests {
		tok := l.NextToken()
	  }
}
