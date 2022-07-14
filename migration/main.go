package main

import (
	"log"
	"go_pos/model"
)

func main(){
	log.Println("Migration started")

	db := model.Connection()

	db.AutoMigrate(&model.Intt{})

	log.Println("Migration finished")

}