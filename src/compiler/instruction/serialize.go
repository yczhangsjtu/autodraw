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

import (
	"fmt"
)

func InstructionsToBytes(insts []*Instruction) []byte {
	ret := []byte{}
	for _, inst := range insts {
		ret = append(ret, inst.ToBytes()...)
	}
	return ret
}

func BytesToInstructions(data []byte) ([]*Instruction,error) {
	ptr := 0
	ret := []*Instruction{}
	for {
		if ptr*2+1 >= len(data) {
			return ret,nil
		}

		command,err := getInt16(data,&ptr)
		if err != nil {
			return ret,NewInstructionError("error in getting command: "+err.Error())
		}
		argNum,err := getInt16(data,&ptr)
		if err != nil {
			return ret,NewInstructionError("error in getting arg number: "+err.Error())
		}
		if !ExpectArgNum(command,int(argNum)) {
			return ret,NewInstructionError(fmt.Sprintf("wrong number of arguments for %s: %d",GetName(command),argNum))
		}
		if(ptr + int(argNum) >= len(data)) {
			return ret,NewInstructionError(fmt.Sprintf("insufficient data at %d",ptr))
		}
		args,err := getInts16(data,&ptr,int(argNum))
		if err != nil {
			return ret,err
		}
		inst,err := GetInstance(command,args)
		if err != nil {
			return ret,err
		}
		ret = append(ret,inst)
	}
	return ret,nil
}
