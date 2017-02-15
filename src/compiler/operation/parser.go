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

import "strings"
import "strconv"

const (
	INVALID int16 = iota
	COMMAND
	NAME
	NUMBER
)

const (
	NEED_COMMAND int16 = iota
	NEED_NAME
	NEED_VALUE
	FINISH
	ERROR
)

type LineParser struct {
	line string

	expectName bool
	expectArgs int16
	needArgNum bool

	command int16
	name string
	args []Value

	state int16
}

type ParseError struct {
	line string
	token string
	reason string
}

func NewParseError(line, token, reason string) *ParseError {
	return &ParseError{line,token,reason}
}

func (e *ParseError) Error() string {
	return e.line + ": " + e.token + " -- " + e.reason
}

func StringToValue(token string) (Value,error) {
	if ValidName(token) {
		return NewVariableValue(token),nil
	}
	i,err := strconv.ParseInt(token,10,16)
	if err != nil {
		return Value{NAN,"",0,nil},err
	}
	return NewNumberValue(int16(i)),nil
}

func TokenIdentify(token string) int16 {
	_,ok := OperationNameMap[token]
	if ok {
		return COMMAND
	}
	if ValidName(token) {
		return NAME
	}
	_,err := strconv.ParseInt(token,10,16)
	if err != nil {
		return INVALID
	}
	return NUMBER
}

func NewLineParser() LineParser {
	lineParser := new(LineParser)
	lineParser.state = NEED_COMMAND
	return *lineParser
}

func (parser *LineParser) Update(token string) error {
	token = strings.Trim(token," ")
	if token == "" {
		return nil
	}
	tokenType := TokenIdentify(token)
	if tokenType == INVALID {
		return NewParseError(parser.line,token,"invalid token")
	}
	switch parser.state {
	case NEED_COMMAND:
		if tokenType == COMMAND {
			parser.command = OperationNameMap[token]
			parser.expectName = ExpectName(parser.command)
			parser.needArgNum = NeedArgNum(parser.command)
			parser.expectArgs = ExpectArgs(parser.command)
			if parser.expectName {
				parser.state = NEED_NAME
			} else if parser.expectArgs > 0 || parser.needArgNum {
				parser.state = NEED_VALUE
			} else {
				parser.state = FINISH
			}
		} else {
			parser.state = ERROR
			return NewParseError(parser.line,token,"expect command")
		}
	case NEED_NAME:
		if tokenType == COMMAND {
			parser.state = ERROR
			return NewParseError(parser.line,token,"is reserved")
		} else if tokenType == NUMBER {
			parser.state = ERROR
			return NewParseError(parser.line,token,"expect name")
		} else if tokenType == NAME {
			parser.name = token
			if parser.expectArgs > 0 || parser.needArgNum {
				parser.state = NEED_VALUE
			} else {
				parser.state = FINISH
			}
		} else {
			parser.state = ERROR
			return NewParseError(parser.line,token,"unknown token")
		}
	case NEED_VALUE:
		if tokenType == COMMAND {
			parser.state = ERROR
			return NewParseError(parser.line,token,"is reserved")
		} else if tokenType == NUMBER {
			number,_ := strconv.ParseInt(token,10,16)
			parser.args = append(parser.args,NewNumberValue(int16(number)))
			if int16(len(parser.args)) == parser.expectArgs {
				parser.state = FINISH
			} else if int16(len(parser.args)) > parser.expectArgs && !parser.needArgNum {
				parser.state = ERROR
				return NewParseError(parser.line,token,"too many arguments")
			}
		} else if tokenType == NAME {
			parser.args = append(parser.args,NewVariableValue(token))
			if int16(len(parser.args)) == parser.expectArgs {
				parser.state = FINISH
			} else if int16(len(parser.args)) > parser.expectArgs && !parser.needArgNum {
				parser.state = ERROR
				return NewParseError(parser.line,token,"too many arguments")
			}
		} else {
			parser.state = ERROR
			return NewParseError(parser.line,token,"unknown token")
		}
	case FINISH:
		parser.state = ERROR
		return NewParseError(parser.line,token,"too many arguments")
	case ERROR:
		return NewParseError(parser.line,token,"parser in invalid state")
	}
	return nil
}

func (parser *LineParser) Digest() (Operation,error) {
	if !parser.needArgNum && parser.state != FINISH {
		return NewOperation(UNDEFINED),NewParseError(parser.line,"$","not finished")
	} else {
		op := NewOperation(parser.command)
		if parser.expectName {
			op.Name = parser.name
		}
		if parser.expectArgs > 0 || parser.needArgNum {
			op.Args = parser.args
		}
		return op,nil
	}
}

func (parser *LineParser) ParseLine(line string) (Operation,error) {
	parser.line = line
	line = strings.Trim(line," ")
	tokens := strings.Split(line," ")

	for _,token := range tokens {
		err := parser.Update(token)
		if err != nil {
			return NewOperation(UNDEFINED),err
		}
	}
	return parser.Digest()
}
