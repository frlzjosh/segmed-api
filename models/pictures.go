package models

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/segmed-apis/base"
)

type Pictures struct {
	gorm.Model
	PictureUrl string `json:"picture_url"`
	IsFlagged  bool   `json:"is_flagged"`
}

func GetPictures(c *fiber.Ctx) {
	log.Println("GetPictures()")
	db := base.DB
	var pictures []Pictures
	db.Find(&pictures)
	c.JSON(pictures)
}

func FlagPicture(c *fiber.Ctx) {
	log.Println("FlagPicture()")
	id := c.Params("id")
	db := base.DB
	var picture Pictures
	db.Find(&picture, id)
	picture.IsFlagged = !picture.IsFlagged
	db.Save(&picture)
	c.JSON(picture)
}

func SaveNewPicture(c *fiber.Ctx) {
	log.Println("SaveNewPicture()")
	params := new(struct {
		PictureUrl string `json:"picture_url"`
	})
	c.BodyParser(&params)
	db := base.DB
	var picture Pictures
	picture.PictureUrl = params.PictureUrl
	picture.IsFlagged = false
	db.Create(&picture)
	c.JSON(picture)
}
