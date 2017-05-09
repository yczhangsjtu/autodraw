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

func ParseLine(line string) (Operation, error) {
	variableFinder := GetVariableFinderRegexp()
	for i,_ := range operationRegexp {
		op := int16(i)
		pattern := GetOperationRegexp(op)
		if pattern.MatchString(line) {
			result := pattern.FindStringSubmatch(line)
			oper := NewOperation(op)
			switch op {
				case LINE:
					fallthrough
				case OVAL:
					fallthrough
				case RECT:
					fallthrough
				case POLYGON:
					result := variableFinder.FindAllString(result[1],-1)
					values := NewValues(result...)
					oper.Args = values
					return oper,nil
				case TEXT:
					stringValue := result[2]
					result := variableFinder.FindAllString(result[1],-1)
					values := append(NewValues(result...),NewValue(stringValue))
					oper.Args = values
					return oper,nil
				case USE:
					fallthrough
				case PUSH:
					fallthrough
				case IMPORT:
					fallthrough
				case BEGIN:
					fallthrough
				case DRAW:
					oper.Name = result[1]
					return oper,nil
				case SET:
					fallthrough
				case ROTATE:
					oper.Name = result[1]
					values := NewValues(result[2])
					oper.Args = values
					return oper,nil
				case TRANSFORM:
					oper.Name = result[1]
					result := variableFinder.FindAllString(result[2],-1)
					values := NewValues(result...)
					oper.Args = values
					return oper,nil
				case SCALE:
					fallthrough
				case TRANSLATE:
					oper.Name = result[1]
					result := variableFinder.FindAllString(result[2],-1)
					values := NewValues(result...)
					oper.Args = values
					return oper,nil
				case END:
					fallthrough
				case POP:
					return oper,nil
				default:
					return NewOperation(UNDEFINED),nil
			}
		}
	}
	return NewOperation(UNDEFINED),NewParseError(line,"","invalid syntax")
}
