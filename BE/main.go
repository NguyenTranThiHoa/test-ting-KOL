package main

import (
	"log"
	"wan-api-kol-event/Controllers"
	"wan-api-kol-event/Initializers"
	"wan-api-kol-event/Models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	Initializers.LoadEnvironmentVariables()
	Initializers.ConnectToDB()
	Initializers.DB.AutoMigrate(&Models.Kol{})
}

func main() {
	r := gin.Default()

	// Thêm middleware CORS ở đây 👇
	r.Use(cors.Default())

	// Seed dữ liệu lần đầu (xong rồi thì comment lại)
	Initializers.SeedKols()

	// API route
	r.GET("/kols", Controllers.GetKolsController)

	if err := r.Run(":8081"); err != nil {
		log.Println("Failed to start server")
	}
}
