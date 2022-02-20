package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) >= 2 {
		input := os.Args[1]
		scoreInt, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Input is invalid")
			return
		}
		if scoreInt > 100 || scoreInt < 0 {
			fmt.Println("Input is invalid")
			return
		}
		fmt.Println("Grade : ", calculateGrade(scoreInt))
	} else {
		fmt.Println("Not found value")
	}
}

func calculateGrade(input int) string {
	if input >= 90 {
		return "A"
	} else if input >= 80 && input < 90 {
		return "B"
	} else if input >= 70 && input < 80 {
		return "C"
	} else if input >= 60 && input < 70 {
		return "D"
	} else if input >= 50 && input < 60 {
		return "E"
	} else {
		return "F"
	}
}
