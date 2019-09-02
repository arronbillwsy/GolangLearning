package main

import "fmt"

func main() {
	a := []int{1,2,3,4,5}
	fmt.Println(a[3:])
	fmt.Println(a[:3])
}


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	if len1 < len2{
		t := nums1
		nums1 = nums2
		nums2 = t
	}

	if len(nums1) == 0 {
		if len(nums2) % 2 == 0{
			return float64((nums2[len(nums2)/2] + nums2[len(nums2)/2-1]) / 2)
		}
	}
	a := 0
	if len(nums1) < len(nums2){
		a = nums1[len(nums1)/2]
	}else {
		a = nums2[len(nums2)/2]
	}

}

