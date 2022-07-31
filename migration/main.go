package main

import (
	"log"
	"go_pos/model"
)

func main(){
	log.Println("Migration started")

	db := model.ConnectionByTCP()

	db.AutoMigrate(&model.ItemMaster{})
	db.AutoMigrate(&model.GrillMaster{})
	db.AutoMigrate(&model.GrillPossible{})

	log.Println("Migration finished")

}