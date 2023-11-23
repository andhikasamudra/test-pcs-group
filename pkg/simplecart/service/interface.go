package service

import (
	"github.com/andhikasamudra/test-pcs-group/pkg/simplecart/dto"
	"github.com/gofiber/fiber/v2"
)

type ServiceInterface interface {
	CreateUser(ctx *fiber.Ctx, request dto.CreateUserRequest) error
	LoginRequest(ctx *fiber.Ctx, request dto.LoginRequest) (*dto.LoginResponse, error)

	GetProducts(ctx *fiber.Ctx) (*dto.GetProductsResponse, error)
	AddNewProduct(ctx *fiber.Ctx, request dto.CreateProductRequest) error

	AddCartItem(ctx *fiber.Ctx, request dto.AddCartItem) error
	GetCarts(ctx *fiber.Ctx) (*dto.GetCartListResponse, error)
}
