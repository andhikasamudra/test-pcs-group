package simplecart

import (
	"github.com/andhikasamudra/test-pcs-group/pkg/simplecart/handler"
	"github.com/andhikasamudra/test-pcs-group/pkg/simplecart/middleware"
	"github.com/andhikasamudra/test-pcs-group/pkg/simplecart/repo"
	"github.com/andhikasamudra/test-pcs-group/pkg/simplecart/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

func InitRoute(r fiber.Router, db *bun.DB) {
	v := validator.New()
	m := repo.NewModel(db)
	s := service.NewService(service.Dependency{
		SimpleCartModel: m,
	})
	h := handler.NewHandler(handler.Dependency{
		SimpleCartService: s,
		Validator:         v,
	})

	api := r.Group("/simplecart")
	api.Post("/register", h.RegisterUser())
	api.Post("/login", h.UserLogin())

	api.Get("/products", middleware.Protected(), h.GetProducts())
	api.Post("/product", middleware.Protected(), h.AddNewProduct())

	api.Get("/carts", middleware.Protected(), h.GetCarts())
	api.Post("/cart", middleware.Protected(), h.AddNewCartItem())
}
