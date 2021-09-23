package main

import (
	"bufio"
	"calc-struct/calculator"
	"fmt"
	"os"

	"calc-struct/converter"
)

func main() {
	
	print()
	var str string
	str,err := readInputStr()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//validator.Validate()

	rpnStr, err := converter.Convert(str)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("reverse polish notation")
	fmt.Println(rpnStr)

	result, err :=calculator.Calculate(rpnStr)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(result)
}

func print() {
	fmt.Println("Calculation of an arbitrary notation with output in reverse polish notation")
	fmt.Println("input expression and press Enter")
}

func readInputStr() (string,error) {
	var scan_str string

	in := bufio.NewReader(os.Stdin)
	scan_str, err := in.ReadString('\n')

	fmt.Println(scan_str)

	if err == nil {
		return "", err
	}

	return scan_str, nil
}

