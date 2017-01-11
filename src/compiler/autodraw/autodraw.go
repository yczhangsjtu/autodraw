package main

import "flag"
import "fmt"
import "compiler/operation"

var verbose bool
var help bool

func main() {

	// Process commandline to initialize settings
	flag.BoolVar(&verbose,"v",false,"verbose level")
	flag.BoolVar(&verbose,"verbose",false,"verbose level")
	flag.BoolVar(&help,"h",false,"show help message")
	flag.BoolVar(&help,"help",false,"show help message")

	flag.Parse()
	fmt.Println(verbose)
	fmt.Println(help)

	oper := operation.Parse("poly abc")
	operation.OperationPrint(oper)
	fmt.Println("")

}
