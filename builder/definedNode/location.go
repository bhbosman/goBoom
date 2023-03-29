package definedNode

import (
	"fmt"
	"go/token"
	"strings"
)

type Location struct {
	indent   int
	position token.Position

	pos token.Pos
	end token.Pos
}

func (self *Location) Indent() int {
	return self.indent
}

func (self *Location) Position() token.Position {
	return self.position
}

func NewLocation(indent int, position token.Position, pos token.Pos, end token.Pos) Location {
	return Location{

		indent:   indent,
		position: position,
		pos:      pos,
		end:      end,
	}
}
func NewLocationP(indent int, position token.Position, pos token.Pos, end token.Pos) *Location {
	return &Location{

		indent:   indent,
		position: position,
		pos:      pos,
		end:      end,
	}
}

func (self *Location) Print(s string) {
	value := fmt.Sprintf("%v:%v:%v:%v%v", self.position.Filename, self.position.Line, self.position.Column, strings.Repeat("\t", self.indent), s)
	println(value)
}
func (self *Location) Pos() token.Pos {
	return self.pos
}

func (self *Location) End() token.Pos {
	return self.end
}
