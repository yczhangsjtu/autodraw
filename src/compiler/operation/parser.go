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

//import (
//	"strings"
//	"strconv"
//)
//
//type LineParser struct {
//	line string
//
//	expectName   bool
//	expectArgs   bool
//	expectArgNum int
//	undetermined   bool
//
//	command int16
//	name    string
//	args    []Value
//
//	state int16
//}
//
//func NewLineParser() LineParser {
//	lineParser := new(LineParser)
//	lineParser.Initialize()
//	return *lineParser
//}
//
//func (parser *LineParser) Initialize() {
//	parser.line = ""
//	parser.expectName = false
//	parser.expectArgs = false
//	parser.expectArgNum = 0
//	parser.undetermined = false
//	parser.command = UNDEFINED
//	parser.name = ""
//	parser.args = []Value{}
//	parser.state = NEED_COMMAND
//}
//
//func (parser *LineParser) Error(token string, reason string) *ParseError {
//	parser.state = ERROR
//	return NewParseError(parser.line, token, reason)
//}
//
//func (parser *LineParser) appendNumberArg(arg int16) {
//	parser.args = append(parser.args, NewNumberValue(int16(arg)))
//}
//
//func (parser *LineParser) appendVariableArg(token string) {
//	parser.args = append(parser.args, NewVariableValue(token))
//}
//
//func (parser *LineParser) getArgNum() int {
//	return len(parser.args)
//}
//
//func (parser *LineParser) Update(token string) error {
//	token = strings.Trim(token, " ")
//	if token == "" {
//		return nil
//	}
//	tokenType := tokenIdentify(token)
//	if tokenType == INVALID {
//		return NewParseError(parser.line, token, "invalid token")
//	}
//	switch parser.state {
//	case NEED_COMMAND:
//		if tokenType == COMMAND {
//			parser.command,_ = GetCommand(token)
//			parser.expectName = ExpectName(parser.command)
//			parser.undetermined = NeedArgNum(parser.command)
//			parser.expectArgNum = ExpectArgNum(parser.command)
//			parser.expectArgs = ExpectArgs(parser.command)
//			if parser.expectName {
//				parser.state = NEED_NAME
//			} else if parser.expectArgs {
//				parser.state = NEED_VALUE
//			} else {
//				parser.state = FINISH
//			}
//		} else {
//			return parser.Error(token, "expecting command")
//		}
//	case NEED_NAME:
//		if tokenType == COMMAND {
//			return parser.Error(token, "is reserved")
//		} else if tokenType == NUMBER {
//			return parser.Error(token, "expecting name")
//		} else if tokenType == NAME {
//			parser.name = token
//			if parser.expectArgs {
//				parser.state = NEED_VALUE
//			} else {
//				parser.state = FINISH
//			}
//		} else {
//			return parser.Error(token, "unknown token")
//		}
//	case NEED_VALUE:
//		if tokenType == COMMAND {
//			return parser.Error(token, "is reserved")
//		} else if tokenType == NUMBER {
//			number, _ := strconv.ParseInt(token, 10, 16)
//			parser.appendNumberArg(int16(number))
//			if parser.getArgNum() == parser.expectArgNum {
//				parser.state = FINISH
//			} else if parser.getArgNum() > parser.expectArgNum {
//				if !parser.undetermined {
//					return parser.Error(token, "too many arguments")
//				}
//			}
//		} else if tokenType == NAME {
//			parser.appendVariableArg(token)
//			if parser.getArgNum() == parser.expectArgNum {
//				parser.state = FINISH
//			} else if parser.getArgNum() > parser.expectArgNum {
//				if !parser.undetermined {
//					return parser.Error(token, "too many arguments")
//				}
//			}
//		} else {
//			return parser.Error(token, "unknown token")
//		}
//	case FINISH:
//		return parser.Error(token, "too many arguments")
//	case ERROR:
//		return parser.Error(token, "parser in invalid state")
//	}
//	return nil
//}
//
//func (parser *LineParser) Digest() (Operation, error) {
//	if !parser.undetermined && parser.state != FINISH {
//		return NewOperation(UNDEFINED), parser.Error("$", "not finished")
//	} else {
//		op := NewOperation(parser.command)
//		if parser.expectName {
//			op.Name = parser.name
//		}
//		if parser.expectArgs {
//			op.Args = parser.args
//		}
//		parser.Initialize()
//		return op, nil
//	}
//}
//
//func Parse(line string) (Operation, error) {
//	parser := NewLineParser()
//	parser.line = line
//	line = strings.Trim(line, " ")
//	tokens := strings.Split(line, " ")
//
//	if len(tokens) == 0 {
//		return NewOperation(UNDEFINED), NewParseError("", "", "empty line")
//	}
//
//	for _, token := range tokens {
//		err := parser.Update(token)
//		if err != nil {
//			return NewOperation(UNDEFINED), err
//		}
//	}
//	return parser.Digest()
//}

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
