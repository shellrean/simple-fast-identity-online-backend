package api

import (
	"context"
	"fido-bio/domain"
	"fido-bio/dto"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"time"
)

type challengeApi struct {
	challengeService domain.ChallengeService
}

func NewChallenge(app *fiber.App, challengeService domain.ChallengeService) {
	ca := challengeApi{
		challengeService: challengeService,
	}

	app.Get("challenge/generate", ca.generate)
	app.Post("challenge/validate", ca.validate)
}

func (ca challengeApi) generate(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	data, err := ca.challengeService.Generate(c)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.Response[string]{Message: err.Error()})
	}
	return ctx.JSON(dto.Response[dto.ChallengeData]{Data: data})
}

func (ca challengeApi) validate(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.ChallengeValidate
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	token, err := ca.challengeService.Validate(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.Response[string]{Message: err.Error()})
	}
	return ctx.JSON(dto.Response[dto.UserData]{Data: token})
}
