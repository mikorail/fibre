package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mikorail/fibre/database"
	"github.com/mikorail/fibre/model"
	"gorm.io/gorm"
)

// Create a user
func CreateReqs(c *fiber.Ctx) error {
	db := database.DB.Db
	reqs := new(model.RequestAutomation)
	// Store the body in the user and return error if encountered
	err := c.BodyParser(reqs)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	now := time.Now()
	reqs.CreatedDate = &now

	err = db.Create(&reqs).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}
	// Return the created user
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Request has created", "data": reqs})
}

// Get All Users from db
func GetAllReqs(c *fiber.Ctx) error {
	db := database.DB.Db
	var reqs model.RequestAutomation
	// find all users in the database
	err := db.Order("created_date desc").First(&reqs).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{"data": nil})
		}
		return c.Status(500).JSON(fiber.Map{"data": err})
	}

	// Return the most recently created request
	return c.Status(200).JSON(fiber.Map{"data": reqs})
}

// GetSingleUser from db
func GetSingleReq(c *fiber.Ctx) error {
	db := database.DB.Db
	// get id params
	id := c.Params("id")
	var reqs model.RequestAutomation
	// find single user in the database by id
	db.Find(&reqs, "id = ?", id)
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User Found", "data": reqs})
}

func DeleteReq(c *fiber.Ctx) error {
	db := database.DB.Db
	var reqs model.RequestAutomation
	// get id params
	id := c.Params("id")
	// find single user in the database by id
	db.Find(&reqs, "id = ?", id)
	err := db.Delete(&reqs, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete user", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User deleted"})
}
