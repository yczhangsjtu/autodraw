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
package fsm

import "testing"
import "compiler/operation"

func TestFSMUpdate(t *testing.T) {
	fsm := NewFSM()
	tests := []string{
		"line 120 300 110 310",
		"rect 110 0 0 110",
		"polygon 110 100 0 10 210 220",
		"circle 110 110 100",
		"oval 110 110 100 50 120",
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
		"end",
		"set Z X",
		"set X Y",
		"set U W",
		"set V U",
		"push T",
		"line 120 300 110 310",
		"rect 110 0 0 110",
		"polygon 110 100 0 10 210 220",
		"circle 110 110 100",
		"oval 110 110 100 50 120",
		"pop",
		"use T",
		"line 120 300 110 310",
		"use Q",
		"rect 110 0 0 110",
		"use P",
		"draw subfigure",
		"polygon 110 100 0 10 210 220",
		"circle 110 110 100",
		"oval 110 110 100 50 120",
	}
	results := map[string]int16{
		"Alice": 110, "Bob": -10, "Carror": 110,
	}
	transforms := []string{
		"T", "P", "Q", "U", "V", "W", "X", "Y", "Z",
	}
	unexpect := []string {
		"M",
	}
	for _, line := range tests {
		parser := operation.NewLineParser()
		oper, err := parser.ParseLine(line)
		if err != nil {
			t.Errorf(err.Error())
		}
		err = fsm.Update(oper)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	for k, v := range results {
		value, ok := fsm.Lookup(k)
		if !ok {
			t.Errorf("%s not found", k)
		}
		if value.Number != v {
			t.Errorf("Expect %s = %d, got %d", k, v, value.Number)
		}
	}
	for _, v := range unexpect {
		_, ok := fsm.Lookup(v)
		if ok {
			t.Errorf("%s should not be found", v)
		}
	}
	for _, tf := range transforms {
		value, ok := fsm.Lookup(tf)
		if !ok {
			t.Errorf("transform %s not found", tf)
		}
		if value.Type != operation.TRANSFORMER {
			t.Errorf("%s is not transform", tf)
		}
	}
}
