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

type ParseError struct {
	line   string
	token  string
	reason string
}

func NewParseError(line, token, reason string) *ParseError {
	return &ParseError{line, token, reason}
}

func (e *ParseError) Error() string {
	return e.line + ": " + e.token + " -- " + e.reason
}
