package api

import (
	"github.com/gofiber/fiber/v2"
	"gocv.io/x/gocv"
	"io"
)

func extractImageFrom(c *fiber.Ctx) (*gocv.Mat, error) {
	file, err := c.FormFile("face")
	if err != nil {
		return nil, err
	}

	f, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	mat, err := gocv.IMDecode(b, gocv.IMReadUnchanged)
	if err != nil {
		return nil, err
	}

	return &mat, err
}
