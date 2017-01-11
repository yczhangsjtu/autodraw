package operation

import "fmt"

const (
	UNDEFINED int16 = iota
	LINE
	RECT
	CIRCLE
	OVAL
	POLYGON
)

var OperationNames = []string { "undefined","line","rect","circle","oval","polygon" }

var OperationNameMap = map[string] int16 { "undefined":UNDEFINED, "line":LINE, "rect":RECT, "circle":CIRCLE, "oval":OVAL, "polygon":POLYGON }

type Instruction struct {
	Op int16
	Args []int16
}

type Operation struct {
	Op int16
	Detail Instruction
}

func OperationPrint(op Operation) {
	fmt.Printf("%s",OperationNames[op.Op])
}
