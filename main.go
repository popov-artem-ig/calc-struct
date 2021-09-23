package main

import (
	"bufio"
	"calc-struct/calculator"
	"fmt"
	"os"
	"strings"

	"calc-struct/converter"
)

func main() {
	
	print()
	str,err := readInputStr()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//validator.Validate()
	rpnArr, err := converter.Convert(str)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("reverse polish notation")
	fmt.Println(rpnArr)

	result, err := calculator.Calculate(rpnArr)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(result)
}

func print() {
	/*
	10*(2+5)-14/(1+2*(1+2))
	(1 + 2) * 4 + 3
	2.2 * 10 - 15
	-5 + 5 * 6
	*/
	fmt.Println("Calculation of an arbitrary notation with output in reverse polish notation")
	fmt.Println("input expression and press Enter")
}

/**
	Read input string
 */
func readInputStr() (string,error) {
	in := bufio.NewReader(os.Stdin)
	scanStr, err := in.ReadString('\n')
	fmt.Println("clear space...")
	scanStr = strings.Replace(scanStr, " ", "", -1)
	fmt.Println("replace ',' to '.'...")
	scanStr = strings.Replace(scanStr, ",", ".", -1)
	if err != nil {
		return "", err
	}

	return scanStr, nil
}

