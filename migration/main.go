package main

import (
	"log"
	"go_pos/model"
)

func main(){
	log.Println("Migration started")

	db := model.Connection()

	db.AutoMigrate(&model.GrillPossible{})
	db.AutoMigrate(&model.ItemMaster{})
	db.AutoMigrate(&model.GrillMaster{})

	log.Println("Migration finished")

}