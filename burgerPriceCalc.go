package main

import (
	"fmt"
)

func main() {
	s := item{
		itemID: 2,
		itemName: "potatoL",
	}

	d := item{
		itemID: 3,
		itemName: "colaL",
	}

	b := item{
		itemID: 1,
		itemName: "DCB",
		price: 390,
		grill: []string{
			"ADD ｹﾁｬｯﾌﾟ",
			"LIGHT ﾏｽﾀｰﾄﾞ",
		},
		itemChild: []*item{
			&s, &d,
		},
	}

	fmt.Println(b)
	for i := 0; i < len(b.itemChild); i ++{
		fmt.Println(*b.itemChild[i])
	}
}

type item struct {
	itemID int
	itemName string
	price int
	grill []string
	itemChild []*item
}

type request_body struct {
	itemID int
	grills []string
	itemChildren []any
}