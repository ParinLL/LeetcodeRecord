package main

// import (
// 	"strings"
// )

func main() {
	/* Call the function `sayHello` */
	romanToInt("MCMXCIV")
}

func romanToInt(s string) int {
	var total int

	for indexOfinput, valueOfinput := range s {

		// fmt.Printf("%d%c\n", i, v)
		symbol := string(valueOfinput)
		// fmt.Printf("%c\n", symbolbyte)
		var value int

		switch symbol {
		case "I":
			value = 1
			if indexOfinput+1 < len(s) && (string(s[indexOfinput+1]) == "V" || string(s[indexOfinput+1]) == "X") {
				value = -1
			}
		case "V":
			value = 5
		case "X":
			value = 10
			if indexOfinput+1 < len(s) && (string(s[indexOfinput+1]) == "L" || string(s[indexOfinput+1]) == "C") {
				value = -10
			}
		case "L":
			value = 50
		case "C":
			value = 100
			if indexOfinput+1 < len(s) && (string(s[indexOfinput+1]) == "D" || string(s[indexOfinput+1]) == "M") {
				value = -100
			}
		case "D":
			value = 500
		case "M":
			value = 1000
		default:
			value = 0
		}
		total = total + value
		// fmt.Printf("Value = %d\n", value)

	}
	return total
}
