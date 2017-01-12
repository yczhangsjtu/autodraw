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

import "fmt"
import "reflect"
import "unicode"

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
)

const Version string = "1.0"

var OperationNames = []string {
	"undefined","line","rect","circle","oval","polygon","set","use",
	"push","pop","transform",
}

var OperationNameMap = map[string] int16 {
	"undefined":UNDEFINED, "line":LINE, "rect":RECT, "circle":CIRCLE,
	"oval":OVAL, "polygon":POLYGON, "set":SET, "use":USE, "push":PUSH,
	"pop":POP, "transform":TRANSFORM,
}

type Instruction struct {
	Op int16
	Args []int16
}

type Transform struct {
	a,b,c,d,x,y int16
}

type Operation struct {
	Op int16
	Name string
	Transform Transform
	Detail Instruction
}

var Verbose bool = false

func f2i(f float64) int16 {
	if f > 0.0 {
		return int16(f*100.0+0.5)
	} else {
		return int16(f*100.0-0.5)
	}
}

func i2f(i int16) float64 {
	return float64(i)/100.0;
}

func NewOperation(op int16) Operation {
	operation := Operation{}
	operation.Op = op
	operation.Detail.Op = UNDEFINED
	return operation
}

func NewLineOperation(x1,y1,x2,y2 float64) Operation {
	operation := NewOperation(LINE)
	operation.Detail = NewLineInstruction(x1,y1,x2,y2)
	return operation
}

func NewRectOperation(x1,y1,x2,y2 float64) Operation {
	operation := NewOperation(RECT)
	operation.Detail = NewRectInstruction(x1,y1,x2,y2)
	return operation
}

func NewCircleOperation(x,y,r float64) Operation {
	operation := NewOperation(CIRCLE)
	operation.Detail = NewCircleInstruction(x,y,r)
	return operation
}

func NewOvalOperation(x,y,a,b,r float64) Operation {
	operation := NewOperation(OVAL)
	operation.Detail = NewOvalInstruction(x,y,a,b,r)
	return operation
}

func NewSetOperation(name string,v float64) Operation {
	if !ValidName(name) {
		return NewOperation(UNDEFINED)
	}
	operation := NewOperation(SET)
	operation.Name = name
	operation.Detail = NewSetInstruction(name,v)
	return operation
}

func NewUseOperation(name string,ins Instruction) Operation {
	if !ValidName(name) {
		return NewOperation(UNDEFINED)
	}
	operation := NewOperation(USE)
	operation.Name = name
	operation.Detail = ins
	return operation
}

func NewPolygonOperation(coords ...float64) Operation {
	if len(coords) >= 4 && len(coords)%2 == 0 {
		operation := NewOperation(POLYGON)
		operation.Detail = NewPolygonInstruction(coords...)
		return operation
	} else {
		return NewOperation(UNDEFINED)
	}
}

func NewInstruction(op int16) Instruction {
	instruction := Instruction{}
	instruction.Op = op
	return instruction
}

func NewLineInstruction(x1,y1,x2,y2 float64) Instruction {
	instruction := NewInstruction(LINE)
	instruction.Args = []int16 {f2i(x1),f2i(y1),f2i(x2),f2i(y2)}
	return instruction;
}

func NewRectInstruction(x1,y1,x2,y2 float64) Instruction {
		instruction := NewInstruction(RECT)
		instruction.Args = []int16 {f2i(x1),f2i(y1),f2i(x2),f2i(y2)}
		return instruction;
}

func NewCircleInstruction(x,y,r float64) Instruction {
		instruction := NewInstruction(CIRCLE)
		instruction.Args = []int16 {f2i(x),f2i(y),f2i(r)}
		return instruction;
}

func NewOvalInstruction(x,y,a,b,t float64) Instruction {
		instruction := NewInstruction(OVAL)
		instruction.Args = []int16 {f2i(x),f2i(y),f2i(a),f2i(b),f2i(t)}
		return instruction;
}

func NewPolygonInstruction(coords ...float64) Instruction {
	if len(coords) >= 4 && len(coords)%2 == 0 {
		instruction := NewInstruction(POLYGON)
		instruction.Args = make([]int16,len(coords));
		for i,_ := range(coords) {
			instruction.Args[i] = f2i(coords[i])
		}
		return instruction;
	} else {
		return NewInstruction(UNDEFINED)
	}
}

func NewSetInstruction(name string,v float64) Instruction {
		instruction := NewInstruction(SET)
		instruction.Args = []int16 {f2i(v)}
		return instruction;
}

func HasTransform(op int16) bool {
	return op == PUSH || op == POP || op == TRANSFORM
}

func HasName(op int16) bool {
	return op == USE || op == SET || op == TRANSFORM
}

func ValidName(name string) bool {
	if len(name) == 0 {
		return false
	}
	for _,c := range name {
		if !unicode.IsLetter(c) {
			return false
		}
	}
	return true
}

func OperationPrint(op Operation) {
	fmt.Printf("%s",OperationNames[op.Op])
	if HasName(op.Op) && op.Name != "" {
		fmt.Printf(" %s",op.Name)
	}
	if HasTransform(op.Op) {
		fmt.Print(" ")
		TransformPrint(op.Transform)
	}
	fmt.Print(" ")
	InstructionPrint(op.Detail)
}

func TransformPrint(tf Transform) {
	fmt.Printf("[[%f,%f;%f,%f],[%f,%f]]",i2f(tf.a),i2f(tf.b),i2f(tf.c),i2f(tf.d),
		i2f(tf.x),i2f(tf.y))
}

func InstructionPrint(ins Instruction) {
	fmt.Printf("%s",OperationNames[ins.Op])
	for _,v := range ins.Args {
		fmt.Printf(" %f",i2f(v))
	}
}

func OperationToString(op Operation) string {
	ret := fmt.Sprintf("%s",OperationNames[op.Op])
	if HasName(op.Op) && op.Name != "" {
		ret += fmt.Sprintf(" %s",op.Name)
	}
	if HasTransform(op.Op) {
		ret += fmt.Sprint(" ")
		ret += TransformToString(op.Transform)
	}
	ret += fmt.Sprint(" ")
	ret += InstructionToString(op.Detail)
	return ret
}

func TransformToString(tf Transform) string {
	return fmt.Sprintf("[[%f,%f;%f,%f],[%f,%f]]",i2f(tf.a),i2f(tf.b),i2f(tf.c),
		i2f(tf.d),i2f(tf.x),i2f(tf.y))
}

func InstructionToString(ins Instruction) string {
	ret := fmt.Sprintf("%s",OperationNames[ins.Op])
	for _,v := range ins.Args {
		ret += fmt.Sprintf(" %f",i2f(v))
	}
	return ret
}

func OperationEqual(op1, op2 Operation) bool {
	if op1.Op != op2.Op {
		return false
	}
	if op1.Op == UNDEFINED || op1.Op == POP {
		return true
	}
	if HasName(op1.Op) && op1.Name != op2.Name {
		return false
	}
	if HasTransform(op1.Op) && !TransformEqual(op1.Transform,op2.Transform) {
		return false
	}
	if op1.Op == PUSH {
		return true
	}
	return InstructionEqual(op1.Detail,op2.Detail)
}

func TransformEqual(tf1,tf2 Transform) bool {
	return tf1 == tf2
}

func InstructionEqual(ins1,ins2 Instruction) bool {
	return reflect.DeepEqual(ins1,ins2)
}
