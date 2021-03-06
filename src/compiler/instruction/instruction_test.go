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

import "testing"
import "compiler/operation"

func TestGetInstruction(t *testing.T) {
	tests := []string {
		"line 120 300 110 310",
		"rect 110 0 0 110",
		"polygon 110 100 0 10 210 220",
		"oval 110 110 100 50",
		"polygon 110 100 0 10 210 220",
		"oval 110 110 100 50",
		"rect 110 0 0 110",
		"line 120 300 110 310",
	}
	results := []Instruction {
		{operation.LINE,[]int16{120,300,110,310}},
		{operation.RECT,[]int16{110,0,110,110,0,110,0,0}},
		{operation.POLYGON,[]int16{6,110,100,0,10,210,220}},
		{operation.OVAL,[]int16{210,110,210,160,110,160,10,160,10,110,10,60,110,60,210,60}},
		{operation.POLYGON,[]int16{6,110,100,0,10,210,220}},
		{operation.OVAL,[]int16{210,110,210,160,110,160,10,160,10,110,10,60,110,60,210,60}},
		{operation.RECT,[]int16{110,0,110,110,0,110,0,0}},
		{operation.LINE,[]int16{120,300,110,310}},
	}
	for i := 0; i < len(tests); i++ {
		parser := operation.NewLineParser()
		oper,err := parser.ParseLine(tests[i])
		if err != nil {
			t.Errorf("Failed to parse: %s",tests[i])
		}
		args,ok := operation.ValuesToInt(oper.Args)
		if !ok {
			t.Errorf("Failed to evaluate arguments: %s",oper.ToString())
		}
		if oper.Command == operation.RECT {
			args = expandRect(args)
		}
		if oper.Command == operation.OVAL {
			args = expandOval(args)
		}
		inst,err := GetInstruction(oper.Command,args)
		if err != nil {
			t.Errorf("Failed to generate instruction: %s -- %s",
				oper.ToString(),err.Error())
		}
		if !inst.Equal(results[i]) {
			t.Errorf("Wrong instruction: expect %s, got %s",
				results[i].ToString(),inst.ToString())
		}
	}
}

func TestBytesToInstructions(t *testing.T) {
	tests := []Instruction {
		{operation.LINE,[]int16{120,300,110,310}},
		{operation.RECT,[]int16{110,0,110,110,0,110,0,0}},
		{operation.POLYGON,[]int16{6,110,100,0,10,210,220}},
		{operation.OVAL,[]int16{210,110,210,160,110,160,10,160,10,110,10,60,110,60,210,60}},
		{operation.POLYGON,[]int16{6,110,100,0,10,210,220}},
		{operation.OVAL,[]int16{210,110,210,160,110,160,10,160,10,110,10,60,110,60,210,60}},
		{operation.RECT,[]int16{110,0,110,110,0,110,0,0}},
		{operation.LINE,[]int16{120,300,110,310}},
	}
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
