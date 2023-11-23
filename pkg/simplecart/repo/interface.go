package repo

import (
	"context"
)

type Interface interface {
	CreateUser(ctx context.Context, book User) (*User, error)
	ReadUser(ctx context.Context, request GetUserRequest) (*User, error)

	CreateProduct(ctx context.Context, data Product) (*Product, error)
	GetProduct(ctx context.Context, request GetProductParam) (*Product, error)
	GetProducts(ctx context.Context) ([]Product, error)

	CreateCartItem(ctx context.Context, data Cart) (*Cart, error)
	GetCart(ctx context.Context, request GetCartParam) (*Cart, error)
	GetCarts(ctx context.Context, request GetCartParam) ([]Cart, error)
	UpdateCartItem(ctx context.Context, data Cart, updatedFields []string) error
}
