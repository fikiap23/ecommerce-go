package handlers

import (
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/repository"
	"go-ecommerce-app/internal/service"
	"go-ecommerce-app/pkg/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	// svc service.UserService
	svc service.UserService
}

func SetupUserRoutes(rh *rest.RestHandler) {

	app:= rh.App

	// create an instance of user service & inject to handler
	
	// Buat repository dari db yg ada di RestHandler
    userRepo := repository.NewUserRepository(rh.DB)

    // Buat service dengan inject repository
    userService := service.NewUserService(userRepo)
	handler:=UserHandler{
		svc: *userService,
	}

	// Public endpoints
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	// Private endpoints
	app.Get("/verify", handler.GetVerificationCode)
	app.Post("/verify", handler.Verify)
	app.Post("/profile", handler.CreateProfile)
	app.Get("/profile", handler.GetProfile)

	app.Post("/cart", handler.AddToCart)
	app.Get("/cart", handler.GetCart)
	app.Post("/order", handler.CreateOrder)
	app.Get("/order", handler.GetOrder)
	app.Get("/order/:id", handler.GetOrderById)

	app.Post("/become-seller", handler.BecomeSeller)

	
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	var input dto.UserSignup
	if err := utils.ParseAndValidate(ctx, &input); err != nil {
		return utils.HandleError(ctx, err)
	}

	lang := utils.GetLanguageFromHeader(ctx)
	token, err := h.svc.Signup(input, lang)
	if err != nil {
		return utils.HandleError(ctx, err)
	}

	return utils.SuccessResponse(ctx, 201, http.StatusText(http.StatusCreated), fiber.Map{
		"token": token,
	})
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	var input dto.UserLogin

	if err := utils.ParseAndValidate(ctx, &input); err != nil {
		return utils.HandleError(ctx, err)
	}

	lang := utils.GetLanguageFromHeader(ctx)

	token, err := h.svc.Login(input, lang)
	if err != nil {
		return utils.HandleError(ctx, err)
	}

	return utils.SuccessResponse(ctx, 200, http.StatusText(http.StatusOK), fiber.Map{
		"token": token,
	})
}

func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "get verification code"})
}

func (h *UserHandler) Verify(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "verify"})
}

func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "create profile"})
}

func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "get profile"})
}

func (h *UserHandler) AddToCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "add to cart"})
}

func (h *UserHandler) GetCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "get cart"})
}

func (h *UserHandler) CreateOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "create order"})
}

func (h *UserHandler) GetOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "get order"})
}

func (h *UserHandler) GetOrderById(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "get order by id"})
}

func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "become seller"})
}