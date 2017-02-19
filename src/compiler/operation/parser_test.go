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

import "testing"

func BenchmarkParseLine(b *testing.B) {
	tests := []string{
		"undefined",
		"line 120 300 110 310",
		"rect 110 0 0 110",
		"polygon 110 100 0 10 210 220",
		"circle 110 110 100",
		"oval 110 110 100 50 50",
		"set s 110",
		"use T",
		"push T",
		"pop",
		"transform T -100 10 -10 100 200 200",
		"rotate s 110",
		"scale s 110 -100",
		"draw plane",
		"import plane",
		"import no-plane",
		"line 120 300 110",
		"rect 110 0 0",
		"circle 110 110",
	}
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			parser := NewLineParser()
			parser.ParseLine(test)
		}
	}
}

func TestParseLine(t *testing.T) {
	Verbose = false
	tests := []string{
		"undefined",
		"line 120 300 110 310",
		"rect 110 0 0 110",
		"polygon 110 100 0 10 210 220",
		"circle 110 110 100",
		"oval 110 110 100 50 50",
		"set s 110",
		"use T",
		"push T",
		"pop",
		"transform T -100 10 -10 100 200 200",
		"rotate T 110",
		"scale T 110 -100",
		"draw plane",
		"import plane",
		"line 120 300 110",
		"rect 110 0 0",
		"circle 110 110",
	}
	expects := []Operation{
		NewOperation(UNDEFINED),
		newLineOperation(NewNumberValues(120, 300, 110, 310)...),
		newRectOperation(NewNumberValues(110, 0, 0, 110)...),
		newPolygonOperation(NewNumberValues(110, 100, 0, 10, 210, 220)...),
		newCircleOperation(NewNumberValues(110, 110, 100)...),
		newOvalOperation(NewNumberValues(110, 110, 100, 50, 50)...),
		newSetOperation("s", NewNumberValue(110)),
		newUseOperation("T"),
		newPushOperation("T"),
		newPopOperation(),
		newTransformOperation("T", NewNumberValues(-100, 10, -10, 100, 200, 200)...),
		newRotateOperation("T", NewNumberValue(110)),
		newScaleOperation("T", NewNumberValues(110,-100)...),
		newDrawOperation("plane"),
		newImportOperation("plane"),
		NewOperation(UNDEFINED),
		NewOperation(UNDEFINED),
		NewOperation(UNDEFINED),
	}

	for i, test := range tests {
		parser := NewLineParser()
		result, err := parser.ParseLine(test)
		if !expects[i].Equal(result) || result.Command != UNDEFINED && err != nil {
			errorString := ""
			if err != nil {
				errorString = err.Error()
			}
			t.Errorf("Parser failed for [%s], expect (%s), got (%s): %s\n",
				test, expects[i].ToString(), result.ToString(), errorString)
		}
	}
}
