package controller

import (
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
)

type Requests []struct {
	ItemID int `json:"itemID"`
	GrillsFree []string `json:"grillsFree,omitempty"`
	GrillsPaid Requests `json:"grillsPaid,omitempty"`
	ItemChildren Requests `json:"itemChildren,omitempty"`
}

func GetPrice(c *gin.Context) {
	var r Requests
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		log.Println(err)
		return
	} else {
		res := r.CheckQW(0)
		c.JSON(200, res)
		log.Println(r)
		log.Println(res)
	}

}

// 合計金額を計算するメイン
func (r Requests) PriceCalc() int{
	s := r.ItemDetail()
	if err := s.CheckStructure(); err != nil {
		log.Println(err)
	}
	return s.CalcPrice()

}

// Itemsの合計金額を計算する！
func (s Items) CalcPrice() int {
	return 0
}

// 各アイテムの詳細を取得(RequestsをItemsに変換)
func (r Requests) ItemDetail() Items {
	res := Items{}
	return res
}

// ItemChildrenに入るべきでないものが入っていないか確認
func (it Items) CheckStructure() error {
	return nil
}

// Requestsの最終到達点
type Items []struct{
	ItemID int
	ItemName string
	GrillsPaid []Items
	GrillsFree []string
	ItemChildren Items
}

// Requestのなかのアイテム数をカウント
func (r Requests) CheckQW(num int) int {
	for i := 0; i < len(r); i++ {
		ic := r[i].ItemChildren
		if ic != nil {
			num = ic.CheckQW(num)
			num += 1
		} else {
			num += 1
		}
	}
	return num
}