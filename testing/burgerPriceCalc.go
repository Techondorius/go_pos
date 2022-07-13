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
	fmt.Println(string(jsonInput))

	var decodedInput []request_body
	json.Unmarshal([]byte(jsonInput), &decodedInput)
	fmt.Println(decodedInput)
}

func setCheck(slice []request_body){
	for i := 0; i < len(slice); i++ {
		if slice[i].ItemChildren != nil{
			setCheck(slice[i].ItemChildren)
		} else {

		}
	}
}

type item struct {
	itemID int
	itemName string
	price int
	grill []string
	itemChild []item
}

type request_body struct {
	ItemID int `json:"itemID"`
	Grills []string `json:"grills,omitempty"`
	ItemChildren []request_body `json:"itemChildren,omitempty"`
}
