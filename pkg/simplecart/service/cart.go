package service

import (
	"database/sql"
	"github.com/andhikasamudra/test-pcs-group/pkg/simplecart/dto"
	"github.com/andhikasamudra/test-pcs-group/pkg/simplecart/middleware"
	"github.com/andhikasamudra/test-pcs-group/pkg/simplecart/repo"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (s *SimpleCartService) AddCartItem(ctx *fiber.Ctx, request dto.AddCartItem) error {
	claims := middleware.GetClaims(ctx)
	user, err := s.SimpleCartModel.ReadUser(ctx.Context(), repo.GetUserRequest{
		GUID: uuid.MustParse(claims["guid"].(string)),
	})
	if err != nil {
		return err
	}

	product, err := s.SimpleCartModel.GetProduct(ctx.Context(), repo.GetProductParam{GUID: request.ProductGUID})
	if err != nil {
		return err
	}

	existingCartItem, err := s.SimpleCartModel.GetCart(ctx.Context(), repo.GetCartParam{
		ProductID: product.ID,
		UserID:    user.ID,
	})
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if existingCartItem != nil {
		existingCartItem.Qty = request.Qty
		err = s.SimpleCartModel.UpdateCartItem(ctx.Context(), *existingCartItem, []string{"qty"})
		if err != nil {
			return err
		}
	} else {
		_, err = s.SimpleCartModel.CreateCartItem(ctx.Context(), repo.Cart{
			UserID:    user.ID,
			ProductID: product.ID,
			Qty:       request.Qty,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *SimpleCartService) GetCarts(ctx *fiber.Ctx) (*dto.GetCartListResponse, error) {
	claims := middleware.GetClaims(ctx)
	user, err := s.SimpleCartModel.ReadUser(ctx.Context(), repo.GetUserRequest{
		GUID: uuid.MustParse(claims["guid"].(string)),
	})
	if err != nil {
		return nil, err
	}

	result, err := s.SimpleCartModel.GetCarts(ctx.Context(), repo.GetCartParam{
		UserID: user.ID,
	})
	if err != nil {
		return nil, err
	}

	var resp dto.GetCartListResponse
	var coupons int
	for _, item := range result {
		resp.List = append(resp.List, dto.CartListItem{
			ProductGUID: item.Product.GUID,
			ProductName: item.Product.Name,
			Qty:         item.Qty,
		})
		if item.Product.Price >= 50000 && item.Product.Price < 100000 {
			coupons += 1
		}
		if item.Product.Price >= 100000 {
			coupons += 2
		}
	}

	resp.Coupons = coupons

	return &resp, nil
}
