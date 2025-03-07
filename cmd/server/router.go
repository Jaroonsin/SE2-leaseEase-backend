package server

import (
	"LeaseEase/config"
	"LeaseEase/internal/handlers"
	"LeaseEase/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func (s *FiberHttpServer) initRouter(router fiber.Router) {
	initAuthRouter(router, s.handlers)
	initPropertyRouter(router, s.handlers, s.cfg)
	initLesseeRouter(router, s.handlers, s.cfg)
	initLessorRouter(router, s.handlers, s.cfg)
	initPropertyReviewRouter(router, s.handlers, s.cfg)
	initPaymentRouter(router, s.handlers, s.cfg)
	initUserRouter(router, s.handlers, s.cfg)
}

func initAuthRouter(router fiber.Router, httpHandler handlers.Handler) {
	authRouter := router.Group("/auth")
	authRouter.Post("/register", httpHandler.Auth().Register)
	authRouter.Post("/login", httpHandler.Auth().Login)
	authRouter.Post("/logout", httpHandler.Auth().Logout)
	authRouter.Post("/request-otp", httpHandler.Auth().RequestOTP)
	authRouter.Post("/verify-otp", httpHandler.Auth().VerifyOTP)
	authRouter.Post("/forgot-password", httpHandler.Auth().ResetPasswordRequest)
	authRouter.Post("/reset-password", httpHandler.Auth().ResetPassword)
}

func initPropertyRouter(router fiber.Router, httpHandler handlers.Handler, cfg *config.Config) {
	propertyRouter := router.Group("/properties", middleware.AuthRequired(cfg))
	propertyRouter.Post("/create", httpHandler.Property().CreateProperty)
	propertyRouter.Put("/update/:id", httpHandler.Property().UpdateProperty)
	propertyRouter.Delete("/delete/:id", httpHandler.Property().DeleteProperty)
	propertyRouter.Get("/get", httpHandler.Property().GetAllProperty)
	propertyRouter.Get("/get/:id", httpHandler.Property().GetPropertyByID)
	propertyRouter.Get("/search", httpHandler.Property().SearchProperty)
	propertyRouter.Get("/autocomplete", httpHandler.Property().AutoComplete)
}

func initLesseeRouter(router fiber.Router, httpHandler handlers.Handler, cfg *config.Config) {
	lesseeRouter := router.Group("/lessee", middleware.AuthRequired(cfg))
	lesseeRouter.Post("/create", httpHandler.Lessee().CreateReservation)
	lesseeRouter.Put("/update/:id", httpHandler.Lessee().UpdateReservation)
	lesseeRouter.Delete("/delete/:id", httpHandler.Lessee().DeleteReservation)
}

func initLessorRouter(router fiber.Router, httpHandler handlers.Handler, cfg *config.Config) {
	lessorRouter := router.Group("/lessor", middleware.AuthRequired(cfg))
	lessorRouter.Post("/accept/:id", httpHandler.Lessor().AcceptReservation)
	lessorRouter.Post("/decline/:id", httpHandler.Lessor().DeclineReservation)
}

func initPropertyReviewRouter(router fiber.Router, httpHandler handlers.Handler, cfg *config.Config) {
	propertyReviewRouter := router.Group("/propertyReview", middleware.AuthRequired(cfg))
	propertyReviewRouter.Post("/create", httpHandler.Review().CreateReview)
	propertyReviewRouter.Put("/update/:id", httpHandler.Review().UpdateReview)
	propertyReviewRouter.Delete("/delete/:id", httpHandler.Review().DeleteReview)
	propertyReviewRouter.Get("/get/:propertyID", httpHandler.Review().GetAllReviews)
}

func initPaymentRouter(router fiber.Router, httpHandler handlers.Handler, cfg *config.Config) {
	paymentRouter := router.Group("/payments", middleware.AuthRequired(cfg))
	paymentRouter.Post("/process", httpHandler.Payment().HandlePayment)
}

func initUserRouter(router fiber.Router, httpHandler handlers.Handler, cfg *config.Config) {
	userRouter := router.Group("/user", middleware.AuthRequired(cfg))

	userRouter.Put("/user", httpHandler.User().UpdateUser)
	userRouter.Put("/image", httpHandler.User().UpdateImage)
	userRouter.Post("/check", httpHandler.User().CheckUser)

}
