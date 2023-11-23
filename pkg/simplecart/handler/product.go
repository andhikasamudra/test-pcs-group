package handler

import (
	internalError "github.com/andhikasamudra/test-pcs-group/internal/error"
	"github.com/andhikasamudra/test-pcs-group/pkg/simplecart/dto"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (h *Handler) GetProducts() fiber.Handler {
	return func(c *fiber.Ctx) error {
		resp, err := h.SimpleCartService.GetProducts(c)
		if err != nil {
			return c.Status(http.StatusInternalServerError).
				JSON(h.Resp.SetError(err, internalError.InternalServerError, err.Error()))
		}

		return c.Status(fiber.StatusOK).JSON(h.Resp.SetOk(resp))
	}
}

func (h *Handler) AddNewProduct() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request dto.CreateProductRequest

		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(h.Resp.SetError(err, internalError.RequestError, err.Error()))
		}

		err = h.Validator.Struct(request)
		if err != nil {
			return c.Status(http.StatusBadRequest).
				JSON(h.Resp.SetError(err, internalError.RequestError, err.Error()))
		}

		err = h.SimpleCartService.AddNewProduct(c, request)
		if err != nil {
			return c.Status(http.StatusInternalServerError).
				JSON(h.Resp.SetError(err, internalError.InternalServerError, err.Error()))
		}

		return c.Status(fiber.StatusCreated).JSON(h.Resp.SetOk(nil))
	}
}
