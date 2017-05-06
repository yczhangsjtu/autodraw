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

/*
Package fsm implements a simple Finite State Machine which takes operations as
inputs and updates its state. When the input operations are finished, the
generated instructions can be dumped to byte string.
*/
package fsm

type FSMError struct {
	oper   string
	reason string
}

type VartableError struct {
	reason string
}

type ArgError struct {
	reason string
}

func NewVartableError(reason string) *VartableError {
	e := VartableError{reason}
	return &e
}

func NewArgError(reason string) *ArgError {
	e := ArgError{reason}
	return &e
}

func NewFSMError(oper string, reason string) *FSMError {
	e := FSMError{oper, reason}
	return &e
}

func (e *FSMError) Error() string {
	return "FSM error: " + e.oper + ": " + e.reason
}

func (e *VartableError) Error() string {
	return "Vartable error: " + e.reason
}

func (e *ArgError) Error() string {
	return "Argument error: " + e.reason
}
