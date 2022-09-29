package main

import "fmt"

func main() {
	lists := []int{1, 2, 3}
	fmt.Println(plusOne(lists))

}

func plusOne(digits []int) []int {

	// looking for last number
	var last = int(digits[len(digits)-1])
	if last != 9 {
		last += 1
		digits[len(digits)-1] = last
		return digits
	}

	newarray := []int{0}
	// 進位
	carry := true
	for i := len(digits) - 2; i >= 0; i-- {
		// 一般進位
		if carry == true && digits[i] != 9 {
			newarray = append(newarray, digits[i]+1)
			carry = false
			continue
		}
		// 傳遞到下個進位
		if carry == true && digits[i] == 9 {
			newarray = append(newarray, 0)
			carry = true
			continue
		}
		// 不需進位
		carry = false
		newarray = append(newarray, digits[i])
	}
	// 當最大一位為9且須進位時值為0，需要在原本的array長度上append 1
	if newarray[len(digits)-1] == 0 {
		newarray = append(newarray, 1)
	}

	// reverse the array
	for i, j := 0, len(newarray)-1; i < j; i, j = i+1, j-1 {
		newarray[i], newarray[j] = newarray[j], newarray[i]
	}

	return newarray
}
