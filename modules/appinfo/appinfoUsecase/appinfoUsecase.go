package appinfoUsecase

import "github.com/maxexq/parksoi-shop/modules/appinfo/appinfoRepository"

type IAppinfoUsecase interface {
}

type appinfoUsecase struct {
	appinfoRepository appinfoRepository.IAppinfoRepository
}

func AppinfoUsecase(appinfoRepository appinfoRepository.IAppinfoRepository) IAppinfoUsecase {
	return &appinfoUsecase{
		appinfoRepository: appinfoRepository,
	}
}
