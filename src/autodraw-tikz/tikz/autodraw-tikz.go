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
package tikz

import "fmt"
import "compiler/operation"
import "compiler/instruction"

type Tikz struct {
	scale float64
	offsetx int16
	offsety int16

	instlist []instruction.Instruction
}

type TikzError struct {
	reason string
}

func NewTikzError(reason string) *TikzError {
	return &TikzError{reason}
}

func (e *TikzError) Error() string {
	return e.reason
}

func NewTikz() *Tikz{
	tz := new(Tikz)
	tz.scale = 1.0
	return tz
}

func (tz *Tikz) Update(inst instruction.Instruction) error {
	tz.instlist = append(tz.instlist,inst)
	return nil
}

func (tz *Tikz) GenerateTikzCode() (string,error) {
	code := ""
	options := ""

	for _,inst := range tz.instlist {
		tikzCode,err := InstToTikz(inst,tz.scale)
		if err != nil {
			return "",err
		}
		code += fmt.Sprintf("  %s\n",tikzCode)
	}

	return fmt.Sprintf("\\begin{tikzpicture}%s\n%s\\end{tikzpicture}\n",
		options,code),nil
}

func InstToTikz(inst instruction.Instruction, scale float64) (string,error) {
	switch inst.Command {
	case operation.LINE:
		return fmt.Sprintf("\\draw %s;",GenerateFloatPairs("--",
					IntsToScaledFloats(inst.Args,scale))),nil
	case operation.RECT:
		return fmt.Sprintf("\\draw %s -- cycle;",GenerateFloatPairs("--",
				IntsToScaledFloats(inst.Args,scale))),nil
	case operation.POLYGON:
		return fmt.Sprintf("\\draw %s -- cycle;",GenerateFloatPairs("--",
				IntsToScaledFloats(inst.Args[1:],scale))),nil
	case operation.CIRCLE:
		return fmt.Sprintf("\\draw %s;",GenerateFloatPairs("circle",
					IntsToScaledFloats(inst.Args,scale))),nil
	case operation.OVAL:
		fallthrough
	default:
		return "",NewTikzError("invalid instruction: "+inst.ToString())
	}
}

func IntsToFloats(args []int16) []float64 {
	ret := make([]float64,len(args))
	for i,v := range args {
		ret[i] = float64(v)
	}
	return ret
}

func IntsToScaledFloats(args []int16, scale float64) []float64 {
	ret := make([]float64,len(args))
	for i,v := range args {
		ret[i] = float64(v)/100.0 * scale
	}
	return ret
}

func GenerateFloatPairs(connect string, args []float64) string {
	if len(args) == 0 {
		return ""
	}
	num := len(args)/2
	ret := fmt.Sprintf("(%g,%g)",args[0],args[1])
	for i := 1; i < num; i++ {
		ret += fmt.Sprintf(fmt.Sprintf(" %s %s",connect,"(%g,%g)"),
			args[2*i],args[2*i+1])
	}
	if len(args)%2 == 1 {
		ret += fmt.Sprintf(fmt.Sprintf(" %s %s",connect,"(%g)"),args[len(args)-1])
	}
	return ret
}
