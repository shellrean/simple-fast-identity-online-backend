package main

import (
	"fido-bio/internal/api"
	"fido-bio/internal/config"
	"fido-bio/internal/connection"
	"fido-bio/internal/repository"
	"fido-bio/internal/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()

	dbConnection := connection.GetDatabase(cnf.Database)

	userRepository := repository.NewUser(dbConnection)
	challengeRepository := repository.NewChallenge(dbConnection)

	userService := service.NewUser(userRepository)
	challengeService := service.NewChallenge(challengeRepository, userRepository)

	app := fiber.New()
	api.NewUser(app, userService)
	api.NewChallenge(app, challengeService)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
