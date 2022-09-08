package patterns

import (
	"fmt"
	"strconv"
)

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func Parse(s string) (pattern []Token, err error) {
	t := NewTokenizer(s)

	var last Token
	for c := t.Char; c != 0; c = t.Next() {
		var tt Token
		// fmt.Printf("%c\n", c)

		switch c {
		case '?':
			last.SetQuantifier(Quantifier{Max: 1})
		case '*':
			last.SetQuantifier(Quantifier{Max: -1})
		case '+':
			last.SetQuantifier(Quantifier{Min: 1, Max: -1})
		case '{':
			var digits []byte
			min := -1

		QuantifierLoop:
			for {
				c = t.Next()

				switch {
				case isDigit(c):
					digits = append(digits, c)
				case c == ',':
					min, err = strconv.Atoi(string(digits))
					if err != nil {
						return pattern, err
					}
					digits = []byte{} // clear
				case c == '}':
					max, err := strconv.Atoi(string(digits))
					if err != nil {
						return pattern, err
					}

					if min == -1 {
						min = max
					}

					last.SetQuantifier(Quantifier{min, max})
					break QuantifierLoop
				case c == 0:
					return pattern, fmt.Errorf("unmatched quantifier range")
				default:
					return pattern, fmt.Errorf("invalid character '%c' in quantifier", c)
				}
			}
		case '[':
			cl := &TokenClass{}

			if t.Peek() == '^' {
				t.Next() // consume caret
				cl.Negated = true
			}

			for {
				c = t.Next()

				if c == 0 {
					err = fmt.Errorf("unmatched class")
					return pattern, err
				}

				if c == ']' {
					tt = cl
					break
				}

				if t.Peek() == '-' {
					t.Next() // consume dash
					cl.Parts = append(cl.Parts, &ClassRange{Start: c, End: t.Next()})
				} else {
					cl.Parts = append(cl.Parts, &TokenLiteral{BaseToken{Syntax: string(c)}})
				}
			}

		case '.':
			tt = &TokenDot{}
		default:
			tt = &TokenLiteral{BaseToken{Syntax: string(c)}}
		}

		if tt != nil {
			last = tt
			pattern = append(pattern, tt)
		}

	}

	return
}
