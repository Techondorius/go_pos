package model

import (
	"log"
)

func CreateItemMaster(im []ItemMaster) (error){
	log.Println("CreatingItemMaster")
	db := Connection()
	if err := db.Create(&im); err != nil {
		log.Println(err)
	}
	return nil
}

func ReadItemMaster() ([]ItemMaster, error){
	db := Connection()
	var items []ItemMaster
	db.Find(&items)
	return items, nil
}

func ReadItemMasterByID(id int) (ItemMaster, error){
	db := Connection()
	var items ItemMaster
	db.Where("id=?", id).Find(&items)
	return items, nil
}