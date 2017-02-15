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
		if operation.Command != id || operation.Name != "" ||
		  len(operation.Args) != 0 {
			t.Errorf("NewOperation(%s) failed!",strings.ToUpper(v))
		}
	}
	operation = NewLineOperation(NewNumberValues(-100,-120,100,120)...)
	if operation.Command != LINE || operation.Name != "" ||
		 !reflect.DeepEqual(operation.Args,NewNumberValues(-100,-120,100,120)) {
		t.Errorf("NewLineOperation(-100,-120,100,120) failed, got %s",
			operation.ToString())
	}
	operation = NewRectOperation(NewNumberValues(-200,-220,200,220)...)
	if operation.Command != RECT || operation.Name != "" ||
		 !reflect.DeepEqual(operation.Args,NewNumberValues(-200,-220,200,220)) {
		t.Errorf("NewRectOperation(-200,-220,200,220) failed, got %s",
			operation.ToString())
	}
	operation = NewCircleOperation(NewNumberValues(-220,200,120)...)
	if operation.Command != CIRCLE || operation.Name != "" ||
		 !reflect.DeepEqual(operation.Args,NewNumberValues(-220,200,120)) {
		t.Errorf("NewCircleOperation(-220,200,120) failed, got %s",
			operation.ToString())
	}
	operation = NewPolygonOperation(NewNumberValues(-221,242,123,114,455,126)...)
	if operation.Command != POLYGON || operation.Name != "" ||
		 !reflect.DeepEqual(operation.Args, NewNumberValues(-221,242,123,114,455,126)) {
		t.Errorf("NewPolygonOperation(-221,242,123,114,455,126) failed, got %s",
			operation.ToString())
	}
	operation = NewSetOperation("scale",NewNumberValue(110))
	if operation.Command != SET || operation.Name != "scale" ||
	   !reflect.DeepEqual(operation.Args, NewNumberValues(110)) {
		t.Errorf("NewSetOperation(110) failed, got %s",operation.ToString())
	}
	operation = NewUseOperation("T")
	if operation.Command != USE || operation.Name != "T" ||
	   len(operation.Args)!=0 {
		t.Errorf("NewUseOperation(T) failed, got %s",operation.ToString())
	}
	operation = NewPushOperation("T")
	if operation.Command != PUSH || operation.Name != "T" ||
	   len(operation.Args)!=0 {
		t.Errorf("NewPushOperation(T) failed, got %s",operation.ToString())
	}
	operation = NewPopOperation()
	if operation.Command != POP || operation.Name != "" ||
		 len(operation.Args)!=0 {
		t.Errorf("NewPopOperation() failed, got %s",operation.ToString())
	}
	operation = NewTransformOperation("T",NewNumberValues(110,-10,10,-90,-200,0)...)
	if operation.Command != TRANSFORM || operation.Name != "T" ||
	   !reflect.DeepEqual(operation.Args, NewNumberValues(110,-10,10,-90,-200,0)) {
		t.Errorf("NewTransformOperation(T) failed, got %s",operation.ToString())
	}
	operation = NewDrawOperation("plane")
	if operation.Command != DRAW || operation.Name != "plane" ||
	   len(operation.Args)!=0 {
		t.Errorf("NewDrawOperation(plane) failed, got %s",operation.ToString())
	}
	operation = NewImportOperation("plane")
	if operation.Command != IMPORT || operation.Name != "plane" ||
	   len(operation.Args)!=0 {
		t.Errorf("NewDrawOperation(plane) failed, got %s",operation.ToString())
	}
}

func TestToString(t *testing.T) {
	var operation Operation
	var operationStr string
	var expect string
	for i,v := range OperationNames {
		id := int16(i)
		operation = NewOperation(id)
		operationStr = operation.ToString()
		if operationStr != fmt.Sprintf("%s",v) {
			t.Errorf("OperationToString(%s operation) failed!",v)
		}
	}
	operation = NewLineOperation(NewNumberValues(-100,-120,100,120)...)
	operationStr = operation.ToString()
	expect = fmt.Sprintf("line [%d,%d,%d,%d]",-100,-120,100,120)
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"line",expect,operationStr)
	}
	operation = NewRectOperation(NewNumberValues(-200,-220,200,220)...)
	operationStr = operation.ToString()
	expect = fmt.Sprintf("rect [%d,%d,%d,%d]",-200,-220,200,220)
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"rect",expect,operationStr)
	}
	operation = NewCircleOperation(NewNumberValues(-220,200,120)...)
	operationStr = operation.ToString()
	expect = fmt.Sprintf("circle [%d,%d,%d]",-220,200,120)
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"circle",expect,operationStr)
	}
	operation = NewPolygonOperation(NewNumberValues(-221,242,113,123,455,126)...)
	operationStr = operation.ToString()
	expect = fmt.Sprintf("polygon [%d,%d,%d,%d,%d,%d]",-221,242,113,123,455,126)
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"polygon",expect,operationStr)
	}
	operation = NewOvalOperation(NewNumberValues(-221,242,113,123,455)...)
	operationStr = operation.ToString()
	expect = fmt.Sprintf("oval [%d,%d,%d,%d,%d]",-221,242,113,123,455)
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"oval",expect,operationStr)
	}
	operation = NewSetOperation("scale",NewNumberValue(110))
	operationStr = operation.ToString()
	expect = fmt.Sprintf("set scale [%d]",110)
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"set",expect,operationStr)
	}
	operation = NewUseOperation("T")
	operationStr = operation.ToString()
	expect = fmt.Sprintf("use T")
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"use",expect,operationStr)
	}
	operation = NewPushOperation("T")
	operationStr = operation.ToString()
	expect = fmt.Sprintf("push T")
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"push",expect,operationStr)
	}
	operation = NewPopOperation()
	operationStr = operation.ToString()
	expect = fmt.Sprintf("pop")
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"pop",expect,operationStr)
	}
	operation = NewTransformOperation("T",NewNumberValues(110,-10,-10,110,0,0)...)
	operationStr = operation.ToString()
	expect = fmt.Sprintf("transform T [%d,%d,%d,%d,%d,%d]",110,-10,-10,110,0,0)
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"transform",expect,operationStr)
	}
	operation = NewDrawOperation("plane")
	operationStr = operation.ToString()
	expect = fmt.Sprintf("draw plane")
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"draw",expect,operationStr)
	}
	operation = NewImportOperation("plane")
	operationStr = operation.ToString()
	expect = fmt.Sprintf("import plane")
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"import",expect,operationStr)
	}
}
