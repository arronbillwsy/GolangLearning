package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	max := 0.0
	for i:=0;i<len(s);i++{
		j:=i
		check := make([]int, 255)
		for ;j<len(s);j++{
			if check[int(s[j])] == 0{
				check[int(s[j])] = check[int(s[j])] + 1
			}else {
				max = math.Max(max, float64(j-i))
				break
			}
		}
		if j==len(s){
			max = math.Max(max, float64(j-i))
		}
	}
	return int(max)
}

