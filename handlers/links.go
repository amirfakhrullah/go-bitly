package handlers

import (
	"strconv"

	"github.com/amirfakhrullah/go-bitly/helpers"
	"github.com/amirfakhrullah/go-bitly/model"
	"github.com/amirfakhrullah/go-bitly/services"
	"github.com/amirfakhrullah/go-bitly/utils"
	"github.com/gofiber/fiber/v2"
)

type LinkPayload struct {
	RedirectUrl string `json:"redirect_url" validate:"required,min=4,max=100"`
	ShortenedId string `json:"shortened_id" validate:"max=20"`
}

func GetAllLinks(ctx *fiber.Ctx) error {
	userId, err := helpers.GetUserId(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	links, err := services.GetAllLinks(userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all links " + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(links)
}

func GetLinkById(ctx *fiber.Ctx) error {
	userId, err := helpers.GetUserId(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "invalid id " + err.Error(),
		})
	}
	l, err := services.GetLink(id, userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(l)
}

func CreateLink(ctx *fiber.Ctx) error {
	userId, err := helpers.GetUserId(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	ctx.Accepts("application/json")
	var body LinkPayload
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing body " + err.Error(),
		})
	}

	if err = helpers.ValidatePayload(body); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if body.ShortenedId == "" {
		body.ShortenedId = utils.RandomURL(8)
	}
	err = services.CreateLink(model.Link{
		RedirectUrl: body.RedirectUrl,
		ShortenedId: body.ShortenedId,
		UserID:      userId,
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
	userId, err := helpers.GetUserId(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	ctx.Accepts("application/json")
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "invalid id " + err.Error(),
		})
	}
	var l LinkPayload
	if err = ctx.BodyParser(&l); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing body " + err.Error(),
		})
	}

	if err = helpers.ValidatePayload(l); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if l.ShortenedId == "" {
		l.ShortenedId = utils.RandomURL(8)
	}
	if err = services.UpdateLink(userId, model.Link{
		RedirectUrl: l.RedirectUrl,
		ShortenedId: l.ShortenedId,
	}, id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(l)
}

func DeleteLink(ctx *fiber.Ctx) error {
	userId, err := helpers.GetUserId(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "invalid id " + err.Error(),
		})
	}
	if err = services.DeleteLink(userId, id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.SendStatus(fiber.StatusOK)
}

func Redirect(ctx *fiber.Ctx) error {
	shortenedId := ctx.Params("shortenedId")
	if shortenedId == "" {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "invalid id",
		})
	}
	link, err := services.OpenShortenedId(shortenedId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Redirect(link.RedirectUrl, fiber.StatusTemporaryRedirect)
}
