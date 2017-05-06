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
)

func TestSerialization(t *testing.T) {
	tests := make([]*Instruction,3)
	tests[0],_ = GetInstance(LINE_STRIP,[]int16{120,300,110,310})
	tests[1],_ = GetInstance(CURVE,[]int16{110,0,110,110,0,110,0,0})
	tests[2],_ = GetInstance(TEXT,[]int16{0,0,20,65,66,67,68})
	bytes := InstructionsToBytes(tests)
	results,err := BytesToInstructions(bytes)
	if err != nil {
		t.Errorf("Error in BytesToInstructions: %s",err.Error())
	}
	if len(tests) != len(results) {
		t.Errorf("Got wrong number of results: %d vs %d",len(results),len(tests))
	}
	for i := 0; i < len(tests); i++ {
		if !tests[i].Equal(results[i]) {
			t.Errorf("Result wrong at %d instruction: expect %s, got %s",i+1,
				tests[i].ToString(),results[i].ToString())
		}
	}
}
