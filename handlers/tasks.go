package handlers

import (
	"restgo/db"
	"restgo/models"

	"github.com/gofiber/fiber/v2"
)

func Tasks(c *fiber.Ctx) error {
	var tasks []models.Tasks
	err := db.DB.Find(&tasks)
	if err.Error != nil {
		return err.Error
	}
	return c.JSON(&tasks)
}

func Task(c *fiber.Ctx) error {
	id := c.Params("id")
	var task models.Tasks
	if err := db.DB.First(&task, id); err.Error != nil {
		return err.Error
	}
	return c.JSON(&task)
}

func New(c *fiber.Ctx) error {
	var task models.Tasks
	if err := c.BodyParser(&task); err != nil {
		return err
	}
	created := db.DB.Create(&task)
	if created.Error != nil {
		return created.Error
	}
	return c.JSON(fiber.Map{
		"Status": "Registro exitoso",
	})
}

func Done(c *fiber.Ctx) error {
	id := c.Params("id")
	var task models.Tasks
	if err := db.DB.First(&task, id); err.Error != nil {
		return err.Error
	}
	task.Status = !task.Status
	db.DB.Save(&task)
	return c.JSON(fiber.Map{
		"Status": "Se actualizo correctamente",
	})
}

func Edit(c *fiber.Ctx) error {
	id := c.Params("id")
	var task models.Tasks
	if err := db.DB.First(&task, id); err.Error != nil {
		return err.Error
	}
	if err := c.BodyParser(&task); err != nil {
		return err
	}
	db.DB.Save(&task)
	return c.JSON(fiber.Map{
		"Status": "Se actualizo correctamente",
	})
}

func Remove(c *fiber.Ctx) error {
	id := c.Params("id")
	var task models.Tasks
	if err := db.DB.First(&task, id); err.Error != nil {
		return err.Error
	}
	db.DB.Delete(&task)
	return c.JSON(fiber.Map{
		"Status": "Se elimino correctamente",
	})
}
