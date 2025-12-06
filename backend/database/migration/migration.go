package migration

import (
	"fiber/database"
	"fiber/model/entity"
	"fmt"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{},  &entity.Wallet{}, &entity.Transaction{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("success migate db")
}