package handler

import (
	internalError "github.com/andhikasamudra/test-pcs-group/internal/error"
	"github.com/andhikasamudra/test-pcs-group/internal/handler"
	"github.com/go-playground/validator/v10"
	"net/http"

	"github.com/andhikasamudra/test-pcs-group/pkg/simplecart/dto"
	"github.com/andhikasamudra/test-pcs-group/pkg/simplecart/service"
	"github.com/gofiber/fiber/v2"
)

type Dependency struct {
	SimpleCartService service.ServiceInterface
	Validator         *validator.Validate
	Resp              handler.ResponseInterface
}

type Handler struct {
	SimpleCartService service.ServiceInterface
	Validator         *validator.Validate
	Resp              handler.ResponseInterface
}

func NewHandler(d Dependency) *Handler {
	resp := handler.NewResponse()
	return &Handler{
		SimpleCartService: d.SimpleCartService,
		Validator:         d.Validator,
		Resp:              resp,
	}
}

func (h *Handler) RegisterUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request dto.CreateUserRequest

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

		err = h.SimpleCartService.CreateUser(c, request)
		if err != nil {
			return c.Status(http.StatusInternalServerError).
				JSON(h.Resp.SetError(err, internalError.InternalServerError, err.Error()))
		}

		return c.Status(fiber.StatusCreated).JSON(h.Resp.SetOk(nil))
	}
}

func (h *Handler) UserLogin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request dto.LoginRequest

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

		result, err := h.SimpleCartService.LoginRequest(c, request)
		if err != nil {
			return c.Status(http.StatusInternalServerError).
				JSON(h.Resp.SetError(err, internalError.InternalServerError, err.Error()))
		}

		return c.Status(fiber.StatusOK).JSON(h.Resp.SetOk(result))
	}
}
