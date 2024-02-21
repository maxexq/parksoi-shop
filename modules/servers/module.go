package servers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxexq/parksoi-shop/modules/appinfo/appinfoHandler"
	"github.com/maxexq/parksoi-shop/modules/appinfo/appinfoRepository"
	"github.com/maxexq/parksoi-shop/modules/appinfo/appinfoUsecase"
	"github.com/maxexq/parksoi-shop/modules/middlewares/middlewaresHandlers"
	middlewaresrepositories "github.com/maxexq/parksoi-shop/modules/middlewares/middlewaresRepositories"
	"github.com/maxexq/parksoi-shop/modules/middlewares/middlewaresUsecases"
	monitorHandlers "github.com/maxexq/parksoi-shop/modules/monitor/handlers"
	"github.com/maxexq/parksoi-shop/modules/users/usersHandlers"
	"github.com/maxexq/parksoi-shop/modules/users/usersRepositories"
	"github.com/maxexq/parksoi-shop/modules/users/usersUsecases"
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
	usecase := middlewaresUsecases.MiddlewaresUsecase(repository)
	return middlewaresHandlers.MiddlewaresHandler(s.cfg, usecase)

}

func (m *ModuleFactory) MonitorModule() {
	handler := monitorHandlers.MonitorHandler(m.server.cfg)
	m.router.Get("/", handler.HealthCheck)
}

func (m *ModuleFactory) UsersModule() {
	repository := usersRepositories.UsersRepository(m.server.db)
	usecase := usersUsecases.UsersUsecase(m.server.cfg, repository)
	handler := usersHandlers.UsersHandler(m.server.cfg, usecase)

	router := m.router.Group("/users")
	router.Post("/signup", handler.SignUpCustomer)
	router.Post("/signin", handler.SignIn)
	router.Post("/refresh", handler.RefreshPassport)
	router.Post("/signout", handler.SignOut)
	router.Post("/signup-admin", handler.SignUpAdmin)
	router.Get("/admin/secret", m.mid.JwtAuth(), m.mid.Authorize(2), handler.GenerateAdminToken)
	router.Get("/:user_id", m.mid.JwtAuth(), m.mid.ParamsCheck(), handler.GetUserProfile)
}

func (m *ModuleFactory) AppinfoModule() {
	repository := appinfoRepository.AppinfoRepository(m.server.db)
	usecase := appinfoUsecase.AppinfoUsecase(repository)
	handler := appinfoHandler.AppinfoHandler(m.server.cfg, usecase)

	router := m.router.Group("/appinfo")

	_ = router
	_ = handler

}
