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

import "os"
import "fmt"
import "reflect"
import "unicode"
import "compiler/transformer"

// Operations
const (
	UNDEFINED int16 = iota
	LINE
	RECT
	CIRCLE
	OVAL
	POLYGON
	SET
	USE
	PUSH
	POP
	TRANSFORM
	ROTATE
	SCALE
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

const Version string = "1.0"

var OperationNames = []string{
	"undefined", "line", "rect", "circle", "oval", "polygon", "set", "use",
	"push", "pop", "transform", "rotate", "scale", "draw", "import",
	"begin", "end",
}

var OperationTypes = []int16{
	NOT_OPERATION, DRAW_FIXED, DRAW_FIXED, DRAW_FIXED, DRAW_FIXED, DRAW_UNDETERMINED,
	ASSIGN, STATE, STATE, SINGLE, ASSIGN, ASSIGN, ASSIGN,
	STATE, STATE, STATE, SINGLE,
}

var expectName = []bool{
	false, false, false, true, false, true,
}

var expectArgNum = []int16{
	0, 4, 4, 3, 5, 0, 1, 0,
	0, 0, 6, 1, 2, 0, 0,
	0, 0,
}

var expectArgs = []bool{
	false, true, true, true, false, false,
}

var needArgNum = []bool{
	false, false, true, false, false, false,
}

var finalArgNum = []int16{
	0, 4, 8, 3, 5, 0, 1, 0,
	0, 0, 6, 1, 2, 0, 0,
	0, 0,
}

var OperationNameMap = map[string]int16{
	"undefined": UNDEFINED, "line": LINE, "rect": RECT, "circle": CIRCLE,
	"oval": OVAL, "polygon": POLYGON, "set": SET, "use": USE, "push": PUSH,
	"pop": POP, "transform": TRANSFORM, "rotate": ROTATE, "scale": SCALE,
	"draw": DRAW, "import": IMPORT, "begin": BEGIN, "end": END,
}

type Value struct {
	Type      int16
	Name      string
	Number    int16
	Transform *transformer.Transform
}

type Operation struct {
	Command int16
	Name    string
	Args    []Value
}

var Verbose bool = false

func ExpectName(op int16) bool {
	return expectName[OperationTypes[op]]
}

func ExpectArgNum(op int16) int16 {
	return expectArgNum[op]
}

func ExpectArgs(op int16) bool {
	return expectArgs[OperationTypes[op]]
}

func FinalArgNum(op int16) int16 {
	return finalArgNum[op]
}

func NeedArgNum(op int16) bool {
	return needArgNum[OperationTypes[op]]
}

func NewNumberValue(x int16) Value {
	return Value{INTEGER, "", x, nil}
}

func NewNumberValues(args ...int16) []Value {
	ret := make([]Value, len(args))
	for i, _ := range ret {
		ret[i] = Value{INTEGER, "", args[i], nil}
	}
	return ret
}

func NewVariableValue(name string) Value {
	if ValidName(name) {
		return Value{VARIABLE, name, 0, nil}
	}
	return Value{NAN, "", 0, nil}
}

func NewTransformValue(tf *transformer.Transform) Value {
	return Value{TRANSFORMER, "", 0, tf}
}

func NewOperation(op int16) Operation {
	operation := Operation{}
	operation.Command = op
	return operation
}

func ValidName(name string) bool {
	if len(name) == 0 {
		return false
	}
	for i, c := range name {
		if i == 0 {
			if !unicode.IsLetter(c) {
				return false
			}
		} else {
			if !unicode.IsLetter(c) && !unicode.IsDigit(c) && c != '-' &&
				c != '.' && c != '_' {
				return false
			}
		}
	}
	return true
}

func ValidPath(path string) bool {
	if len(path) == 0 {
		return false
	}
	file, err := os.Open(path)
	if err != nil {
		return false
	} else {
		file.Close()
	}
	return true
}

func (v *Value) Print() {
	switch v.Type {
	case INTEGER:
		fmt.Printf("%d", v.Number)
	case TRANSFORMER:
		v.Transform.Print()
	case VARIABLE:
		fmt.Printf("%s", v.Name)
	default:
		fmt.Printf("undefined")
	}
}

func ValuesPrint(vs []Value) {
	fmt.Printf("[")
	for i, v := range vs {
		if i > 0 {
			fmt.Printf(",")
		}
		v.Print()
	}
	fmt.Printf("]")
}

func (op *Operation) Print() {
	fmt.Printf("%s", OperationNames[op.Command])
	if op.Name != "" {
		fmt.Printf(" %s", op.Name)
	}
	if len(op.Args) > 0 {
		fmt.Printf(" ")
		ValuesPrint(op.Args)
	}
}

func (v *Value) ToString() string {
	switch v.Type {
	case INTEGER:
		return fmt.Sprintf("%d", v.Number)
	case TRANSFORMER:
		return v.Transform.ToString()
	case VARIABLE:
		return fmt.Sprintf("%s", v.Name)
	}
	return fmt.Sprintf("undefined")
}

func ValuesToString(vs []Value) string {
	ret := "["
	for i, v := range vs {
		if i > 0 {
			ret += ","
		}
		ret += v.ToString()
	}
	ret += "]"
	return ret
}

func (op *Operation) ToString() string {
	ret := fmt.Sprintf("%s", OperationNames[op.Command])
	if op.Name != "" {
		ret += fmt.Sprintf(" %s", op.Name)
	}
	if len(op.Args) > 0 {
		ret += " "
		ret += ValuesToString(op.Args)
	}
	return ret
}

func (op *Operation) Equal(op2 Operation) bool {
	return reflect.DeepEqual(*op, op2)
}

func ValuesToInt(values []Value) ([]int16, bool) {
	ret := make([]int16, len(values))
	for i, v := range values {
		if v.Type != INTEGER {
			return ret, false
		}
		ret[i] = v.Number
	}
	return ret, true
}
