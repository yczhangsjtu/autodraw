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
package main

import "flag"
import "fmt"
import "os"
import "log"
import "io/ioutil"
import "compiler/instruction"
import "tikz/tikz"

const Version string = "1.0"

var verbose bool
var help bool
var inputFileName string
var outputFileName string

func usage(info string) {
	fmt.Fprintf(os.Stderr, "This is atikz, version %s\n", Version)
	fmt.Fprintln(os.Stderr, info)
	flag.Usage()
}

func main() {

	// Process commandline to initialize settings
	flag.BoolVar(&verbose, "v", false, "verbose level")
	flag.BoolVar(&verbose, "verbose", false, "verbose level")
	flag.BoolVar(&help, "h", false, "show help message")
	flag.BoolVar(&help, "help", false, "show help message")
	flag.StringVar(&outputFileName, "o", "", "output file name")
	flag.StringVar(&outputFileName, "output", "-", "output file name")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s inputFile [options]\noptions:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	args := flag.Args()

	if help {
		usage("")
		return
	}

	if len(args) == 0 {
		usage("Missing input file!")
		return
	}

	inputFileName = args[0]

	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := make([]byte,65536)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	insts,err := instruction.BytesToInstructions(data[:count])
	if err != nil {
		log.Fatal(err)
	}

	tz := tikz.NewTikz()
	for _,inst := range insts {
		err := tz.Update(inst)
		if err != nil {
			log.Fatal(err)
		}
	}

	code,err := tz.GenerateTikzCode()
	if err != nil {
		log.Fatal(err)
	}

	if outputFileName == "" || outputFileName == "-" {
		fmt.Println(code)
	} else {
		err = ioutil.WriteFile(outputFileName, []byte(code), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}
