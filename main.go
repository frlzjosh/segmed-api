package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jinzhu/gorm"
	"github.com/segmed-apis/base"
	"github.com/segmed-apis/models"
)

func initializeDB() {
	var err error
	base.DB, err = gorm.Open("mysql", base.DBUrl(base.BuildDBConfig()))
	if err != nil {
		panic(err)
	}
	log.Println("DB connection successful!")

	base.DB.AutoMigrate(&models.Pictures{})
	log.Println("Database Migrated")

}

func initRoutes(app *fiber.App) {
	app.Get("/get/all", models.GetPictures)
	app.Get("/flag/:id", models.FlagPicture)
	app.Post("/save/", models.SaveNewPicture)
	app.Post("/delete", models.DeletePicture)
}

func main() {
	app := fiber.New()
	app.Use(cors.New())
	initializeDB()
	defer base.DB.Close()
	initRoutes(app)
	app.Listen(":3000")
}
