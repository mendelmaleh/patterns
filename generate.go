package patterns

import (
	"math/rand"
	"strings"
)

type Generator struct {
	Pattern string
	Tokens  []Token

	// Max for quantifiers like '*' and '+'
	Max  int
	Rand *rand.Rand
}

// NewGenerator returns a string generator from pattern
func NewGenerator(pattern string, r *rand.Rand) (*Generator, error) {
	tokens, err := Parse(pattern)
	if err != nil {
		return nil, err
	}

	return &Generator{Pattern: pattern, Tokens: tokens, Max: 128, Rand: r}, nil
}

// Generate a string from the pattern provided to the generator
func (g *Generator) Generate() string {
	var b strings.Builder
	for _, t := range g.Tokens {
		b.WriteString(Quantified(t, g.Max, g.Rand))
	}

	return b.String()
}

// Quantified returns a token's value quantified
func Quantified(t Token, max int, r *rand.Rand) string {
	// Can't be inherited from BaseToken because it relies on the actual Token's implementation Value()
	var b strings.Builder

	for i := 0; i < t.Quantity(max, r); i++ {
		b.WriteString(t.Value(r))
	}

	return b.String()
}
