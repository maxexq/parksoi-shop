package appinfoHandler

import (
	"github.com/maxexq/parksoi-shop/config"
	"github.com/maxexq/parksoi-shop/modules/appinfo/appinfoUsecase"
)

type IAppinfoHandler interface {
}

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
