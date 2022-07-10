package main

import (
	"fmt"
)

func main() {
	a := burger{100}
	fmt.Println(a)
}

type burger struct {
	price string
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