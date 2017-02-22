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
	"fmt"
	"reflect"
	"unicode"
	"compiler/transformer"
)

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

func ExpectName(op int16) bool {
	return expectName[GetType(op)]
}

func ExpectArgNum(op int16) int {
	return expectArgNum[op]
}

func ExpectArgs(op int16) bool {
	return expectArgs[GetType(op)]
}

func FinalArgNum(op int16) int {
	return finalArgNum[op]
}

func NeedArgNum(op int16) bool {
	return needArgNum[GetType(op)]
}

func NewNumberValue(x int16) Value {
	return Value{INTEGER, "", x, nil}
}

func NewNumberValues(args ...int16) []Value {
	ret := make([]Value, len(args))
	for i, v := range args {
		ret[i] = NewNumberValue(v)
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

func GetCommand(name string) (int16,bool) {
	ret,ok := operationNameMap[name]
	return ret,ok
}

func GetName(op int16) string {
	if int(op) >= len(operationNames) {
		return "undefined"
	}
	return operationNames[op]
}

func GetType(op int16) int16 {
	if int(op) >= len(operationTypes) {
		return UNDEFINED
	}
	return operationTypes[op]
}

func GetOperationNum() int {
	return len(operationTypes)
}

func ValidName(name string) bool {
	if len(name) == 0 {
		return false
	}
	for i, c := range name {
		if !unicode.IsLetter(c) {
			if i == 0 {
				return false
			}
			if !unicode.IsDigit(c) && c != '-' && c != '.' && c != '_' {
				return false
			}
		}
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
	fmt.Printf("%s", GetName(op.Command))
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
	ret := fmt.Sprintf("%s", GetName(op.Command))
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
