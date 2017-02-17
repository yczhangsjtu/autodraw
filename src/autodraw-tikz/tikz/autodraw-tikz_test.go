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
		//{operation.OVAL,[]int16{110,110,100,50,50}},
	}
	expect := "\\begin{tikzpicture}\n"+
		"  \\draw (1.2,3) -- (1.1,3.1);\n"+
		"  \\draw (1.1,0) rectangle (0,1.1);\n"+
		"  \\draw (1.1,1) -- (0,0.1) -- (2.1,2.2) -- cycle;\n"+
		"  \\draw (1.1,1.1) circle (1);\n"+
		"\\end{tikzpicture}\n"
	tz := NewTikz()
	for _,inst := range tests {
		err := tz.Update(inst)
		if err != nil {
			t.Errorf("Failed to process instruction %s: %s",
				inst.ToString(),err.Error())
		}
	}
	code,err := tz.GenerateTikzCode()
	if err != nil {
		t.Errorf("Failed to generate tikz code: %s",err.Error())
	}
	if code != expect {
		t.Errorf("Wrong tikz code generated, expected \n%s\n, got \n%s\n",
			expect,code)
	}
}

func TestInstToTikz(t *testing.T) {
	tests := []instruction.Instruction {
		{operation.LINE,[]int16{120,300,110,310}},
		{operation.RECT,[]int16{110,0,0,110}},
		{operation.POLYGON,[]int16{6,110,100,0,10,210,220}},
		{operation.CIRCLE,[]int16{110,110,100}},
	}
	expects := []string {
		"\\draw (1.2,3) -- (1.1,3.1);",
		"\\draw (1.1,0) rectangle (0,1.1);",
		"\\draw (1.1,1) -- (0,0.1) -- (2.1,2.2) -- cycle;",
		"\\draw (1.1,1.1) circle (1);",
	}
	for i,inst := range tests {
		tikzCode,err := InstToTikz(inst,1.0)
		if err != nil {
			t.Errorf("Failed to generate tikz code with instruction %s: %s",
				inst.ToString(),err.Error())
		}
		if tikzCode != expects[i] {
			t.Errorf("Wrong tikz code generated, expected %s, got %s",
				expects[i],tikzCode)
		}
	}
}
