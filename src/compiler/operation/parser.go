package operation

import "strings"
import "log"
import "strconv"

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

	nnums := len(tokens)-1
	if operation.Op == POLYGON || operation.Op == LINE || operation.Op == RECT ||
		 operation.Op == CIRCLE {
		if operation.Op == POLYGON && (nnums%2 != 0 || nnums < 6) {
			log.Output(1,"odd number of polygon coordinates")
			operation.Op = UNDEFINED
			return operation
		} else if (operation.Op == LINE || operation.Op == RECT) && nnums != 4 {
			log.Output(1,"incorrect number of line coordinates")
			operation.Op = UNDEFINED
			return operation
		} else if (operation.Op == CIRCLE && nnums != 3) {
			log.Output(1,"incorrect number of circle coordinates")
			operation.Op = UNDEFINED
			return operation
		}
		operation.Detail.Args = make([]int16,nnums)
		for i := 1; i <= nnums; i++ {
			x,err := strconv.ParseFloat(tokens[i],64)
			if err != nil {
				log.Output(1,"invalid float "+tokens[i])
				operation.Op = UNDEFINED
				return operation
			}
			operation.Detail.Args[i-1] = int16(x*100)
		}
		operation.Detail.Op = operation.Op
		return operation
	}

	return operation
}
