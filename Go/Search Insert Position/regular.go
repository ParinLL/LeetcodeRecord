package main

// import (
// 	"fmt"
// )

func main() {
	/* Call the function `sayHello` */
	lists := []int{1, 3, 5, 6}
	searchInsert(lists, 7)

}
func searchInsert(nums []int, target int) int {
	for index, value := range nums {
		if value >= target {
			// fmt.Println(index)
			return index
		} else {
			continue
		}
	}
	if nums[len(nums)-1] < target {
		// fmt.Println(len(nums))
		return len(nums)
	}
	// fmt.Println(index)
	return 0
}
