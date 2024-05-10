package api

import (
	"context"
	"fido-bio/domain"
	"fido-bio/dto"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"time"
)

type userApi struct {
	userService domain.UserService
}

func NewUser(app *fiber.App, userService domain.UserService) {
	da := userApi{
		userService: userService,
	}

	app.Post("users/register", da.register)
}

func (d userApi) register(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.RegisterUser
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	err := d.userService.Register(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.Response[string]{Message: err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(dto.Response[string]{Message: "success"})
}
