package main

import (
	"fmt"
)

func main()  {
	mymap := make(map[int]string , 3)
	for i:=0;i<3;i++{
		mymap[i] = "abc"
	}
	fmt.Println(mymap)

}