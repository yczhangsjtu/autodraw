// This file is part of autodraw.
//
// Autodraw is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Autodraw is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with autodraw.  If not, see <http://www.gnu.org/licenses/>.
package operation

// Operations
const (
	UNDEFINED int16 = iota
	LINE
	RECT
	OVAL
	POLYGON
	SET
	USE
	PUSH
	POP
	TRANSFORM
	ROTATE
	SCALE
	TRANSLATE
	DRAW
	IMPORT
	BEGIN
	END
)

// Value types
const (
	VARIABLE int16 = iota
	INTEGER
	TRANSFORMER
	NAN
)

// Operation types
const (
	NOT_OPERATION int16 = iota
	DRAW_FIXED
	DRAW_UNDETERMINED
	ASSIGN
	SINGLE
	STATE
)

// Consts for parsers
const (
	INVALID int16 = iota
	COMMAND
	NAME
	NUMBER
)

const (
	NEED_COMMAND int16 = iota
	NEED_NAME
	NEED_VALUE
	FINISH
	ERROR
)

var operationNames = []string{
	"undefined", "line", "rect", "oval", "polygon", "set", "use",
	"push", "pop", "transform", "rotate", "scale", "translate", "draw", "import",
	"begin", "end",
}

var operationTypes = []int16{
	NOT_OPERATION, DRAW_FIXED, DRAW_FIXED, DRAW_FIXED, DRAW_UNDETERMINED,
	ASSIGN, STATE, STATE, SINGLE, ASSIGN, ASSIGN, ASSIGN, ASSIGN, STATE, STATE,
	STATE, SINGLE,
}

var expectName = []bool{
	false, false, false, true, false, true,
}

var expectArgNum = []int{
	0, 4, 4, 4, 0, 1, 0,
	0, 0, 6, 1, 2, 2, 0, 0,
	0, 0,
}

var expectArgs = []bool{
	false, true, true, true, false, false,
}

var needArgNum = []bool{
	false, false, true, false, false, false,
}

var finalArgNum = []int{
	0, 4, 8,16, 0, 1, 0,
	0, 0, 6, 1, 2, 2, 0, 0,
	0, 0,
}

var operationNameMap = map[string]int16{
	"undefined": UNDEFINED, "line": LINE, "rect": RECT,
	"oval": OVAL, "polygon": POLYGON, "set": SET, "use": USE, "push": PUSH,
	"pop": POP, "transform": TRANSFORM, "rotate": ROTATE, "scale": SCALE,
	"translate": TRANSLATE,"draw": DRAW, "import": IMPORT, "begin": BEGIN,
	"end": END,
}
