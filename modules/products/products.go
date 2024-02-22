package products

import (
	"github.com/maxexq/parksoi-shop/modules/appinfo"
	"github.com/maxexq/parksoi-shop/modules/entities"
)

type Product struct {
	Id          string            `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Category    *appinfo.Category `json:"category"`
	Price       float64           `json:"price"`
	Images      []*entities.Image `json:"images"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"uploaded_at"`
}

type ProductFilter struct {
	Id     string `query:"id"`
	Search string `query:"search"`
	*entities.PaginationReq
	*entities.SortReq
}
