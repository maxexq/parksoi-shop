package productsHandlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/maxexq/parksoi-shop/config"
	"github.com/maxexq/parksoi-shop/modules/entities"
	"github.com/maxexq/parksoi-shop/modules/files/filesUsecases"
	"github.com/maxexq/parksoi-shop/modules/products"
	"github.com/maxexq/parksoi-shop/modules/products/productsUsecases"
)

type IProductsHandler interface {
	FindOneProduct(c *fiber.Ctx) error
	FindProduct(c *fiber.Ctx) error
}

type productsHandler struct {
	cfg             config.IConfig
	productsUsecase productsUsecases.IProductsUsecase
	filesUsecase    filesUsecases.IFilesUsecase
}

type productsHandlerErrCode string

const (
	findOneProductErr productsHandlerErrCode = "products-001"
	findProductErr    productsHandlerErrCode = "products-002"
)

func ProductsHandler(cfg config.IConfig, productsUsecase productsUsecases.IProductsUsecase, filesUsecase filesUsecases.IFilesUsecase) IProductsHandler {
	return &productsHandler{
		cfg:             cfg,
		productsUsecase: productsUsecase,
		filesUsecase:    filesUsecase,
	}
}

func (h *productsHandler) FindOneProduct(c *fiber.Ctx) error {

	productId := strings.Trim(c.Params("product_id"), " ")

	product, err := h.productsUsecase.FindOneProduct(productId)
	if err != nil {
		return entities.NewResponse(c).Error(fiber.ErrInternalServerError.Code, string(findOneProductErr), err.Error()).Res()
	}

	return entities.NewResponse(c).Success(fiber.StatusOK, product).Res()
}

func (h *productsHandler) FindProduct(c *fiber.Ctx) error {
	req := &products.ProductFilter{
		PaginationReq: &entities.PaginationReq{},
		SortReq:       &entities.SortReq{},
	}

	if err := c.QueryParser(req); err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(findProductErr),
			err.Error(),
		).Res()
	}

	if req.Page < 1 {
		req.Page = 1
	}
	if req.Limit < 5 {
		req.Limit = 5
	}

	if req.OrderBy == "" {
		req.OrderBy = "title"
	}
	if req.Sort == "" {
		req.Sort = "ASC"
	}

	products := h.productsUsecase.FindProduct(req)
	return entities.NewResponse(c).Success(fiber.StatusOK, products).Res()
}
