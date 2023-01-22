package server

import (
	"strconv"

	"github.com/amirfakhrullah/go-bitly/model"
	"github.com/amirfakhrullah/go-bitly/utils"
	"github.com/gofiber/fiber/v2"
)

type CreateLinkPayload struct {
	RedirectUrl string `json:"redirectUrl"`
	ShortenedId string `json:"shortenedId"`
}

func GetAllLinks(ctx *fiber.Ctx) error {
	links, err := model.GetAllLinks()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all links " + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(links)
}

func GetLinkById(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "invalid id " + err.Error(),
		})
	}
	l, err := model.GetLink(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(l)
}

func CreateLink(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")
	var body CreateLinkPayload
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing body " + err.Error(),
		})
	}
	if body.ShortenedId == "" {
		body.ShortenedId = utils.RandomURL(8)
	}
	dbErr := model.CreateLink(model.Link{
		RedirectUrl: body.RedirectUrl,
		ShortenedId: body.ShortenedId,
	})
	if dbErr != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": dbErr.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"shortenedId": body.ShortenedId,
	})
}
