package appinfoHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxexq/parksoi-shop/config"
	"github.com/maxexq/parksoi-shop/modules/appinfo/appinfoUsecase"
	"github.com/maxexq/parksoi-shop/modules/entities"
	"github.com/maxexq/parksoi-shop/pkg/auth"
)

type IAppinfoHandler interface {
	GenerateApiKey(c *fiber.Ctx) error
}

type generateApiKeyErCode string

const (
	generateApiKeyErr generateApiKeyErCode = "appinfo-001"
)

type appinfoHandler struct {
	cfg            config.IConfig
	appinfoUsecase appinfoUsecase.IAppinfoUsecase
}

func AppinfoHandler(cfg config.IConfig, appinfoUsecase appinfoUsecase.IAppinfoUsecase) IAppinfoHandler {
	return &appinfoHandler{
		cfg:            cfg,
		appinfoUsecase: appinfoUsecase,
	}
}

func (h *appinfoHandler) GenerateApiKey(c *fiber.Ctx) error {
	apiKey, err := auth.NewParksoiAuth(
		auth.ApiKey,
		h.cfg.Jwt(),
		nil,
	)
	if err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrInternalServerError.Code,
			string(generateApiKeyErr),
			err.Error(),
		).Res()
	}

	return entities.NewResponse(c).Success(
		fiber.StatusOK,
		&struct {
			Key string `json:"Key"`
		}{
			Key: apiKey.SignToken(),
		},
	).Res()
}
