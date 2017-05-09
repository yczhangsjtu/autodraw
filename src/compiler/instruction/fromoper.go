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
	"compiler/operation"
)

func OperationToInstruction(oper operation.Operation) (*Instruction,error) {
	var args []int16
	switch oper.Command {
		case operation.LINE:
		  args,ok := operation.ValuesToInt(oper.Args)
			if !ok {
				return nil,NewInstructionError("Operation argument not integer")
			}
			return GetInstance(LINE_STRIP,args)
		case operation.RECT:
			args = make([]int16,10)
			args[0],args[1] = oper.Args[0].Number,oper.Args[1].Number
			args[2],args[3] = oper.Args[0].Number,oper.Args[3].Number
			args[4],args[5] = oper.Args[2].Number,oper.Args[3].Number
			args[6],args[7] = oper.Args[2].Number,oper.Args[1].Number
			args[8],args[9] = oper.Args[0].Number,oper.Args[1].Number
			return GetInstance(LINE_STRIP,args)
		case operation.OVAL:
			args = make([]int16,26)
			x,y,a,b := oper.Args[0].Number,oper.Args[1].Number,oper.Args[2].Number,oper.Args[3].Number
			x0,y0,x1,y1,x2,y2,x3,y3 := x+a,y+b,x-a,y+b,x-a,y-b,x+a,y-b
			args[0], args[1]  = x+a,y
			args[6], args[7]  = x,y+b
			args[12],args[13] = x-a,y
			args[18],args[19] = x,y-b
			args[24],args[25] = x+a,y
			args[2], args[3]  = int16(float64(x0)*0.55+float64(args[0])*0.45), int16(float64(y0)*0.55+float64(args[1])*0.45)
			args[4], args[5]  = int16(float64(x0)*0.55+float64(args[6])*0.45), int16(float64(y0)*0.55+float64(args[7])*0.45)
			args[8], args[9]  = int16(float64(x1)*0.55+float64(args[6])*0.45), int16(float64(y1)*0.55+float64(args[7])*0.45)
			args[10],args[11] = int16(float64(x1)*0.55+float64(args[12])*0.45),int16(float64(y1)*0.55+float64(args[13])*0.45)
			args[14],args[15] = int16(float64(x2)*0.55+float64(args[12])*0.45),int16(float64(y2)*0.55+float64(args[13])*0.45)
			args[16],args[17] = int16(float64(x2)*0.55+float64(args[18])*0.45),int16(float64(y2)*0.55+float64(args[19])*0.45)
			args[20],args[21] = int16(float64(x3)*0.55+float64(args[18])*0.45),int16(float64(y3)*0.55+float64(args[19])*0.45)
			args[22],args[23] = int16(float64(x3)*0.55+float64(args[24])*0.45),int16(float64(y3)*0.55+float64(args[25])*0.45)
			return GetInstance(CURVE,args)
		case operation.POLYGON:
			args,ok := operation.ValuesToInt(oper.Args)
			if !ok {
				return nil,NewInstructionError("Operation argument not integer")
			}
			args = append(args,args[0],args[1])
			return GetInstance(LINE_STRIP,args)
		case operation.TEXT:
			args = make([]int16,3)
			args[0],args[1],args[2] = oper.Args[0].Number,oper.Args[1].Number,oper.Args[2].Number
			for _,c := range oper.Args[3].Text {
				args = append(args,int16(c))
			}
			return GetInstance(TEXT,args)
	}
	return nil,NewInstructionError("Invalid command "+operation.GetName(oper.Command))
}
