package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=fiber port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(20)            
	sqlDB.SetMaxIdleConns(10)            
	sqlDB.SetConnMaxLifetime(time.Hour)  
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	fmt.Println("connect to database")

}