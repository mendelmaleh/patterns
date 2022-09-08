package patterns

import (
	"math/rand"
)

const (
	AsciiStart = 33
	AsciiEnd   = 126
)

// AsciiRandom returns a random ASCII character (bytes 33-126)
func AsciiRandom(r *rand.Rand) byte {
	return byte(AsciiStart + r.Intn(AsciiEnd-AsciiStart))
}

type Token interface {
	String() string
	SetQuantifier(q Quantifier)

	// Quantity returns the value from the quantifier
	Quantity(r *rand.Rand) int

	// Value for the token, not quantified
	Value(r *rand.Rand) string
}

// BaseToken contains common fields that all tokens should have and some basic methods they can inherit
type BaseToken struct {
	Quantifier

	Syntax string
}

func (t *BaseToken) String() string { return t.Syntax }

func (t *BaseToken) SetQuantifier(q Quantifier) { t.Quantifier = q }

// Quantifier is a quantifier with a min-max range from which to pick a random value
type Quantifier struct {
	Min, Max int
}

// Quantity resolves the quantifier and implements the Token.Quantity interface
func (q Quantifier) Quantity(r *rand.Rand) int {
	// default value
	if q == (Quantifier{}) {
		return 1
	}

	// no range
	if q.Min == q.Max {
		return q.Min
	}

	return q.Min + r.Intn(1+q.Max-q.Min)
}

// TokenLiteral represents a literal character
type TokenLiteral struct {
	BaseToken
}

var _ Token = &TokenLiteral{BaseToken{Syntax: "a"}}

func (t TokenLiteral) Value(_ *rand.Rand) string { return string(t.Syntax) }

func (t TokenLiteral) IsValid(b byte) bool { return t.Syntax == string(b) }

// TokenDot
type TokenDot struct {
	BaseToken
}

func (t TokenDot) Value(r *rand.Rand) string {
	return string(AsciiRandom(r))
}
