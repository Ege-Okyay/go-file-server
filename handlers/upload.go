package handlers

import (
	"path/filepath"

	"github.com/Ege-Okyay/go-file-server/config"
	"github.com/Ege-Okyay/go-file-server/services"
	"github.com/gofiber/fiber/v2"
)

func UploadFile(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "File upload failed",
		})
	}

	file, err := fileHeader.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to open file",
		})
	}
	defer file.Close()

	extension := filepath.Ext(fileHeader.Filename)
	filename, err := services.SaveFile(file, extension)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save file",
		})
	}

	fileURL := config.UploadURL + filename

	return c.JSON(fiber.Map{
		"url": fileURL,
	})
}
