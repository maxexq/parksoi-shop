package middlewaresHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/maxexq/parksoi-shop/config"
	"github.com/maxexq/parksoi-shop/modules/entities"
	"github.com/maxexq/parksoi-shop/modules/middlewares/middlewaresUsecases"
)

type IMiddlewaresHandler interface {
	Cors() fiber.Handler
	RouterCheck() fiber.Handler
}

type middlewareHandlerErrCode string

const (
	routerCheckErr middlewareHandlerErrCode = "middleware-001"
)

type middlewaresHandler struct {
	cfg                config.IConfig
	middlewaresUsecase middlewaresUsecases.IMiddlewaresUsecases
}

func MiddlewaresHandlers(cfg config.IConfig, middlewaresUsecase middlewaresUsecases.IMiddlewaresUsecases) IMiddlewaresHandler {
	return &middlewaresHandler{
		cfg:                cfg,
		middlewaresUsecase: middlewaresUsecase,
	}
}

func (h *middlewaresHandler) Cors() fiber.Handler {
	return cors.New(cors.Config{
		Next:             cors.ConfigDefault.Next,
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST, HEAD, PUT, DELETE, PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	})
}

func (h *middlewaresHandler) RouterCheck() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return entities.NewResponse(c).Error(fiber.ErrNotFound.Code, string(routerCheckErr), "Router not found").Res()
	}
}
