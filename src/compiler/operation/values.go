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

func ExpectName(op int16) bool {
	return expectName[GetType(op)]
}

func ExpectArgNum(op int16) int {
	return expectArgNum[op]
}

func ExpectArgs(op int16) bool {
	return expectArgs[GetType(op)]
}

func FinalArgNum(op int16) int {
	return finalArgNum[op]
}

func NeedArgNum(op int16) bool {
	return needArgNum[GetType(op)]
}

func NewNumberValue(x int16) Value {
	return Value{INTEGER, "", x, nil}
}

func NewNumberValues(args ...int16) []Value {
	ret := make([]Value, len(args))
	for i, v := range args {
		ret[i] = NewNumberValue(v)
	}
	return ret
}

func GetCommand(name string) (int16,bool) {
	ret,ok := operationNameMap[name]
	return ret,ok
}

func GetName(op int16) string {
	if int(op) >= len(operationNames) {
		return "undefined"
	}
	return operationNames[op]
}

func GetType(op int16) int16 {
	if int(op) >= len(operationTypes) {
		return UNDEFINED
	}
	return operationTypes[op]
}

func GetOperationNum() int {
	return len(operationTypes)
}
