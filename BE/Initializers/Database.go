package Initializers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	dsn := "host=localhost user=postgres password=123456 dbname=DemoGoLang port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to PostgreSQL: ", err)
	}

	fmt.Println("✅ Connected to PostgreSQL successfully")
}
