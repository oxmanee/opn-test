package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) >= 2 {
		input := os.Args[1]
		fmt.Println(checkIp(input))
	} else {
		fmt.Println("Not Found Value")
	}
}

func checkIp(input string) bool {
	parts := strings.Split(input, ".")

	if len(parts) < 4 {
		return false
	}

	for _, x := range parts {
		if i, err := strconv.Atoi(x); err == nil {
			zeroCase := x[:1]
			if zeroCase == "0" && len(x) > 1 {
				return false
			}

			if i < 0 || i > 255 {
				return false
			}
		} else {
			return false
		}

	}
	return true
}
