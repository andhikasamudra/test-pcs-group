package dto

import "github.com/google/uuid"

type ProductListItem struct {
	GUID  uuid.UUID `json:"guid"`
	Name  string    `json:"name"`
	Price int64     `json:"price"`
}

type GetProductsResponse struct {
	List []ProductListItem `json:"list"`
}

type CreateProductRequest struct {
	Name  string `json:"name" validate:"required"`
	Price int64  `json:"price" validate:"required"`
}
