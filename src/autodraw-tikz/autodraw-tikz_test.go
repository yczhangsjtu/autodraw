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

import "testing"
import "compiler/operation"
import "compiler/instruction"

func TestUpdate(t *testing.T) {
	tests := []instruction.Instruction {
		{operation.LINE,[]int16{120,300,110,310}},
		{operation.RECT,[]int16{110,0,0,110}},
		{operation.POLYGON,[]int16{6,110,100,0,10,210,220}},
		{operation.CIRCLE,[]int16{110,110,100}},
		{operation.OVAL,[]int16{110,110,100,50,50}},
	}
	tz := NewTikz()
	for _,inst := range tests {
		err := tz.Update(inst)
		if err != nil {
			t.Errorf("Failed to process instruction %s: %s",
				inst.ToString(),err.Error())
		}
	}
}
