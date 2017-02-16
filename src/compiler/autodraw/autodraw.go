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
import "os"
import "log"
import "bufio"
import "strings"
import "io/ioutil"
import "compiler/fsm"
import "compiler/operation"
//import "compiler/instruction"

const Version string = "1.0"

var verbose bool
var help bool
var inputFileName string
var outputFileName string

func usage(info string) {
	fmt.Fprintf(os.Stderr,"This is autodraw, version %s\n",Version)
	fmt.Fprintln(os.Stderr,info)
	flag.Usage()
}

func main() {

	// Process commandline to initialize settings
	flag.BoolVar(&verbose,"v",false,"verbose level")
	flag.BoolVar(&verbose,"verbose",false,"verbose level")
	flag.BoolVar(&help,"h",false,"show help message")
	flag.BoolVar(&help,"help",false,"show help message")
	flag.StringVar(&outputFileName,"o","","output file name")
	flag.StringVar(&outputFileName,"output","a.anm","output file name")
	flag.Usage = func (){
		fmt.Fprintf(os.Stderr, "Usage: %s inputFile [options]\noptions:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	args := flag.Args()

	if help {
		usage("")
		return
	}

	operation.Verbose = verbose

	if len(args) == 0 {
		usage("Missing input file!")
		return
	}

	inputFileName = args[0]

	file,err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineno := 1
	compiler := fsm.NewFSM()
	compiler.Verbose = verbose

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Trim(line," ")
		if line == "" {
			continue
		}
		parser := operation.NewLineParser()
		oper,err := parser.ParseLine(line)
		if err != nil {
			log.Fatal(err)
		}
		compiler.Update(oper)

		lineno++
	}

	err = ioutil.WriteFile(outputFileName, compiler.DumpInstructions(), 0644)
}
