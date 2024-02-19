package servers

import (
	"github.com/gofiber/fiber/v2"
	monitorHandlers "github.com/maxexq/parksoi-shop/modules/monitor/handlers"
)

type IModuleFactory interface {
	MonitorModule()
}

type ModuleFactory struct {
	router fiber.Router
	server *server
}

func InitModule(r fiber.Router, server *server) *ModuleFactory {
	return &ModuleFactory{
		router: r,
		server: server,
	}
}

func (m *ModuleFactory) MonitorModule() {
	handler := monitorHandlers.MonitorHandler(m.server.cfg)
	m.router.Get("/", handler.HealthCheck)
}
