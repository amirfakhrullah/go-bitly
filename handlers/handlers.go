package handlers

import (
	"fmt"
	"strconv"

	"github.com/amirfakhrullah/go-bitly/model"
	"github.com/amirfakhrullah/go-bitly/services"
	"github.com/amirfakhrullah/go-bitly/utils"
	"github.com/gofiber/fiber/v2"
)

type CreateLinkPayload struct {
	RedirectUrl string `json:"redirectUrl"`
	ShortenedId string `json:"shortenedId"`
}

func GetAllLinks(ctx *fiber.Ctx) error {
	links, err := services.GetAllLinks()
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
	l, err := services.GetLink(id)
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
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing body " + err.Error(),
		})
	}
	if body.ShortenedId == "" {
		body.ShortenedId = utils.RandomURL(8)
	}
	err := services.CreateLink(model.Link{
		RedirectUrl: body.RedirectUrl,
		ShortenedId: body.ShortenedId,
	})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"shortenedId": body.ShortenedId,
	})
}

func UpdateLink(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "invalid id " + err.Error(),
		})
	}
	var l model.Link
	if err = ctx.BodyParser(&l); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing body " + err.Error(),
		})
	}
	if err = services.UpdateLink(l, id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(l)
}

func DeleteLink(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "invalid id " + err.Error(),
		})
	}
	if err = services.DeleteLink(id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.SendStatus(fiber.StatusOK)
}

func Redirect(ctx *fiber.Ctx) error {
	shortenedId := ctx.Params("shortenedId")
	link, err := services.FindByShortenedId(shortenedId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	link.Clicked += 1
	if err = services.UpdateLink(link, link.ID); err != nil {
		fmt.Printf("error updating link")
	}
	return ctx.Redirect(link.RedirectUrl, fiber.StatusTemporaryRedirect)
}
