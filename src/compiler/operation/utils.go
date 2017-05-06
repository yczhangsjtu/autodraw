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
	"unicode"
	"strconv"
)

func ValidName(name string) bool {
	if len(name) == 0 {
		return false
	}
	for i, c := range name {
		if !unicode.IsLetter(c) {
			if i == 0 {
				return false
			}
			if !unicode.IsDigit(c) && c != '-' && c != '.' && c != '_' {
				return false
			}
		}
	}
	return true
}

func tokenIdentify(token string) int16 {
	_, ok := GetCommand(token)
	if ok {
		return COMMAND
	}
	if ValidName(token) {
		return NAME
	}
	_, err := strconv.ParseInt(token, 10, 16)
	if err != nil {
		return INVALID
	}
	return NUMBER
}
