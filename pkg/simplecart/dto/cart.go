package dto

import "github.com/google/uuid"

type AddCartItem struct {
	ProductGUID uuid.UUID `json:"product_guid" validate:"required"`
	Qty         int       `json:"qty" validate:"required"`
}

type CartListItem struct {
	ProductGUID uuid.UUID `json:"product_guid"`
	ProductName string    `json:"product_name"`
	Qty         int       `json:"qty"`
}

type GetCartListResponse struct {
	List    []CartListItem `json:"list"`
	Coupons int            `json:"coupons"`
}
