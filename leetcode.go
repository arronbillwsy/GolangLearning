package main

import (
	"encoding/json"
	"fmt"
	"math"
)

func main()  {
	//num:=0
	//num = 1
	//fmt.Println(num)
	//nums:=[]int{2,7,11,15}
	//target:=9
	//fmt.Println(twoSum(nums, target))
	//var l11 ListNode = ListNode{
	//	Val:  9,
	//	Next: nil,
	//}
	//var l12 ListNode = ListNode{
	//	Val:  8,
	//	Next: nil,
	//}
	//l11.Next = &l12
	//var l21 ListNode = ListNode{
	//	Val:  1,
	//	Next: nil,
	//}
	//res := addTwoNumbers(&l11, &l21)
	//fmt.Println(res)
	//fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	//var event Event = Event{
	//	ID:   10,
	//	Name: "wsyEvent",
	//	SubEvent: struct {
	//		SubID int
	//		IP    string
	//	}{2, "10.5.117.135"},
	//}
	//fmt.Println(structToString(event))
	//mymap := make(map[int]string , 3)
	//for i:=0;i<3;i++{
	//	mymap[i] = "abc"
	//}
	//fmt.Println(mymap)
	a := []int{1,2,3,4,5}
	fmt.Println(a[3:])
	fmt.Println(a[:3])
}

type ListNode struct {
	Val int
	Next *ListNode
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head ListNode = ListNode{
		Val:  0,
		Next: nil,
	}
	cur := &head
	flag := 0
	for ;l1!=nil && l2!=nil;{
		var next ListNode
		next.Val = (l1.Val + l2.Val + flag) % 10
		flag = (l1.Val + l2.Val + flag) / 10
		cur.Next = &next
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1==nil{
		for ;l2!=nil;{
			t := (l2.Val + flag) / 10
			l2.Val = (l2.Val + flag) % 10
			flag = t
			cur.Next = l2
			cur = cur.Next
			l2 = l2.Next
		}
	}
	if l2==nil{
		for ;l1!=nil;{
			t := (l1.Val + flag) / 10
			l1.Val = (l1.Val + flag) % 10
			flag = t
			cur.Next = l1
			cur = cur.Next
			l1 = l1.Next
		}
	}
	if flag==1{
		var tail ListNode = ListNode{
			Val:  1,
			Next: nil,
		}
		cur.Next = &tail
	}
	return head.Next
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

type Event struct {
	ID int
	Name string
	SubEvent struct{
		SubID int
		IP string
	}
}

func structToString(event Event) string {
	res, err := json.Marshal(event)
	if err != nil{
		fmt.Printf("Error: %s", err)
		return ""
	}
	return string(res)
}













