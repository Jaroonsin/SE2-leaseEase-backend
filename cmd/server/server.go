package server

import (
	"LeaseEase/config"
	"LeaseEase/internal/handlers"
	"LeaseEase/internal/middleware"
	"context"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	_ "LeaseEase/cmd/docs/v2"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pakornv/scalar-go"
	"go.uber.org/zap"

	swagger "github.com/arsmn/fiber-swagger/v2"
)

type FiberHttpServer struct {
	app      *fiber.App
	cfg      *config.Config
	logger   *zap.Logger
	handlers handlers.Handler
}

func NewFiberHttpServer(cfg *config.Config, logger *zap.Logger, handlers handlers.Handler) *FiberHttpServer {
	return &FiberHttpServer{
		app:      fiber.New(),
		cfg:      cfg,
		logger:   logger,
		handlers: handlers,
	}
}

func (s *FiberHttpServer) initHttpServer(version string) fiber.Router {
	// set global prefix
	router := s.app.Group("/api/" + version)

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

	// swagger with scalar
	filePath := filepath.Join("cmd", "docs", version, "swagger.yaml")
	apiRef, err := scalar.New(filePath, &scalar.Config{
		Theme: scalar.ThemeElysiajs,
	})
	if err != nil {
		panic(err)
	}
	router.Get("/reference", func(c *fiber.Ctx) error {
		htmlContent, err := apiRef.RenderHTML()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.Type("html").SendString(htmlContent)
	})

	router.Get("/swagger/*", swagger.HandlerDefault)

	// healthcheck
	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("server is running !")
	})

	return router
}

func (s *FiberHttpServer) Start() {
	// init http handler

	version := []string{"v1", "v2"}
	for _, v := range version {
		router := s.initHttpServer(v)
		// init modules
		s.initAuthRouter(v, router, s.handlers)
		s.initPropertyRouter(v, router, s.handlers, s.cfg)
		s.initPropertyReviewRouter(v, router, s.handlers, s.cfg)
	}

	// Setup signal capturing for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Run the server in a goroutine so it doesn't block
	go func() {
		if err := s.app.Listen(":" + s.cfg.ServerPort); err != nil {
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

func (s *FiberHttpServer) initAuthRouter(version string, router fiber.Router, httpHandler handlers.Handler) {
	authRouter := router.Group("/auth")

	if version == "v1" {
		authRouter.Post("/register", httpHandler.Auth().Register)
		authRouter.Post("/login", httpHandler.Auth().Login)

	} else if version == "v2" {
		authRouter.Post("/register", httpHandler.Auth().Register)
		authRouter.Post("/login", httpHandler.Auth().Login)
		authRouter.Get("/check", httpHandler.Auth().AuthCheck)
		authRouter.Post("/logout", httpHandler.Auth().Logout)
		authRouter.Post("/request-otp", httpHandler.Auth().RequestOTP)
		authRouter.Post("/verify-otp", httpHandler.Auth().VerifyOTP)
		authRouter.Post("/forgot-password", httpHandler.Auth().ResetPasswordRequest)
		authRouter.Post("/reset-password", httpHandler.Auth().ResetPassword)
	}
}

func (s *FiberHttpServer) initPropertyRouter(version string, router fiber.Router, httpHandler handlers.Handler, cfg *config.Config) {
	propertyRouter := router.Group("/properties", middleware.AuthRequired(cfg))
	reservationRounter := router.Group("/reservations", middleware.AuthRequired(cfg))

	if version == "v1" {
		// property
		propertyRouter.Post("/create", httpHandler.Property().CreateProperty)
		propertyRouter.Put("/update/:id", httpHandler.Property().UpdateProperty)
		propertyRouter.Delete("/delete/:id", httpHandler.Property().DeleteProperty)
		propertyRouter.Get("/", httpHandler.Property().GetAllProperty)
		propertyRouter.Get("/:id", httpHandler.Property().GetPropertyByID)

	} else if version == "v2" {
		// property
		propertyRouter.Post("/create", httpHandler.Property().CreateProperty)
		propertyRouter.Put("/update/:id", httpHandler.Property().UpdateProperty)
		propertyRouter.Delete("/delete/:id", httpHandler.Property().DeleteProperty)
		propertyRouter.Get("/get", httpHandler.Property().GetAllProperty)
		propertyRouter.Get("/get/:id", httpHandler.Property().GetPropertyByID)
		propertyRouter.Get("/search", httpHandler.Property().SearchProperty)
		propertyRouter.Get("/autocomplete", httpHandler.Property().AutoComplete)

		// reservation
		reservationRounter.Post("/create", httpHandler.Reservation().CreateReservation)
		reservationRounter.Put("/update/:id", httpHandler.Reservation().UpdateReservation)
		reservationRounter.Delete("/delete/:id", httpHandler.Reservation().DeleteReservation)
	}

}

func (s *FiberHttpServer) initPropertyReviewRouter(version string, router fiber.Router, httpHandler handlers.Handler, cfg *config.Config) {
	propertyReviewRouter := router.Group("/propertyReview", middleware.AuthRequired(cfg))
	if version == "v2" {
		propertyReviewRouter.Post("/create", httpHandler.Review().CreateReview)
		propertyReviewRouter.Put("/update/:id", httpHandler.Review().UpdateReview)
		propertyReviewRouter.Delete("/delete/:id", httpHandler.Review().DeleteReview)
		propertyReviewRouter.Get("/get/:propertyID", httpHandler.Review().GetAllReviews)
	}
}
