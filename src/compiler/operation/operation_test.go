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
	for i, v := range operationNames {
		id := int16(i)
		operation = NewOperation(id)
		if operation.Command != id || operation.Name != "" ||
			len(operation.Args) != 0 {
			t.Errorf("NewOperation(%s) failed!", strings.ToUpper(v))
		}
	}
	operation = newLineOperation(NewNumberValues(-100, -120, 100, 120)...)
	if operation.Command != LINE || operation.Name != "" ||
		!reflect.DeepEqual(operation.Args, NewNumberValues(-100, -120, 100, 120)) {
		t.Errorf("NewLineOperation(-100,-120,100,120) failed, got %s",
			operation.ToString())
	}
	operation = newRectOperation(NewNumberValues(-200, -220, 200, 220)...)
	if operation.Command != RECT || operation.Name != "" ||
		!reflect.DeepEqual(operation.Args, NewNumberValues(-200, -220, 200, 220)) {
		t.Errorf("NewRectOperation(-200,-220,200,220) failed, got %s",
			operation.ToString())
	}
	operation = newPolygonOperation(NewNumberValues(-221, 242, 123, 114, 455, 126)...)
	if operation.Command != POLYGON || operation.Name != "" ||
		!reflect.DeepEqual(operation.Args, NewNumberValues(-221, 242, 123, 114, 455, 126)) {
		t.Errorf("NewPolygonOperation(-221,242,123,114,455,126) failed, got %s",
			operation.ToString())
	}
	operation = newSetOperation("s", NewNumberValue(110))
	if operation.Command != SET || operation.Name != "s" ||
		!reflect.DeepEqual(operation.Args, NewNumberValues(110)) {
		t.Errorf("NewSetOperation(110) failed, got %s", operation.ToString())
	}
	operation = newUseOperation("T")
	if operation.Command != USE || operation.Name != "T" ||
		len(operation.Args) != 0 {
		t.Errorf("NewUseOperation(T) failed, got %s", operation.ToString())
	}
	operation = newPushOperation("T")
	if operation.Command != PUSH || operation.Name != "T" ||
		len(operation.Args) != 0 {
		t.Errorf("NewPushOperation(T) failed, got %s", operation.ToString())
	}
	operation = newPopOperation()
	if operation.Command != POP || operation.Name != "" ||
		len(operation.Args) != 0 {
		t.Errorf("NewPopOperation() failed, got %s", operation.ToString())
	}
	operation = newTransformOperation("T", NewNumberValues(110, -10, 10, -90, -200, 0)...)
	if operation.Command != TRANSFORM || operation.Name != "T" ||
		!reflect.DeepEqual(operation.Args, NewNumberValues(110, -10, 10, -90, -200, 0)) {
		t.Errorf("NewTransformOperation(T) failed, got %s", operation.ToString())
	}
	operation = newRotateOperation("T", NewNumberValue(110))
	if operation.Command != ROTATE || operation.Name != "T" ||
		!reflect.DeepEqual(operation.Args, NewNumberValues(110)) {
		t.Errorf("NewRotateOperation(T) failed, got %s", operation.ToString())
	}
	operation = newScaleOperation("T", NewNumberValues(110, -100)...)
	if operation.Command != SCALE || operation.Name != "T" ||
		!reflect.DeepEqual(operation.Args, NewNumberValues(110, -100)) {
		t.Errorf("NewScaleOperation(T) failed, got %s", operation.ToString())
	}
	operation = newTranslateOperation("T", NewNumberValues(110, -100)...)
	if operation.Command != TRANSLATE || operation.Name != "T" ||
		!reflect.DeepEqual(operation.Args, NewNumberValues(110, -100)) {
		t.Errorf("NewTranslateOperation(T) failed, got %s", operation.ToString())
	}
	operation = newDrawOperation("plane")
	if operation.Command != DRAW || operation.Name != "plane" ||
		len(operation.Args) != 0 {
		t.Errorf("NewDrawOperation(plane) failed, got %s", operation.ToString())
	}
	operation = newImportOperation("plane")
	if operation.Command != IMPORT || operation.Name != "plane" ||
		len(operation.Args) != 0 {
		t.Errorf("NewDrawOperation(plane) failed, got %s", operation.ToString())
	}
}

func TestToString(t *testing.T) {
	var operation Operation
	var operationStr string
	var expect string
	for i, v := range operationNames {
		id := int16(i)
		operation = NewOperation(id)
		operationStr = operation.ToString()
		if operationStr != fmt.Sprintf("%s", v) {
			t.Errorf("OperationToString(%s operation) failed!", v)
		}
	}
	operation = newLineOperation(NewNumberValues(-100, -120, 100, 120)...)
	operationStr = operation.ToString()
	expect = fmt.Sprintf("line [%d,%d,%d,%d]", -100, -120, 100, 120)
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"line", expect, operationStr)
	}
	operation = newRectOperation(NewNumberValues(-200, -220, 200, 220)...)
	operationStr = operation.ToString()
	expect = fmt.Sprintf("rect [%d,%d,%d,%d]", -200, -220, 200, 220)
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"rect", expect, operationStr)
	}
	operation = newPolygonOperation(NewNumberValues(-221, 242, 113, 123, 455, 126)...)
	operationStr = operation.ToString()
	expect = fmt.Sprintf("polygon [%d,%d,%d,%d,%d,%d]", -221, 242, 113, 123, 455, 126)
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"polygon", expect, operationStr)
	}
	operation = newOvalOperation(NewNumberValues(-221, 242, 113, 123)...)
	operationStr = operation.ToString()
	expect = fmt.Sprintf("oval [%d,%d,%d,%d]", -221, 242, 113, 123)
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"oval", expect, operationStr)
	}
	operation = newSetOperation("s", NewNumberValue(110))
	operationStr = operation.ToString()
	expect = fmt.Sprintf("set s [%d]", 110)
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"set", expect, operationStr)
	}
	operation = newUseOperation("T")
	operationStr = operation.ToString()
	expect = fmt.Sprintf("use T")
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"use", expect, operationStr)
	}
	operation = newPushOperation("T")
	operationStr = operation.ToString()
	expect = fmt.Sprintf("push T")
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"push", expect, operationStr)
	}
	operation = newPopOperation()
	operationStr = operation.ToString()
	expect = fmt.Sprintf("pop")
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"pop", expect, operationStr)
	}
	operation = newTransformOperation("T", NewNumberValues(110, -10, -10, 110, 0, 0)...)
	operationStr = operation.ToString()
	expect = fmt.Sprintf("transform T [%d,%d,%d,%d,%d,%d]", 110, -10, -10, 110, 0, 0)
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"transform", expect, operationStr)
	}
	operation = newRotateOperation("T", NewNumberValue(110))
	operationStr = operation.ToString()
	expect = fmt.Sprintf("rotate T [%d]", 110)
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"rotate", expect, operationStr)
	}
	operation = newScaleOperation("T", NewNumberValues(110, -100)...)
	operationStr = operation.ToString()
	expect = fmt.Sprintf("scale T [%d,%d]", 110, -100)
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"scale", expect, operationStr)
	}
	operation = newTranslateOperation("T", NewNumberValues(110, -100)...)
	operationStr = operation.ToString()
	expect = fmt.Sprintf("translate T [%d,%d]", 110, -100)
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"translate", expect, operationStr)
	}
	operation = newDrawOperation("plane")
	operationStr = operation.ToString()
	expect = fmt.Sprintf("draw plane")
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"draw", expect, operationStr)
	}
	operation = newImportOperation("plane")
	operationStr = operation.ToString()
	expect = fmt.Sprintf("import plane")
	if operationStr != expect {
		t.Errorf("OperationToString(%s operation) failed! Expect %s, got %s",
			"import", expect, operationStr)
	}
}

///////////////////////////////////////////////////////////////////////////////
// New Operations for convenience //////////////////////////////////////////////
func newLineOperation(coords ...Value) Operation {
	if len(coords) == int(ExpectArgNum(LINE)) {
		return newOperationTypeDrawFixed(LINE, coords...)
	} else {
		return NewOperation(UNDEFINED)
	}
}

func newRectOperation(coords ...Value) Operation {
	if len(coords) == int(ExpectArgNum(RECT)) {
		return newOperationTypeDrawFixed(RECT, coords...)
	} else {
		return NewOperation(UNDEFINED)
	}
}

func newOvalOperation(coords ...Value) Operation {
	if len(coords) == int(ExpectArgNum(OVAL)) {
		return newOperationTypeDrawFixed(OVAL, coords...)
	} else {
		return NewOperation(UNDEFINED)
	}
}

func newPolygonOperation(coords ...Value) Operation {
	if len(coords) >= 4 && len(coords)%2 == 0 {
		return newOperationTypeDrawNondetermined(POLYGON, coords...)
	} else {
		return NewOperation(UNDEFINED)
	}
}

func newSetOperation(name string, v Value) Operation {
	return newOperationTypeAssign(SET, name, v)
}

func newRotateOperation(name string, v Value) Operation {
	return newOperationTypeAssign(ROTATE, name, v)
}

func newTransformOperation(name string, coords ...Value) Operation {
	if len(coords) == int(ExpectArgNum(TRANSFORM)) {
		return newOperationTypeAssign(TRANSFORM, name, coords...)
	} else {
		return NewOperation(UNDEFINED)
	}
}

func newScaleOperation(name string, coords ...Value) Operation {
	if len(coords) == 2 {
		return newOperationTypeAssign(SCALE, name, coords...)
	} else {
		return NewOperation(UNDEFINED)
	}
}

func newTranslateOperation(name string, coords ...Value) Operation {
	if len(coords) == 2 {
		return newOperationTypeAssign(TRANSLATE, name, coords...)
	} else {
		return NewOperation(UNDEFINED)
	}
}

func newPopOperation() Operation {
	return newOperationTypeSingle(POP)
}

func newPushOperation(name string) Operation {
	return newOperationTypeState(PUSH, name)
}

func newDrawOperation(name string) Operation {
	return newOperationTypeState(DRAW, name)
}

func newImportOperation(path string) Operation {
	return newOperationTypeState(IMPORT, path)
}

func newUseOperation(name string) Operation {
	return newOperationTypeState(USE, name)
}

func newOperationTypeDrawFixed(op int16, args ...Value) Operation {
	operation := NewOperation(op)
	operation.Args = args
	return operation
}

func newOperationTypeDrawNondetermined(op int16, args ...Value) Operation {
	operation := NewOperation(op)
	operation.Args = args
	return operation
}

func newOperationTypeAssign(op int16, name string, args ...Value) Operation {
	if !ValidName(name) {
		return NewOperation(UNDEFINED)
	}
	operation := NewOperation(op)
	operation.Name = name
	operation.Args = args
	return operation
}

func newOperationTypeSingle(op int16) Operation {
	operation := NewOperation(op)
	return operation
}

func newOperationTypeState(op int16, name string) Operation {
	if !ValidName(name) {
		return NewOperation(UNDEFINED)
	}
	operation := NewOperation(op)
	operation.Name = name
	return operation
}
///////////////////////////////////////////////////////////////////////////////
