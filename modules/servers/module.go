package servers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxexq/parksoi-shop/modules/middlewares/middlewaresHandlers"
	middlewaresrepositories "github.com/maxexq/parksoi-shop/modules/middlewares/middlewaresRepositories"
	"github.com/maxexq/parksoi-shop/modules/middlewares/middlewaresUsecases"
	monitorHandlers "github.com/maxexq/parksoi-shop/modules/monitor/handlers"
)

type IModuleFactory interface {
	MonitorModule()
}

type ModuleFactory struct {
	router fiber.Router
	server *server
	mid    middlewaresHandlers.IMiddlewaresHandler
}

func InitModule(r fiber.Router, server *server, mid middlewaresHandlers.IMiddlewaresHandler) *ModuleFactory {
	return &ModuleFactory{
		router: r,
		server: server,
		mid:    mid,
	}
}

func InitMiddleware(s *server) middlewaresHandlers.IMiddlewaresHandler {
	repository := middlewaresrepositories.MiddlewaresRepository(s.db)
	usecase := middlewaresUsecases.MiddlewaresRepository(repository)
	return middlewaresHandlers.MiddlewaresHandlers(s.cfg, usecase)

}

func (m *ModuleFactory) MonitorModule() {
	handler := monitorHandlers.MonitorHandler(m.server.cfg)
	m.router.Get("/", handler.HealthCheck)
}
