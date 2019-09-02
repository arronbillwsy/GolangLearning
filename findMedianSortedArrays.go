package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{1,3,5}
	b := []int{2,4,6,8,10}
	fmt.Println(findMedianSortedArrays(a,b))
}

func swap(nums1 []int, nums2 []int) ([]int, []int) {
	t := nums1
	nums1 = nums2
	nums2 = t
	return nums1,nums2
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	if len1>len2{
		return findMedianSortedArrays(nums2, nums1)
	}
	low := 0
	high := len1
	for low <= high{
		index1 := (low+high)/2
		left1 := -1
		right1 := -1
		if index1>0{
			left1 = nums1[len1-1]
		}
		if index1<len1{
			right1 = nums1[index1]
		}
		index2 := (len1+len2+1)/2-index1
		left2 := -1
		right2 := -1
		if index2>0{
			left2 = nums2[len2-1]
		}
		if index2<len2{
			right2 = nums2[index2]
		}
		if (left1<=right2) && (left2<=right1){
			if (len1+len2)%2 == 0{
				return (math.Max(float64(left1), float64(left2))+math.Min(float64(right1), float64(right2)))/2
			}else {
				return math.Max(float64(left1), float64(left2))
			}
		}else if left1 > right2{
			high = index1-1
		}else {
			low = index1+1
		}
	}
	return 0

	//if len1 < len2{
	//	swap(nums1,nums2)
	//}
	//if nums2[len1] <= nums1[0]{
	//	swap(nums1,nums2)
	//}
	//if nums1[len1-1] <= nums2[0] {
	//	index := (len1+len2)/2
	//	if (len1 + len2) % 2 == 1 {
	//		if index < len1{
	//			return float64(nums1[index])
	//		}else {
	//			return float64(nums2[index-len1])
	//		}
	//	}else {
	//		if index < len1{
	//			return float64((nums1[index]+nums1[index-1])/2)
	//		}else if index < len1+1{
	//			return float64((nums1[index-1]+nums2[index-len1])/2)
	//		}else {
	//			return float64((nums2[index-len1-1]+nums2[index-len1])/2)
	//		}
	//	}
	//}
	//if len2 == 0 {
	//	if len1 % 2 == 0{
	//		return float64((nums1[len1/2] + nums1[len1/2-1]) / 2)
	//	}else {
	//		return float64(nums1[len1/2])
	//	}
	//}else if len2 == 1 {
	//	nums1 = append(nums1, nums2[0])
	//	sort.Ints(nums1)
	//	if len1 % 2 == 0 {
	//		return float64((nums1[len1/2] + nums1[len1/2-1]) / 2)
	//	}else {
	//		return float64(nums1[len1/2])
	//	}
	//}else {
	//	median1 := nums1[len1/2]
	//	median2 := nums2[len2/2]
	//	if median1 <= median2{
	//		nums1 = nums1[len1/2:]
	//		nums2 = nums2[:len2-len1/2+1]
	//	}else
	//}
	//
	//return 0
}
