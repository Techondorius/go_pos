package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	r := []request_body{
		{
			ItemID: 1001,
			ItemChildren: []request_body{
				{ItemID: 1002},
				{ItemID: 1003},
			},
		},
		{
			ItemID: 1001,
		},
	}

	jsonInput , _ := json.Marshal(&r)
	fmt.Println("requestBody(Converted json to string): " + string(jsonInput))

	var decodedInput []request_body
	json.Unmarshal([]byte(jsonInput), &decodedInput)
	fmt.Print("requestbody converted to Struct: ")
	fmt.Println(decodedInput)
}

func setCheck(r []request_body){
	for i := 0; i < len(r); i++ {
		if r[i].ItemChildren != nil{
			setCheck(r[i].ItemChildren)
		} else {

		}
	}
}

type item struct {
	itemID int
	itemName string
	price int
	grill_free []string
	grill_paid []item
	itemChild []item
}

type request_body struct {
	ItemID int `json:"itemID"`
	Grills []string `json:"grills,omitempty"`
	ItemChildren []request_body `json:"itemChildren,omitempty"`
}
