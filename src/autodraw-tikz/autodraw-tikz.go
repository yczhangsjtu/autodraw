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
		switch inst.Command {
		case operation.LINE:
		case operation.RECT:
		case operation.POLYGON:
		case operation.CIRCLE:
		case operation.OVAL:
		default:
			return "",NewTikzError("invalid instruction: "+inst.ToString())
		}
	}

	return fmt.Sprintf("\\begin{tikzpicture}%s\n%s\\end{tikzpicture}\n",
		options,code),nil
}
