package move_goyacc

import (
	"bytes"
)

const eof = 0

type exprLex struct {
	line []byte
	pos  int
	st   bool
}

func (x *exprLex) Lex(yylval *exprSymType) int {
	var buff bytes.Buffer
	if len(x.line) == 0 {
		return eof
	}
	for {
		c := x.next()
		switch c {
		case eof:
			return eof
		case ':':
			x.st = true
			return int(c)
		case ';':
			return int(c)
		case '#':
			x.pos = len(x.line)
			return int(c)
		case '-':
			return int(c)
		default:
			buff.WriteRune(c)
			if buff.String() == "Print" && x.peek() == ':' {

			}
		}
	}
}

func (x *exprLex) next() rune {
	if x.pos == len(x.line) {
		return eof
	}
	defer func() {
		x.pos++
	}()
	return rune(x.line[x.pos])
}

func()
