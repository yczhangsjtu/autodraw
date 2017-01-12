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
Package autodraw/main is the executable program for the autodraw language
compiler
*/
package main

import "flag"
import "fmt"
import "compiler/operation"

const Version string = "1.0"

var verbose bool
var help bool

func usage(info string) {
	fmt.Println(info)
	flag.Usage()
}

func main() {

	// Process commandline to initialize settings
	flag.BoolVar(&verbose,"v",false,"verbose level")
	flag.BoolVar(&verbose,"verbose",false,"verbose level")
	flag.BoolVar(&help,"h",false,"show help message")
	flag.BoolVar(&help,"help",false,"show help message")

	flag.Parse()

	if help {
		usage(fmt.Sprintf("This is autodraw, version %s",Version))
		return
	}

	operation.Verbose = verbose

	oper := operation.Parse("polygon 0.0 0.0 1.0 0.0 1.0 1.1 0.0 1.0")
	operation.OperationPrint(oper)
	fmt.Println("")

}
