package main

import (
	"fmt"
	"strings"
)

func main() {
	/* Call the function `sayHello` */
	romanToInt("III")
}

// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.
func romanToInt(s string) int {
	var total int
	var validRoman map[int]string
	validRoman = make(map[int]string)
	var skip bool
	for indexOfinput, valueOfinput := range s {
		if skip {
			skip = false
			continue
		}

		// symbolbyte := s[indexOfinput]
		symbol := string(valueOfinput)
		// fmt.Println(indexOfinput)
		// fmt.Println(symbol)
		// fmt.Printf("%c\n", symbolbyte)

		switch {
		case strings.Contains(symbol, "I"):
			if indexOfinput < len(s)-1 && string(s[indexOfinput+1]) == "V" {
				validRoman[indexOfinput] = "IV"
				skip = true
			} else if indexOfinput < len(s)-1 && string(s[indexOfinput+1]) == "X" {
				validRoman[indexOfinput] = "IX"
				skip = true
			} else {
				validRoman[indexOfinput] = "I"
			}
		case strings.Contains(symbol, "V"):
			validRoman[indexOfinput] = "V"
		case strings.Contains(symbol, "X"):
			if indexOfinput < len(s)-1 && string(s[indexOfinput+1]) == "L" {
				validRoman[indexOfinput] = "XL"
				skip = true
			} else if indexOfinput < len(s)-1 && string(s[indexOfinput+1]) == "C" {
				validRoman[indexOfinput] = "XC"
				skip = true
			} else {
				validRoman[indexOfinput] = "X"
			}
		case strings.Contains(symbol, "L"):
			validRoman[indexOfinput] = "L"

		case strings.Contains(symbol, "C"):
			if indexOfinput < len(s)-1 && string(s[indexOfinput+1]) == "D" {
				validRoman[indexOfinput] = "CD"
				skip = true
			} else if indexOfinput < len(s)-1 && string(s[indexOfinput+1]) == "M" {
				validRoman[indexOfinput] = "CM"
				skip = true
			} else {
				validRoman[indexOfinput] = "C"
			}
		case strings.Contains(symbol, "D"):
			validRoman[indexOfinput] = "D"

		case strings.Contains(symbol, "M"):
			validRoman[indexOfinput] = "M"
		default:
			skip = false
		}
		// fmt.Printf("Value = %d\n", value)
	}
	// fmt.Println(validRoman)
	for _, valueOfmap := range validRoman {
		var value int
		switch {
		case valueOfmap == "I":
			value = 1
		case valueOfmap == "IV":
			value = 4
		case valueOfmap == "V":
			value = 5
		case valueOfmap == "IX":
			value = 9
		case valueOfmap == "X":
			value = 10
		case valueOfmap == "XL":
			value = 40
		case valueOfmap == "L":
			value = 50
		case valueOfmap == "XC":
			value = 90
		case valueOfmap == "C":
			value = 100
		case valueOfmap == "CD":
			value = 400
		case valueOfmap == "D":
			value = 500
		case valueOfmap == "CM":
			value = 900
		case valueOfmap == "M":
			value = 1000
		default:
			value = 0
		}
		total += value
	}
	fmt.Printf("Value = %d\n", total)
	return total
}
