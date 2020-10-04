package models

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/segmed-apis/base"
)

type Pictures struct {
	gorm.Model
	PictureUrl string `json:"picture_url"`
	IsFlagged  bool   `json:"is_flagged"`
	CustomTime string `json:"custom_time"`
	Metadata string `json:"metadata"`
}

func GetPictures(c *fiber.Ctx) error {
	log.Println("GetPictures()")
	db := base.DB
	var pictures []Pictures
	db.Find(&pictures)
	return c.Status(fiber.StatusOK).JSON(pictures)
}

func FlagPicture(c *fiber.Ctx) error {
	log.Println("FlagPicture()")
	id := c.Params("id")
	db := base.DB
	var picture Pictures
	db.Find(&picture, id)
	picture.IsFlagged = !picture.IsFlagged
	firstPicture := getFirstUser()
	if picture.IsFlagged == false {
		picture.CustomTime = firstPicture.CreatedAt.String()
	} else {
		picture.CustomTime = picture.UpdatedAt.String()
	}
	db.Save(&picture)
	return c.Status(fiber.StatusOK).JSON(picture)
}

func getFirstUser() Pictures {
	db := base.DB
	var firstPicture Pictures
	db.First(&firstPicture)
	return firstPicture
}

func SaveNewPicture(c *fiber.Ctx) error {
	log.Println("SaveNewPicture()")
	params := new(struct {
		PictureUrl string `json:"picture_url"`
		Metadata string `json:"metadata"`
	})
	c.BodyParser(&params)
	db := base.DB
	var picture Pictures
	firstPicture := getFirstUser()
	picture.PictureUrl = params.PictureUrl
	picture.Metadata = params.Metadata
	picture.IsFlagged = false
	picture.CustomTime = firstPicture.CreatedAt.String()
	db.Create(&picture)
	return c.Status(fiber.StatusOK).JSON(picture)
}

func DeletePicture(c *fiber.Ctx) error {
	log.Println("DeletePicture()")
	params := new(struct {
		ID int `json:"ID"`
	})
	c.BodyParser(&params)
	db := base.DB
	pictures := db.Delete(&Pictures{}, params.ID)
	return c.JSON(pictures)
}
