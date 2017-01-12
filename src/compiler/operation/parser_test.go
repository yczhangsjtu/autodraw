package operation

import "fmt"
import "testing"

func TestParse(t *testing.T) {
	tests := []string {
		"undefined",
		"line 1.2 3.0 1.1 3.1",
		"rect 1.1 0.0 0.0 1.1",
		"polygon 1.1 1.0 0.0 0.1 2.1 2.2",
		"circle 1.1 1.1 1.0",
	}
	expects := []Operation{
		NewOperation(UNDEFINED),
		NewLineOperation(1.2,3.0,1.1,3.1),
		NewRectOperation(1.1,0.0,0.0,1.1),
		NewPolygonOperation(1.1,1.0,0.0,0.1,2.1,2.2),
		NewCircleOperation(1.1,1.1,1.0),
	}

	for i,test := range tests {
		result := Parse(test)
		if !OperationEqual(expects[i],result) {
			t.Errorf("Parser failed for [%s]\n",test)
			OperationPrint(expects[i])
			fmt.Print(" vs. ")
			OperationPrint(result)
			fmt.Println("")
		} else {
			t.Logf("Parser correct for [%s]:\n",test)
		}
	}
}
