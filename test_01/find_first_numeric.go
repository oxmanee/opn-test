package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) >= 2 {
		input := os.Args[1]
		fmt.Println(*findFirstNumeric(input))
	} else {
		fmt.Println("Not Found Value")
	}
}

func findFirstNumeric(input string) *string {
	re := regexp.MustCompile("[0-9]")
	strArrNum := re.FindAllString(input, -1)
	if len(strArrNum) > 0 {
		return &strArrNum[0]
	} else {
		return nil
	}
}
