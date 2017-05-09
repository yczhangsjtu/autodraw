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

import (
	"testing"
)

func BenchmarkParseLine(b *testing.B) {
	tests := []string{
		"undefined",
		"line 120 300 110 310",
		"line 120 300 110 Ab_",
		"rect 110 0 0 110",
		"polygon 110 100 0 10 210 220",
		"oval 110 110 100 50",
		"text 0 0 256 \"Hello world!\"",
		"set s 110",
		"use T",
		"push T",
		"pop",
		"transform T -100 10 -10 100 200 200",
		"transform T A B C D E F",
		"transform T Alice 0 Bob 100 Carror F",
		"rotate s 110",
		"scale s 110 -100",
		"translate s 110 -100",
		"draw plane",
		"import plane",
		"import no-plane",
		"line 120 300 110",
		"rect 110 0 0",
		"line 120 300 110 310",
		"rect 110 0 0 110",
		"polygon 110 100 0 10 210 220",
		"oval 110 110 100 50",
		"set Alice 15",
		"set Bob 20",
		"set Bob Alice",
		"set Alice 110",
		"set Carror -10",
		"set Bob Carror",
		"set Carror Alice",
		"transform T Alice 0 Bob 0 110 Carror",
		"rotate X Alice",
		"scale Y Bob Carror",
		"translate W Bob Carror",
		"set Q T",
		"set P Q",
		"begin subfigure",
		"set X 0",
		"set Y 1",
		"set M 2",
		"line 120 300 110 310",
		"rect 110 0 0 110",
		"polygon 110 100 0 10 210 220",
		"oval 110 110 100 50",
		"end",
		"set Z X",
		"set X Y",
		"set U W",
		"set V U",
		"push T",
		"line 120 300 110 310",
		"rect 110 0 0 110",
		"polygon 110 100 0 10 210 220",
		"oval 110 110 100 50",
		"pop",
		"use T",
		"line 120 300 110 310",
		"use Q",
		"rect 110 0 0 110",
		"use P",
		"draw subfigure",
		"polygon 110 100 0 10 210 220",
		"oval 110 110 100 50",
	}
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			ParseLine(test)
		}
	}
}

func TestParseLine(t *testing.T) {
	tests := []string{
		"undefined",
		"line 120 300 110 310",
		"rect 110 0 0 110",
		"polygon 110 100 0 10 210 220",
		"oval 110 110 100 50",
		"text 0 0 256 \"Hello world!\"",
		"set s 110",
		"set t \"Hello world!\"",
		"use T",
		"push T",
		"pop",
		"transform T -100 10 -10 100 200 200",
		"rotate T 110",
		"scale T 110 -100",
		"translate T 110 -100",
		"transform T Alice 0 Bob 0 110 Carror",
		"draw plane",
		"import plane",
		"line 120 300 110",
		"rect 110 0 0",
	}
	expects := []Operation{
		NewOperation(UNDEFINED),
		newLineOperation(NewNumberValues(120, 300, 110, 310)...),
		newRectOperation(NewNumberValues(110, 0, 0, 110)...),
		newPolygonOperation(NewNumberValues(110, 100, 0, 10, 210, 220)...),
		newOvalOperation(NewNumberValues(110, 110, 100, 50)...),
		newTextOperation(append(NewNumberValues(0,0,256),NewStringValue("Hello world!"))...),
		newSetOperation("s", NewNumberValue(110)),
		newSetOperation("t", NewStringValue("Hello world!")),
		newUseOperation("T"),
		newPushOperation("T"),
		newPopOperation(),
		newTransformOperation("T", NewNumberValues(-100, 10, -10, 100, 200, 200)...),
		newRotateOperation("T", NewNumberValue(110)),
		newScaleOperation("T", NewNumberValues(110,-100)...),
		newTranslateOperation("T", NewNumberValues(110,-100)...),
		newTransformOperation("T", NewValues("Alice","0","Bob","0","110","Carror")...),
		newDrawOperation("plane"),
		newImportOperation("plane"),
		NewOperation(UNDEFINED),
		NewOperation(UNDEFINED),
		NewOperation(UNDEFINED),
	}

	for i, test := range tests {
		result, err := ParseLine(test)
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
