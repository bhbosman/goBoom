package builder

import (
	"go/token"
)

type Field struct {
	Location
}

func NewField(indent int, position token.Position, pos token.Pos, end token.Pos) Field {
	return Field{
		Location: NewLocation(indent, position, pos, end),
	}
}
