package service

import (
	"github.com/andhikasamudra/test-pcs-group/pkg/simplecart/dto"
	"github.com/andhikasamudra/test-pcs-group/pkg/simplecart/repo"
	"github.com/gofiber/fiber/v2"
)

func (s *SimpleCartService) GetProducts(ctx *fiber.Ctx) (*dto.GetProductsResponse, error) {
	result, err := s.SimpleCartModel.GetProducts(ctx.Context())
	if err != nil {
		return nil, err
	}

	var resp dto.GetProductsResponse
	for _, item := range result {
		resp.List = append(resp.List, dto.ProductListItem{
			GUID:  item.GUID,
			Name:  item.Name,
			Price: item.Price,
		})
	}

	return &resp, nil
}

func (s *SimpleCartService) AddNewProduct(ctx *fiber.Ctx, request dto.CreateProductRequest) error {
	newProduct := repo.Product{
		Name:  request.Name,
		Price: request.Price,
	}
	_, err := s.SimpleCartModel.CreateProduct(ctx.Context(), newProduct)
	if err != nil {
		return err
	}

	return nil
}
