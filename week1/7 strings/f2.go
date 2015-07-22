package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	miToKm = 1.60934
)

func main() {
	from := os.Args[1]
	to := os.Args[2]

	switch {
	case strings.HasSuffix(from, "mi"):
		switch to {
		case "km":
			miles, _ := strconv.ParseFloat(from[:len(from)-2], 64)
			fmt.Println(miles * miToKm)
		case "me":
			miles, _ := strconv.ParseFloat(from[:len(from)-2], 64)
			fmt.Println(miles * miToKm * 1000)
		case "ft":
			miles, _ := strconv.ParseInt(from[:len(from)-2], 64)

		}

	}
}
