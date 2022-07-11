package main

import (
	"fmt"
)

func main() {
	s := item{
		itemID: 2,
		itemName: "potatoL",
	}
	b := item{
		itemID: 1,
		itemName: "DCB",
		grill: []string{
			"ADD ｹﾁｬｯﾌﾟ",
			"LIGHT ﾏｽﾀｰﾄﾞ",
		},
		itemChild: []*item{
			&s,
		},
	}

	fmt.Println(b)
	for i := 0; i < len(b.itemChild); i ++{
		fmt.Println(*b.itemChild[i])
	}

	fmt.Println(s)
}

type item struct {
	itemID int
	itemName string
	grill []string
	itemChild []*item
}
/*

{
	responceDatas: blablabla
	items: [
		{
			name: "ﾀﾞﾌﾞﾙﾁｰｽﾞﾊﾞｰｶﾞｰ"
			Price: 390
			Grills: [
				"ADD ｹﾁｬｯﾌﾟ",
				"LIGHT ﾏｽﾀｰﾄﾞ"
			]
			items: [
				{
					Name: "ﾎﾟﾃﾄL"
					Grills:[
						"LIGHT ｼｵ"
					]
				},
				{
					Name: "ｺｰﾗL"
					Grills: [
						"NO ｺｵﾘ"
					]
				}
			]
		}
	]
}
これを落とし込む

 */