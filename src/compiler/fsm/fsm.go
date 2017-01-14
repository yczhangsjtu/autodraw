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
package fsm

import "log"
import "compiler/operation"
import "compiler/transformer"

type Operation operation.Operation
type Transformer transformer.Transform
type VarTable map[string] float64

type FSM struct {
	tfstack transformer.TFStack
	vartable VarTable
}

func f2i(f float64) int16 {
	if f > 0.0 {
		return int16(f*100.0+0.5)
	} else {
		return int16(f*100.0-0.5)
	}
}

func i2f(i int16) float64 {
	return float64(i)/100.0;
}

func NewVarTable() VarTable {
	return make(VarTable)
}

func NewFSM() FSM {
	fsm := FSM{transformer.NewTFStack(),NewVarTable()}
	return fsm
}

func Update(fsm *FSM, oper Operation) {
	switch oper.Op {
	case operation.UNDEFINED:
		if operation.Verbose {
			log.Output(1,"Undefined operation")
		}
	case operation.LINE:
		fallthrough
	case operation.RECT:
		fallthrough
	case operation.CIRCLE:
		fallthrough
	case operation.OVAL:
		fallthrough
	case operation.POLYGON:
		return
	case operation.SET:
		fsm.vartable[oper.Name] = i2f(oper.Detail.Args[0])
	}
}
