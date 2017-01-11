package operation

import "strings"
import "log"

func Parse(line string) Operation {
	var operation Operation
	var ok bool

	line = strings.Trim(line," ")
	tokens := strings.Split(line," ")

	operation.Op, ok = OperationNameMap[tokens[0]]
	if !ok {
		log.Output(1,"operation not recognized: "+tokens[0])
		operation.Op = UNDEFINED
		return operation
	}

	return operation
}
