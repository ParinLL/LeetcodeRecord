package main

import "fmt"

func main() {
	/* Call the function `sayHello` */
	lists := []int{1, 3, 5, 6}
	ans := searchInsert(lists, 5)
	fmt.Println("Ans: ", ans)

}
func searchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for end >= start {
		mid := start + (end-start)/2
		// mid := (start + end)/2
		// mid := (start + end)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
		// fmt.Println(start)
	}
	// fmt.Println(start)
	// O(log n)
	return start
}
