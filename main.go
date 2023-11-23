package main

import (
	"fmt"
	"github.com/andhikasamudra/test-pcs-group/pkg/simplecart"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/andhikasamudra/test-pcs-group/config"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/uptrace/bun/extra/bundebug"

	"github.com/gofiber/fiber/v2"
)

func main() {
	run()
}

func run() {
	app := fiber.New()

	db := config.GetConnection()
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("services is up"))
	})
	app.Get("/ready", func(c *fiber.Ctx) error {
		return c.SendString("ready")
	})
	api := app.Group("/api")
	simplecart.InitRoute(api, db)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	var serverShutdown sync.WaitGroup

	go func() {
		_ = <-c //nolint:all
		fmt.Println("Gracefully shutting down...")
		serverShutdown.Add(1)
		defer serverShutdown.Done()
		_ = app.ShutdownWithTimeout(60 * time.Second)
	}()

	// ...

	if err := app.Listen(os.Getenv("APP_PORT")); err != nil {
		log.Panic(err)
	}

	serverShutdown.Wait()

	fmt.Println("Running cleanup tasks...")
}
