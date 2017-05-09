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
	"regexp"
)

func GetPattern(op int16) string {
	if int(op) >= len(operationNames) {
		return ""
	}
	return operationPatterns[op]
}

func GetOperationRegexp(op int16) *regexp.Regexp {
	if int(op) >= len(operationRegexp) {
		return nil
	}
	if operationRegexp[op] == nil {
		var err error
		operationRegexp[op],err = regexp.Compile(operationPatterns[op])
		if err != nil {
			return nil
		}
	}
	return operationRegexp[op]
}

func GetVariableRegexp() *regexp.Regexp {
	if variableRegexp == nil {
		var err error
		variableRegexp,err = regexp.Compile(variablePattern)
		if err != nil {
			return nil
		}
	}
	return variableRegexp
}

func GetVariableFinderRegexp() *regexp.Regexp {
	if variableFinderRegexp == nil {
		var err error
		variableFinderRegexp,err = regexp.Compile(variableFinderPattern)
		if err != nil {
			return nil
		}
	}
	return variableFinderRegexp
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

