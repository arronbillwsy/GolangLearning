package main

import "fmt"

func main() {
	nums := []int{2,7,11,15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	res:= make([]int, 2)
	set:= make(map[int]int, len(nums))
	for i:=0;i<len(nums);i++{
		value, ok:=set[target-nums[i]]
		if ok{
			res[0] = value
			res[1] = i
			return res
		}else {
			set[nums[i]] = i
		}
	}
	return res
}