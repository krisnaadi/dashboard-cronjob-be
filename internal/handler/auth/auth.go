package auth

import (
	"net/http"

	"github.com/krisnaadi/dashboard-cronjob-be/internal/usecase/auth"
	"github.com/krisnaadi/dashboard-cronjob-be/pkg/logger"
	"github.com/krisnaadi/dashboard-cronjob-be/pkg/writer"
	"github.com/labstack/echo/v4"
)

func (handler *Handler) HandleShowUser(c echo.Context) error {
	ctx := c.Request().Context()

	user, err := handler.auth.GetAuthUser(ctx, 1)
	if err != nil {
		logger.Error(ctx, nil, err, "handler.auth.GetAuthUser() error - HandleShowUser")
		response := writer.APIResponse("Get User Fail", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if user.ID == 0 {
		response := writer.APIResponse("User Not Found", false, nil)
		return c.JSON(http.StatusNotFound, response)
	}

	userResponse := UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	response := writer.APIResponse("Show User Successfully", true, userResponse)
	return c.JSON(http.StatusOK, response)
}

func (handler *Handler) HandleLogin(c echo.Context) error {
	ctx := c.Request().Context()

	var input auth.LoginRequest
	if err := c.Bind(&input); err != nil {
		logger.Error(ctx, nil, err, "c.Bind() error - HandleLogin")
		response := writer.APIResponse("Login Fail", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(input); err != nil {
		logger.Error(ctx, nil, err, "c.Validate() error - HandleLogin")
		response := writer.APIResponse("Login Fail", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	user, err := handler.auth.Login(ctx, input)
	if err != nil {
		logger.Error(ctx, nil, err, "handler.auth.Login() error - HandleLogin")
		response := writer.APIResponse("Email or Password Wrong", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	token, err := handler.auth.GenerateToken(ctx, user)
	if err != nil {
		logger.Error(ctx, nil, err, "handler.auth.GenerateToken() error - HandleLogin")
		response := writer.APIResponse("Generate Token Fail", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	loginResponse := LoginResponse{
		Token: token,
	}

	response := writer.APIResponse("Login Successfully", true, loginResponse)
	return c.JSON(http.StatusOK, response)
}

func (handler *Handler) HandleRegister(c echo.Context) error {
	ctx := c.Request().Context()

	var input auth.RegisterRequest
	if err := c.Bind(&input); err != nil {
		logger.Error(ctx, nil, err, "c.Bind() error - HandleRegister")
		response := writer.APIResponse("Register Fail", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(input); err != nil {
		logger.Error(ctx, nil, err, "c.Validate() error - HandleRegister")
		response := writer.APIResponse("Register Fail", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	_, err := handler.auth.Register(ctx, input)
	if err != nil {
		logger.Error(ctx, nil, err, "handler.auth.Register() error - HandleRegister")
		response := writer.APIResponse("Register Fail", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	response := writer.APIResponse("Register Successfully", true, nil)
	return c.JSON(http.StatusOK, response)
}
