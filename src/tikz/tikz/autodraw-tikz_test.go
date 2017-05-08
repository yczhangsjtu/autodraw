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

import (
	"testing"
	"compiler/instruction"
)

func TestUpdate(t *testing.T) {
	tests := []instruction.Instruction {
		{instruction.LINE_STRIP,[]int16{120,300,110,310}},
		{instruction.CURVE,[]int16{210,110,210,137,115,88,110,160,55,137,5,88,10,110,10,37,5,33,110,60,165,82,115,33,210,110}},
		{instruction.TEXT,[]int16{0,0,256,65,66,67,68,69,70}},
	}
	expect := "\\begin{tikzpicture}\n"+
		"  \\draw (1.2,3) -- (1.1,3.1);\n"+
		"  \\draw (2.1,1.1) .. controls (2.1,1.37) and (1.15,0.88) .. (1.1,1.6) .. controls (0.55,1.37) and (0.05,0.88) .. (0.1,1.1) .. controls (0.1,0.37) and (0.05,0.33) .. (1.1,0.6) .. controls (1.65,0.82) and (1.15,0.33) .. (2.1,1.1);\n"+
		"  \\node[scale=1] at (0,0) {ABCDEF};\n"+
		"\\end{tikzpicture}\n"
	tz := NewTikz()
	for i,_ := range tests {
		err := tz.Update(&tests[i])
		if err != nil {
			t.Errorf("Failed to process instruction %s: %s",
				tests[i].ToString(),err.Error())
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
		{instruction.LINE_STRIP,[]int16{120,300,110,310}},
		{instruction.CURVE,[]int16{210,110,210,137,115,88,110,160,55,137,5,88,10,110,10,37,5,33,110,60,165,82,115,33,210,110}},
		{instruction.TEXT,[]int16{0,0,256,65,66,67,68,69,70}},
	}
	expects := []string {
		"\\draw (1.2,3) -- (1.1,3.1);",
		"\\draw (2.1,1.1) .. controls (2.1,1.37) and (1.15,0.88) .. (1.1,1.6) .. controls (0.55,1.37) and (0.05,0.88) .. (0.1,1.1) .. controls (0.1,0.37) and (0.05,0.33) .. (1.1,0.6) .. controls (1.65,0.82) and (1.15,0.33) .. (2.1,1.1);",
		"\\node[scale=1] at (0,0) {ABCDEF};",
	}
	for i,inst := range tests {
		tikzCode,err := InstToTikz(&inst,1.0)
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
