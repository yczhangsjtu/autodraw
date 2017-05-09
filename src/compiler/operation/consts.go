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

import (
	"regexp"
)

// Operations
const (
	UNDEFINED int16 = iota
	LINE
	RECT
	OVAL
	POLYGON
	TEXT
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
	STRING
	NAN
)

var operationNames = []string{
	"undefined", "line", "rect", "oval", "polygon", "text", "set", "use",
	"push", "pop", "transform", "rotate", "scale", "translate", "draw", "import",
	"begin", "end",
}

var operationPatterns = []string{
	"^\\s*undefined\\s*$",
	"^\\s*line((?:\\s+(?:-?\\d+|[_A-Za-z]\\w*)){4})\\s*$",
	"^\\s*rect((?:\\s+(?:-?\\d+|[_A-Za-z]\\w*)){4})\\s*$",
	"^\\s*oval((?:\\s+(?:-?\\d+|[_A-Za-z]\\w*)){4})\\s*$",
	"^\\s*polygon(((?:\\s+(?:-?\\d+|[_A-Za-z]\\w*)){2})+)\\s*$",
	"^\\s*text((?:\\s+(?:-?\\d+|[_A-Za-z]\\w*)){3})\\s+(\"[^\"]*\"|[_A-Za-z]\\w*)\\s*$",
	"^\\s*set\\s+([_A-Za-z]\\w*)\\s+(-?\\d+|\"[^\"]*\"|[_A-Za-z]\\w*)\\s*$",
	"^\\s*use\\s+([_A-Za-z]\\w*)\\s*$",
	"^\\s*push\\s+([_A-Za-z]\\w*)\\s*$",
	"^\\s*pop\\s*$",
	"^\\s*transform\\s+([_A-Za-z]\\w*)((?:\\s+(?:-?\\d+|[_A-Za-z]\\w*)){6})\\s*$",
	"^\\s*rotate\\s+([_A-Za-z]\\w*)\\s+(-?\\d+|[_A-Za-z]\\w*)\\s*$",
	"^\\s*scale\\s+([_A-Za-z]\\w*)((?:\\s+(?:-?\\d+|[_A-Za-z]\\w*)){2})\\s*$",
	"^\\s*translate\\s+([_A-Za-z]\\w*)((?:\\s+(?:-?\\d+|[_A-Za-z]\\w*)){2})\\s*$",
	"^\\s*draw\\s+([_A-Za-z]\\w*)\\s*$",
	"^\\s*import\\s+([!-~]+)\\s*$",
	"^\\s*begin\\s+([_A-Za-z]\\w*)\\s*$",
	"^\\s*end\\s*$",
}

var variablePattern string = "^[_A-Za-z]\\w*$"
var variableRegexp *regexp.Regexp = nil
var variableFinderPattern string = "-?\\d+|[_A-Za-z]\\w*"
var variableFinderRegexp *regexp.Regexp = nil

var operationRegexp = []*regexp.Regexp {
	nil,nil,nil,nil,nil,
	nil,nil,nil,nil,nil,
	nil,nil,nil,nil,nil,
	nil,nil,nil,
}

var operationNameMap = map[string]int16{
	"undefined": UNDEFINED, "line": LINE, "rect": RECT,
	"oval": OVAL, "polygon": POLYGON, "set": SET, "use": USE, "push": PUSH,
	"pop": POP, "transform": TRANSFORM, "rotate": ROTATE, "scale": SCALE,
	"translate": TRANSLATE,"draw": DRAW, "import": IMPORT, "begin": BEGIN,
	"end": END,
}
