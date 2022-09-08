package patterns

import "math/rand"

// TokenClass represents a character class
type TokenClass struct {
	BaseToken

	// Parts is a list of literals and class ranges
	Parts []Part

	// Negated if the class starts with a '^'
	Negated bool
}

var _ Token = &TokenClass{}

func (t TokenClass) Value(r *rand.Rand) string {
Random:
	for {
		b := AsciiRandom(r)

		for _, p := range t.Parts {
			if p.IsValid(b) {
				if !t.Negated {
					return string(b)
				} else {
					// new random char
					continue Random
				}
			}
		}

		if t.Negated {
			return string(b)
		}
	}
}

// Part is a part of a class
type Part interface {
	// IsValid checks if b would satisfy part
	IsValid(b byte) bool

	// value, might be random
	Value(r *rand.Rand) string
}

var (
	_ Part = &TokenLiteral{}
	_ Part = &ClassRange{}
)

// ClassRange represents a character class range
type ClassRange struct {
	BaseToken

	Start, End byte
}

var _ Token = &ClassRange{BaseToken{Syntax: "a-f"}, 'a', 'f'}

func (t ClassRange) Value(r *rand.Rand) string {
	s, e := int(t.Start), int(t.End)
	return string(s + r.Intn(e-s))
}

func (t ClassRange) IsValid(b byte) bool {
	return t.Start <= b && b <= t.End
}
