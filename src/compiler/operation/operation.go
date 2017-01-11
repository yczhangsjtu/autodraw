package operation

import "fmt"
import "reflect"

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
	instruction.Args = []int16 {int16(x1*100),int16(y1*100),
															int16(x2*100),int16(y2*100)}
	return instruction;
}

func NewRectInstruction(x1,y1,x2,y2 float64) Instruction {
		instruction := NewInstruction(RECT)
		instruction.Args = []int16 {int16(x1*100),int16(y1*100),
																int16(x2*100),int16(y2*100)}
		return instruction;
}

func NewCircleInstruction(x,y,r float64) Instruction {
		instruction := NewInstruction(CIRCLE)
		instruction.Args = []int16 {int16(x*100),int16(y*100),int16(r*100)}
		return instruction;
}

func NewPolygonInstruction(coords ...float64) Instruction {
	if len(coords) >= 4 && len(coords)%2 == 0 {
		instruction := NewInstruction(POLYGON)
		instruction.Args = make([]int16,len(coords));
		for i,_ := range(coords) {
			instruction.Args[i] = int16(coords[i]*100)
		}
		return instruction;
	} else {
		return NewInstruction(UNDEFINED)
	}
}

func HasTransform(op int16) bool {
	return op == USE || op == PUSH || op == POP || op == TRANSFORM
}

func HasName(op int16) bool {
	return op == SET || op == TRANSFORM
}

func OperationPrint(op Operation) {
	fmt.Printf("%s",OperationNames[op.Op])
	if HasName(op.Op) {
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
	fmt.Printf("[[%f,%f;%f,%f],[%f,%f]]",float64(tf.a)/100,float64(tf.b)/100,
			float64(tf.c)/100,float64(tf.d)/100,float64(tf.x)/100,float64(tf.y)/100)
}

func InstructionPrint(ins Instruction) {
	fmt.Printf("%s",OperationNames[ins.Op])
	for _,v := range ins.Args {
		fmt.Printf(" %f",float64(v)/100)
	}
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
