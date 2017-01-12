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
import "strings"
import "reflect"

func TestNewOperation(t *testing.T) {
	var operation Operation
	for i,v := range OperationNames {
		id := int16(i)
		operation = NewOperation(id)
		if operation.Op != id || operation.Detail.Op != UNDEFINED{
			t.Errorf("NewOperation(%s) failed!",strings.ToUpper(v))
		}
	}
	operation = NewLineOperation(-1.0,-1.2,1.0,1.2)
	if operation.Op != LINE || operation.Detail.Op != LINE ||
		 reflect.DeepEqual(operation.Detail.Args,[4]int16{-100,-120,100,120}) {
		t.Errorf("NewLineOperation(-1.0,-1.2,1.0,1.2) failed")
	}
	operation = NewRectOperation(-2.0,-2.2,2.0,2.2)
	if operation.Op != RECT || operation.Detail.Op != RECT ||
		 reflect.DeepEqual(operation.Detail.Args,[4]int16{-200,-220,200,220}) {
		t.Errorf("NewRectOperation(-2.0,-2.2,2.0,2.2) failed")
	}
	operation = NewCircleOperation(-2.2,2.0,1.2)
	if operation.Op != CIRCLE || operation.Detail.Op != CIRCLE ||
		 reflect.DeepEqual(operation.Detail.Args,[3]int16{-220,200,120}) {
		t.Errorf("NewCircleOperation(-2.2,2.0,1.2) failed")
	}
	operation = NewPolygonOperation(-2.21,2.42,1.23,1.14,4.55,1.26)
	if operation.Op != POLYGON || operation.Detail.Op != POLYGON ||
		 reflect.DeepEqual(operation.Detail.Args,
		 [6]int16{-221,242,123,114,455,126}) {
		t.Errorf("NewPolygonOperation(-2.21,2.42,1.23,1.14,4.55,1.26) failed")
	}
	operation = NewSetOperation("scale",1.1)
	if operation.Op != SET || operation.Detail.Op != SET ||
		 operation.Name != "scale" || reflect.DeepEqual(operation.Detail.Args,
		 [1]int16{110}) {
		t.Errorf("NewSetOperation(1.1) failed")
	}
	operation = NewUseOperation("T",NewLineInstruction(0.0,0.0,-1.0,-1.0))
	if operation.Op != USE || operation.Detail.Op != LINE ||
		 operation.Name != "T" || !InstructionEqual(operation.Detail,
			 NewLineInstruction(0.0,0.0,-1.0,-1.0)) {
		t.Errorf("NewUseOperation(T,line) failed")
	}
}

func TestToString(t *testing.T) {
	var operation Operation
	var operationStr string
	var expect string
	var zero float64 = 0.0
	for i,v := range OperationNames {
		id := int16(i)
		operation = NewOperation(id)
		operationStr = OperationToString(operation)
		if !HasTransform(id) {
			if operationStr != fmt.Sprintf("%s undefined",v) {
				t.Errorf("OperationToString(%s operation) failed!",v)
			}
		} else {
			if operationStr != fmt.Sprintf("%s [[%f,%f;%f,%f],[%f,%f]] undefined",
																		v,zero,zero,zero,zero,zero,zero) {
				t.Errorf("OperationToString(%s operation) failed!",v)
			}
		}
	}
	operation = NewLineOperation(-1.0,-1.2,1.0,1.2)
	operationStr = OperationToString(operation)
	expect = fmt.Sprintf("line line %f %f %f %f",
			float64(-1.0),float64(-1.2),float64(1.0),float64(1.2))
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"line",expect,operationStr)
	}
	operation = NewRectOperation(-2.0,-2.2,2.0,2.2)
	operationStr = OperationToString(operation)
	expect = fmt.Sprintf("rect rect %f %f %f %f",
			float64(-2.0),float64(-2.2),float64(2.0),float64(2.2))
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"rect",expect,operationStr)
	}
	operation = NewCircleOperation(-2.2,2.0,1.2)
	operationStr = OperationToString(operation)
	expect = fmt.Sprintf("circle circle %f %f %f",
			float64(-2.2),float64(2.0),float64(1.2))
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"circle",expect,operationStr)
	}
	operation = NewPolygonOperation(-2.21,2.42,1.13,1.23,4.55,1.26)
	operationStr = OperationToString(operation)
	expect = fmt.Sprintf("polygon polygon %f %f %f %f %f %f",
			float64(-2.21),float64(2.42),float64(1.13),
			float64(1.23),float64(4.55),float64(1.26))
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"polygon",expect,operationStr)
	}
	operation = NewOvalOperation(-2.21,2.42,1.13,1.23,4.55)
	operationStr = OperationToString(operation)
	expect = fmt.Sprintf("oval oval %f %f %f %f %f",
			float64(-2.21),float64(2.42),float64(1.13),float64(1.23),float64(4.55))
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"oval",expect,operationStr)
	}
	operation = NewSetOperation("scale",1.1)
	operationStr = OperationToString(operation)
	expect = fmt.Sprintf("set scale set %f",float64(1.1))
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"set",expect,operationStr)
	}
	operation = NewUseOperation("T",NewLineInstruction(0.0,1.1,-0.1,-1.0))
	operationStr = OperationToString(operation)
	expect = fmt.Sprintf("use T line %f %f %f %f",
			float64(0.0),float64(1.1),float64(-0.1),float64(-1.0))
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"use",expect,operationStr)
	}
}
