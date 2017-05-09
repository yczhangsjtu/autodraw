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
	"testing"
	"compiler/operation"
)

func TestOperationToInstruction(t *testing.T) {
	tests := []string {
		"line 120 300 110 310",
		"rect 110 0 0 110",
		"polygon 110 100 0 10 210 220",
		"oval 110 110 100 50",
		"text 0 0 256 \"HELLO\"",
		"node 0 0 256 hello \"HELLO\"",
	}
	results := []Instruction {
		{LINE_STRIP,[]int16{120,300,110,310}},
		{LINE_STRIP,[]int16{110,0,110,110,0,110,0,0,110,0}},
		{LINE_STRIP,[]int16{110,100,0,10,210,220,110,100}},
		{CURVE,[]int16{210,110,210,137,165,160,110,160,55,160,10,137,10,110,10,82,55,60,110,60,165,60,210,82,210,110}},
		{TEXT,[]int16{0,0,256,72,69,76,76,79}},
		{NODE,[]int16{0,0,256,72,69,76,76,79}},
	}
	for i := 0; i < len(tests); i++ {
		oper,err := operation.ParseLine(tests[i])
		if err != nil {
			t.Errorf("Failed to parse: %s",tests[i])
		}
		inst,err := OperationToInstruction(oper)
		if err != nil {
			t.Errorf("Failed to generate instruction: %s -- %s",
				oper.ToString(),err.Error())
		}
		if !inst.Equal(&results[i]) {
			t.Errorf("Wrong instruction: expect %s, got %s",
				results[i].ToString(),inst.ToString())
		}
	}
}
