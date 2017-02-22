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
package instruction

import "fmt"
import "strconv"
import "reflect"
import "compiler/operation"

type Instruction struct {
	Command int16
	Args []int16
}

type InstructionError struct {
	reason string
}

func NewInstruction() Instruction {
	return Instruction{0,[]int16{}}
}

func NewInstructionError(reason string) *InstructionError {
	e := InstructionError{reason}
	return &e
}

func (e *InstructionError) Error() string {
	return e.reason
}

func (inst *Instruction) Equal(inst2 Instruction) bool {
	return reflect.DeepEqual(*inst,inst2)
}

func (inst *Instruction) ToString() string {
	return fmt.Sprintf("%s %d",operation.GetName(inst.Command),inst.Args)
}

func (inst *Instruction) ToBytes() []byte {
	ret := make([]byte,len(inst.Args)*2+2)
	ret[0] = byte(inst.Command/256)
	ret[1] = byte(inst.Command%256)
	for i := 0; i < len(inst.Args); i++ {
		ret[2*i+2] = byte(uint(inst.Args[i])/256)
		ret[2*i+3] = byte(uint(inst.Args[i])%256)
	}
	return ret
}

func InstructionsToBytes(insts []Instruction) []byte {
	ret := []byte{}
	for _, inst := range insts {
		ret = append(ret, inst.ToBytes()...)
	}
	return ret
}

func composeBytes(b1,b2 byte) int16 {
	return int16(uint(b1)*256+uint(b2))
}

func getInt16(data []byte, i *int) (int16,error) {
	if (*i)*2+1 >= len(data) {
		return 0,NewInstructionError("index out of range")
	}
	ret := composeBytes(data[(*i)*2],data[(*i)*2+1])
	(*i)++
	return ret,nil
}

func getInts16(data []byte, ptr *int, count int) ([]int16,error) {
	ret := make([]int16,count)
	var err error
	for i := 0; i < count; i++ {
		ret[i],err = getInt16(data,ptr)
		if err != nil {
			return ret,err
		}
	}
	return ret,err
}

func BytesToInstructions(data []byte) ([]Instruction,error) {
	ptr := 0
	ret := []Instruction{}
	for {
		if ptr*2+1 >= len(data) {
			return ret,nil
		}

		command,err := getInt16(data,&ptr)
		commandType := operation.GetType(command)
		if commandType != operation.DRAW_FIXED &&
			commandType != operation.DRAW_UNDETERMINED {
			return ret,NewInstructionError(
				"Invalid command number "+strconv.Itoa(int(command)))
		}

		var argNum int16
		if commandType == operation.DRAW_FIXED {
			argNum = int16(operation.FinalArgNum(command))
		} else {
			argNum,err = getInt16(data,&ptr)
			if err != nil {
				return ret,err
			}
		}

		args,err := getInts16(data,&ptr,int(argNum))
		if err != nil {
			return ret,err
		}
		inst,err := GetInstruction(command,args)
		if err != nil {
			return ret,err
		}
		ret = append(ret,inst)
	}
	return ret,nil
}

func GetInstruction(command int16, args []int16) (Instruction,error) {
	inst := NewInstruction()
	inst.Command = command
	switch command {
	case operation.RECT:
		fallthrough
	case operation.LINE:
		fallthrough
	case operation.OVAL:
		if len(args) != int(operation.FinalArgNum(command)) {
			return NewInstruction(),NewInstructionError(
				"invalid number of arguments: "+operation.GetName(command)+
			  " requires "+strconv.Itoa(int(operation.FinalArgNum(command)))+
				" args, got "+strconv.Itoa(len(args)))
		}
		inst.Args = args
		return inst,nil
	case operation.POLYGON:
		if len(args) < 4 || len(args)%2 == 1 {
			return NewInstruction(),NewInstructionError(
				"invalid number of arguments: got "+
				strconv.Itoa(len(args))+" for polygon")
		}
		inst.Args = addLengthPrefix(args)
		return inst,nil
	default:
		return NewInstruction(),NewInstructionError(
			"invalid draw command: "+operation.GetName(command))
	}
}

func addLengthPrefix(args []int16) []int16 {
	return append([]int16{int16(len(args))},args...)
}

func expandRect(args []int16) []int16 {
	ret := make([]int16,8)
	ret[0],ret[1] = args[0],args[1]
	ret[2],ret[3] = args[0],args[3]
	ret[4],ret[5] = args[2],args[3]
	ret[6],ret[7] = args[2],args[1]
	return ret
}

func expandOval(args []int16) []int16 {
	ret := make([]int16,16)
	x,y,a,b := args[0],args[1],args[2],args[3]
	ret[0],ret[1] = x+a,y
	ret[2],ret[3] = x+a,y+b
	ret[4],ret[5] = x,y+b
	ret[6],ret[7] = x-a,y+b
	ret[8],ret[9] = x-a,y
	ret[10],ret[11] = x-a,y-b
	ret[12],ret[13] = x,y-b
	ret[14],ret[15] = x+a,y-b
	return ret
}
