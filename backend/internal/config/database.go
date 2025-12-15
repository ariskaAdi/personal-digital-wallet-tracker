package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	dialect := "host=localhost user=postgres password=postgres dbname=fiber port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dialect), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	
	return db
}