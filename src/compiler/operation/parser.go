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
		if Verbose {
			log.Output(1,"operation not recognized: "+tokens[0])
		}
		operation.Op = UNDEFINED
		return operation
	}

	nnums := len(tokens)-1
	if operation.Op == POLYGON || operation.Op == LINE || operation.Op == RECT ||
		 operation.Op == CIRCLE {
		if operation.Op == POLYGON && (nnums%2 != 0 || nnums < 6) {
			if Verbose {
				log.Output(1,"odd number of polygon coordinates")
			}
			operation.Op = UNDEFINED
			return operation
		} else if (operation.Op == LINE || operation.Op == RECT) && nnums != 4 {
			if Verbose {
				log.Output(1,"incorrect number of line coordinates")
			}
			operation.Op = UNDEFINED
			return operation
		} else if (operation.Op == CIRCLE && nnums != 3) {
			if Verbose {
				log.Output(1,"incorrect number of circle coordinates")
			}
			operation.Op = UNDEFINED
			return operation
		}
		operation.Detail.Args = make([]int16,nnums)
		for i := 1; i <= nnums; i++ {
			x,err := strconv.ParseFloat(tokens[i],64)
			if err != nil {
				if Verbose {
					log.Output(1,"invalid float "+tokens[i])
				}
				operation.Op = UNDEFINED
				return operation
			}
			operation.Detail.Args[i-1] = f2i(x)
		}
		operation.Detail.Op = operation.Op
		return operation
	}

	return operation
}
