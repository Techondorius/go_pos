package controller

import (
	"go_pos/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Requests []struct {
	ItemID int `json:"itemID"`
	GrillsFree []string `json:"grillsFree,omitempty"`
	GrillsPaid Requests `json:"grillsPaid,omitempty"`
	ItemChildren Requests `json:"itemChildren,omitempty"`
}

type Masters [](model.ItemMaster)

// GET /api/price/testddata
// SQLにTestDataを挿れる
func InsertTestData(c *gin.Context) {
	im := []model.ItemMaster{
		{
			ID: 1001,
			Name: "DCB",
		},{
			ID: 1002,
			Name: "ﾎﾟﾃﾄM",
		},{
			ID: 1003,
			Name: "ｺｰﾗM",
		},{
			ID: 1004,
			Name: "ﾄﾏﾄ",
		},
	}
	if err := model.CreateItemMaster(im); err != nil {
		log.Println(err)
	}
}

// GET /api/price/
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

	r.PriceCalc()
}

// 合計金額を計算するメイン
func (r *Requests) PriceCalc() int{
	s := r.ItemDetail()
	log.Println(s)
	if err := s.CheckStructure(); err != nil {
		log.Println(err)
	}
	return s.CalcPrice()

}

// 各アイテムの詳細を取得(RequestsをMastersに変換)
func (r Requests) ItemDetail() Masters {
	var ret Masters
	for i := 0; i < len(r); i++ {
		id := r[i].ItemID
		ic := r[i].ItemChildren
		if ic != nil {
			ms, err := model.ReadItemMasterByID(id)
			if err != nil{
				log.Println(err)
			}
			ms.Children = ic.ItemDetail()
			ret = append(ret, ms)
		} else {
			k, err := model.ReadItemMasterByID(id)
			if err != nil{
				log.Println(err)
			}
			ret = append(ret, k)
		}
	}
	return ret
}

// ItemChildrenに入るべきでないものが入っていないか確認
func (it *Masters) CheckStructure() error {
	return nil
}

// Mastersの合計金額を計算する！
func (s Masters) CalcPrice() int {
	return 0
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