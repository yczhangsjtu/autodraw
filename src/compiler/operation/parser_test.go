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
package operation

import "fmt"
import "testing"

func TestParse(t *testing.T) {
	Verbose = false
	tests := []string {
		"undefined",
		"line 1.2 3.0 1.1 3.1",
		"rect 1.1 0.0 0.0 1.1",
		"polygon 1.1 1.0 0.0 0.1 2.1 2.2",
		"circle 1.1 1.1 1.0",
		"line 1.2 3.0 1.1",
		"rect 1.1 0.0 0.0",
		"polygon 1.1 1.0 0.0 0.1 2.2",
		"circle 1.1 1.1",
	}
	expects := []Operation{
		NewOperation(UNDEFINED),
		NewLineOperation(1.2,3.0,1.1,3.1),
		NewRectOperation(1.1,0.0,0.0,1.1),
		NewPolygonOperation(1.1,1.0,0.0,0.1,2.1,2.2),
		NewCircleOperation(1.1,1.1,1.0),
		NewOperation(UNDEFINED),
		NewOperation(UNDEFINED),
		NewOperation(UNDEFINED),
		NewOperation(UNDEFINED),
	}

	for i,test := range tests {
		result := Parse(test)
		if !OperationEqual(expects[i],result) {
			t.Errorf("Parser failed for [%s]\n",test)
			OperationPrint(expects[i])
			fmt.Print(" vs. ")
			OperationPrint(result)
			fmt.Println("")
		}
	}
}
