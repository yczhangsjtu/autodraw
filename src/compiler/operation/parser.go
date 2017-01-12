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
		return NewOperation(UNDEFINED)
	}

	nargs := len(tokens)-1
	if operation.Op == POLYGON || operation.Op == LINE || operation.Op == RECT ||
		 operation.Op == CIRCLE || operation.Op == OVAL {
		if operation.Op == POLYGON && (nargs%2 != 0 || nargs < 6) {
			if Verbose {
				log.Output(1,"odd number of polygon coordinates")
			}
			return NewOperation(UNDEFINED)
		} else if (operation.Op == LINE || operation.Op == RECT) && nargs != 4 {
			if Verbose {
				log.Output(1,"incorrect number of line coordinates")
			}
			return NewOperation(UNDEFINED)
		} else if operation.Op == CIRCLE && nargs != 3 {
			if Verbose {
				log.Output(1,"incorrect number of circle coordinates")
			}
			return NewOperation(UNDEFINED)
		} else if operation.Op == OVAL && nargs != 5 {
			if Verbose {
				log.Output(1,"incorrect number of oval coordinates")
			}
			return NewOperation(UNDEFINED)
		}
		operation.Detail.Args = make([]int16,nargs)
		for i := 1; i <= nargs; i++ {
			x,err := strconv.ParseFloat(tokens[i],64)
			if err != nil {
				if Verbose {
					log.Output(1,"invalid float "+tokens[i])
				}
				return NewOperation(UNDEFINED)
			}
			operation.Detail.Args[i-1] = f2i(x)
		}
		operation.Detail.Op = operation.Op
		return operation
	}

	if operation.Op == SET {
		if nargs != 2 {
			if Verbose {
				log.Output(1,"incorrect number of set arguments")
			}
			return NewOperation(UNDEFINED)
		}
		if !ValidName(tokens[1]) {
			if Verbose {
				log.Output(1,"invalid name")
			}
			return NewOperation(UNDEFINED)
		}
		x,err := strconv.ParseFloat(tokens[2],64)
		if err != nil {
			if Verbose {
				log.Output(1,"invalid float "+tokens[2])
			}
			return NewOperation(UNDEFINED)
		}
		operation.Name = tokens[1]
		operation.Detail.Op = SET
		operation.Detail.Args = []int16{f2i(x)}
		return operation
	}

	if operation.Op == USE {
		if nargs < 3 {
			if Verbose {
				log.Output(1,"incorrect number of use arguments")
			}
			return NewOperation(UNDEFINED)
		}
		if !ValidName(tokens[1]) {
			if Verbose {
				log.Output(1,"invalid name")
			}
			return NewOperation(UNDEFINED)
		}
		tmpOperation := Parse(strings.Join(tokens[2:]," "))
		if tmpOperation.Op == UNDEFINED {
			if Verbose {
				log.Output(1,"invalid suboperation in use")
			}
			return NewOperation(UNDEFINED)
		}
		operation.Name = tokens[1]
		operation.Detail = tmpOperation.Detail
		return operation
	}

	if operation.Op == PUSH {
		if nargs != 1 {
			if Verbose {
				log.Output(1,"incorrect number of push arguments")
			}
			return NewOperation(UNDEFINED)
		}
		if !ValidName(tokens[1]) {
			if Verbose {
				log.Output(1,"invalid name")
			}
			return NewOperation(UNDEFINED)
		}
		operation.Name = tokens[1]
		operation.Detail = NewInstruction(PUSH)
		return operation
	}

	if operation.Op == POP {
		if nargs != 0 {
			if Verbose {
				log.Output(1,"incorrect number of pop arguments")
			}
			return NewOperation(UNDEFINED)
		}
		operation.Detail = NewInstruction(POP)
		return operation
	}

	if operation.Op == TRANSFORM {
		if nargs != 7 {
			if Verbose {
				log.Output(1,"incorrect number of transform arguments")
			}
			return NewOperation(UNDEFINED)
		}
		if !ValidName(tokens[1]) {
			if Verbose {
				log.Output(1,"invalid name")
			}
			return NewOperation(UNDEFINED)
		}
		args := [6]int16{}
		for i := 2; i <= nargs; i++ {
			x,err := strconv.ParseFloat(tokens[i],64)
			if err != nil {
				if Verbose {
					log.Output(1,"invalid float "+tokens[i])
				}
				return NewOperation(UNDEFINED)
			}
			args[i-2] = f2i(x)
		}
		operation.Transform.a = args[0]
		operation.Transform.b = args[1]
		operation.Transform.c = args[2]
		operation.Transform.d = args[3]
		operation.Transform.x = args[4]
		operation.Transform.y = args[5]
		operation.Name = tokens[1]
		operation.Detail.Op = TRANSFORM
		return operation
	}

	if operation.Op == DRAW {
		if nargs != 1 {
			if Verbose {
				log.Output(1,"incorrect number of draw arguments")
			}
			return NewOperation(UNDEFINED)
		}
		if !ValidName(tokens[1]) {
			if Verbose {
				log.Output(1,"invalid name")
			}
			return NewOperation(UNDEFINED)
		}
		operation.Name = tokens[1]
		operation.Detail = NewInstruction(DRAW)
		return operation
	}

	if operation.Op == IMPORT {
		if nargs != 1 {
			if Verbose {
				log.Output(1,"incorrect number of import arguments")
			}
			return NewOperation(UNDEFINED)
		}
		if !ValidPath(tokens[1]) {
			if Verbose {
				log.Output(1,"invalid path")
			}
			return NewOperation(UNDEFINED)
		}
		operation.Name = tokens[1]
		operation.Detail = NewInstruction(IMPORT)
		return operation
	}
	return operation
}
