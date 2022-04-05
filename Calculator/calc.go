package main

import (
	"fmt"
	"os"
)

func restart() {
	procAttr := new(os.ProcAttr)
	procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
	os.StartProcess(os.Args[0], []string{"", "test"}, procAttr)
}

func main() {
	var controlSum bool = false
	var firstFloat float64
	var secondFloat float64
	var operation string
	var tempString string
	fmt.Println("First var: ")
	fmt.Scanln(&firstFloat)
	fmt.Println("Second var: ")
	fmt.Scanln(&secondFloat)
	fmt.Println("Which operation do you want to perform with your variables? (+, -, == , /, *) ")
	fmt.Scanln(&operation)

	switch {
	case operation == "+":
		fmt.Println(firstFloat + secondFloat)
	case operation == "-":
		fmt.Println(firstFloat - secondFloat)
	case operation == "==":
		fmt.Println(firstFloat == secondFloat)
	case operation == "/":
		fmt.Println(firstFloat / secondFloat)
	case operation == "*":
		fmt.Println(firstFloat * secondFloat)
	default:
		fmt.Println("Error!")
	}

	fmt.Println("Do you want to restart program? (y/N) ")
	fmt.Scanln(&tempString)

	switch {
	case tempString == "y" || tempString == "Y":
		controlSum = true
	case tempString == "n" || tempString == "N":
		controlSum = false
	}

	if controlSum == true {
		restart()
	}

}
