package handlers

import (
	"go-ecommerce-app/internal/api/rest"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/helper"
	"go-ecommerce-app/internal/repository"
	"go-ecommerce-app/internal/service"
	"go-ecommerce-app/pkg/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	svc  service.UserService
	auth helper.Auth
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func SetupUserRoutes(rh *rest.RestHandler) {

	app:= rh.App

	// create an instance of user service & inject to handler
	
	// Buat repository dari db yg ada di RestHandler
    userRepo := repository.NewUserRepository(rh.DB)

    // Buat service dengan inject repository
    userService := service.NewUserService(userRepo, rh.Auth, rh.Config)
	handler := UserHandler{
		svc: userService, 
		auth: rh.Auth,
	}

	publicRoutes:= app.Group("/users")
	// Public endpoints
	publicRoutes.Post("/register", handler.Register)
	publicRoutes.Post("/login", handler.Login)

	privateRoutes:= app.Group("/users", rh.Auth.Authorize)
	// Private endpoints
	privateRoutes.Get("/verify", handler.GetVerificationCode)
	privateRoutes.Post("/verify", handler.Verify)
	privateRoutes.Post("/profile", handler.CreateProfile)
	privateRoutes.Get("/profile", handler.GetProfile)

	privateRoutes.Post("/cart", handler.AddToCart)
	privateRoutes.Get("/cart", handler.GetCart)
	privateRoutes.Post("/order", handler.CreateOrder)
	privateRoutes.Get("/order", handler.GetOrder)
	privateRoutes.Get("/order/:id", handler.GetOrderById)

	privateRoutes.Post("/become-seller", handler.BecomeSeller)

	
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	input:= dto.UserSignup{}
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
	 input:= dto.UserLogin{}

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
	lang := utils.GetLanguageFromHeader(ctx)
	userId:= h.auth.GetCurrentUser(ctx).Sub

	 err := h.svc.GetVerificationCode(userId, lang)
	if err != nil {
		return utils.HandleError(ctx, err)
	}

	return utils.SuccessResponse(ctx, 200, http.StatusText(http.StatusOK),nil)
}

func (h *UserHandler) Verify(ctx *fiber.Ctx) error {
	input:= dto.UserVerification{}
	lang := utils.GetLanguageFromHeader(ctx)
	userId:= h.auth.GetCurrentUser(ctx).Sub


	if err := utils.ParseAndValidate(ctx, &input); err != nil {
		return utils.HandleError(ctx, err)
	}

	err := h.svc.VerifyCode(userId, input.Code, lang)

	if err != nil {
		return utils.HandleError(ctx, err)
	}

	return utils.SuccessResponse(ctx, 200, http.StatusText(http.StatusOK),nil)

}

func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": "create profile"})
}

func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	lang := utils.GetLanguageFromHeader(ctx)
	idUser := h.auth.GetCurrentUser(ctx).Sub

	user, err := h.svc.GetProfile(idUser, lang)

	if err != nil {
		return utils.HandleError(ctx, err)
	}

	return utils.SuccessResponse(ctx, 200, http.StatusText(http.StatusOK), user)
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