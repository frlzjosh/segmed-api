package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber"
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
	fmt.Println("DB connection successful!")

	base.DB.AutoMigrate(&models.Pictures{})
	fmt.Println("Database Migrated")

}

func initRoutes(app *fiber.App) {
	app.Get("/get/all", models.GetPictures)
	app.Get("/flag/:id", models.FlagPicture)
	app.Post("/save/", models.SaveNewPicture)
}

func main() {
	app := fiber.New()
	initializeDB()
	defer base.DB.Close()
	initRoutes(app)
	app.Listen(":3000")
}
