package servers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxexq/parksoi-shop/modules/appinfo/appinfoHandler"
	"github.com/maxexq/parksoi-shop/modules/appinfo/appinfoRepository"
	"github.com/maxexq/parksoi-shop/modules/appinfo/appinfoUsecase"
	"github.com/maxexq/parksoi-shop/modules/files/filesHandlers"
	"github.com/maxexq/parksoi-shop/modules/files/filesUsecases"
	"github.com/maxexq/parksoi-shop/modules/middlewares/middlewaresHandlers"
	middlewaresrepositories "github.com/maxexq/parksoi-shop/modules/middlewares/middlewaresRepositories"
	"github.com/maxexq/parksoi-shop/modules/middlewares/middlewaresUsecases"
	monitorHandlers "github.com/maxexq/parksoi-shop/modules/monitor/handlers"
	"github.com/maxexq/parksoi-shop/modules/products/productsHandlers"
	"github.com/maxexq/parksoi-shop/modules/products/productsRepositories"
	"github.com/maxexq/parksoi-shop/modules/products/productsUsecases"
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
	router.Post("/signup", m.mid.ApiKeyAuth(), handler.SignUpCustomer)
	router.Post("/signin", m.mid.ApiKeyAuth(), handler.SignIn)
	router.Post("/refresh", m.mid.ApiKeyAuth(), handler.RefreshPassport)
	router.Post("/signout", m.mid.ApiKeyAuth(), handler.SignOut)
	router.Post("/signup-admin", handler.SignUpAdmin)
	router.Get("/admin/secret", m.mid.JwtAuth(), m.mid.Authorize(2), handler.GenerateAdminToken)
	router.Get("/:user_id", m.mid.JwtAuth(), m.mid.ParamsCheck(), handler.GetUserProfile)
}

func (m *ModuleFactory) AppinfoModule() {
	repository := appinfoRepository.AppinfoRepository(m.server.db)
	usecase := appinfoUsecase.AppinfoUsecase(repository)
	handler := appinfoHandler.AppinfoHandler(m.server.cfg, usecase)

	router := m.router.Group("/appinfo")

	router.Post("/categories", m.mid.JwtAuth(), m.mid.Authorize(2), handler.AddCategory)
	router.Delete("/:category_id/categories", m.mid.JwtAuth(), m.mid.Authorize(2), handler.RemoveCategory)
	router.Get("/categories", m.mid.ApiKeyAuth(), handler.FindCategory)
	router.Get("/apikey", m.mid.JwtAuth(), m.mid.Authorize(2), handler.GenerateApiKey)
}

func (m *ModuleFactory) FilesModule() {
	usecase := filesUsecases.FileUsecase(m.server.cfg)
	handler := filesHandlers.FileHandler(m.server.cfg, usecase)

	router := m.router.Group("/files")

	router.Post("/upload", m.mid.JwtAuth(), m.mid.Authorize(2), handler.UploadFiles)
	router.Patch("/delete", m.mid.JwtAuth(), m.mid.Authorize(2), handler.DeleteFile)
}

func (m *ModuleFactory) ProductsModule() {
	filesUsecase := filesUsecases.FileUsecase(m.server.cfg)

	repository := productsRepositories.ProductsRepository(m.server.db, m.server.cfg, filesUsecase)
	usecase := productsUsecases.ProductsUsecase(repository)
	handler := productsHandlers.ProductsHandler(m.server.cfg, usecase, filesUsecase)

	router := m.router.Group("/products")

	router.Get("/:product_id", m.mid.ApiKeyAuth(), handler.FindOneProduct)
	router.Get("/", m.mid.ApiKeyAuth(), handler.FindProduct)
	router.Post("/", m.mid.JwtAuth(), m.mid.Authorize(2), handler.AddProduct)
	router.Patch("/:product_id", m.mid.JwtAuth(), m.mid.Authorize(2), handler.UpdateProduct)
	router.Delete("/:product_id", m.mid.JwtAuth(), m.mid.Authorize(2), handler.DeleteProduct)

}
