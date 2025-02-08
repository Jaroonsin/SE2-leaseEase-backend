package server

import (
	"LeaseEase/config"
	"LeaseEase/internal/handlers"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"

	swagger "github.com/arsmn/fiber-swagger/v2"
)

type FiberHttpServer struct {
	app      *fiber.App
	cfg      *config.DBConfig
	logger   *zap.Logger
	handlers handlers.Handler
}

func NewFiberHttpServer(cfg *config.DBConfig, logger *zap.Logger, handlers handlers.Handler) *FiberHttpServer {
	return &FiberHttpServer{
		app:      fiber.New(),
		cfg:      cfg,
		logger:   logger,
		handlers: handlers,
	}
}

func (s *FiberHttpServer) initHttpServer() fiber.Router {
	// set global prefix
	router := s.app.Group("/api/v1")

	// enable cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Origin,X-PINGOTHER,Accept,Authorization,Content-Type,X-CSRF-Token",
		ExposeHeaders:    "Link",
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// init logger
	// router.Use(logger.New(logger.Config{
	// 	Format:     "${time} ${status} - ${method} ${path}\n",
	// 	TimeFormat: "2006/01/02 15:04:05",
	// 	TimeZone:   "Asia/Bangkok",
	// }))

	// swagger
	router.Get("/swagger/*", swagger.HandlerDefault)

	// healthcheck
	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("server is running !")
	})

	return router
}

func (s *FiberHttpServer) Start() {
	// init http handler
	router := s.initHttpServer()

	// init modules
	s.initAuthRouter(router, s.handlers)
	s.initPropertyRouter(router, s.handlers)

	// Setup signal capturing for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Run the server in a goroutine so it doesn't block
	go func() {
		if err := s.app.Listen(":3000"); err != nil {
			s.logger.Sugar().Fatalf("Error while starting server: %v", err)
		}
	}()

	// Wait for a termination signal
	<-quit
	s.logger.Sugar().Info("Gracefully shutting down server...")

	// Create a deadline for shutdown
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shut down the server
	if err := s.app.Shutdown(); err != nil {
		s.logger.Sugar().Fatalf("Error during server shutdown: %v", err)
	}

	s.logger.Sugar().Info("Server shutdown complete.")
}

func (s *FiberHttpServer) initAuthRouter(router fiber.Router, httpHandler handlers.Handler) {
	authRouter := router.Group("/auth")

	authRouter.Post("/register", httpHandler.Auth().Register)
}

func (s *FiberHttpServer) initPropertyRouter(router fiber.Router, httpHandler handlers.Handler) {
	propertyRouter := router.Group("/properties")

	propertyRouter.Post("/create", httpHandler.Property().CreateProperty)
	propertyRouter.Put("/update/:id", httpHandler.Property().UpdateProperty)
	propertyRouter.Delete("/delete/:id", httpHandler.Property().DeleteProperty)
}
