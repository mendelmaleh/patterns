package patterns

type Tokenizer struct {
	String string

	Pos  int
	Char byte
}

func NewTokenizer(s string) (t Tokenizer) {
	t.String = s
	t.Char = t.String[0]

	return
}

func (t Tokenizer) Peek() byte {
	if t.Pos+1 >= len(t.String) {
		return 0
	}

	return t.String[t.Pos+1]
}

// Next returns the next character or 0 if the end was reached
func (t *Tokenizer) Next() byte {
	t.Pos += 1

	if t.Pos == len(t.String) {
		return 0
	}

	t.Char = t.String[t.Pos]
	return t.Char
}
