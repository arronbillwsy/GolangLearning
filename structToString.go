package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var event Event = Event{
		ID:   10,
		Name: "wsyEvent",
		SubEvent: struct {
			SubID int
			IP    string
		}{2, "10.5.117.135"},
	}
	fmt.Println(structToString(event))
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