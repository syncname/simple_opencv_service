package api

import (
	"github.com/gofiber/fiber/v2"
)

func (a *App) GetFacePos(c *fiber.Ctx) error {

	mat, err := extractImageFrom(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := a.rec.Facebox.GetFaces(mat)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"facebox": res})
}

func (a *App) GetEmotionONNX(c *fiber.Ctx) error {

	mat, err := extractImageFrom(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := a.rec.GetOnnxEmotion(mat)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"results": res})
}

func (a *App) GetCaffeEmotion(c *fiber.Ctx) error {

	mat, err := extractImageFrom(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := a.rec.GetCaffeEmotion(mat)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"results": res})
}

func (a *App) GetAge(c *fiber.Ctx) error {

	mat, err := extractImageFrom(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := a.rec.GetAge(mat)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"results": res})
}

func (a *App) GetGender(c *fiber.Ctx) error {

	mat, err := extractImageFrom(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := a.rec.GetGender(mat)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"results": res})
}

func (a *App) GetFullInfo(c *fiber.Ctx) error {

	mat, err := extractImageFrom(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := a.rec.GetFullIno(mat)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"results": res})
}
