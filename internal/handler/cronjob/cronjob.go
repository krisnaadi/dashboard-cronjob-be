package cronjob

import (
	"context"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/krisnaadi/dashboard-cronjob-be/internal/usecase/cronjob"
	"github.com/krisnaadi/dashboard-cronjob-be/pkg/logger"
	"github.com/krisnaadi/dashboard-cronjob-be/pkg/writer"
	"github.com/labstack/echo/v4"
)

func (handler *Handler) HandleGetCronjob(c echo.Context) error {
	ctx := c.Request().Context()

	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		response := writer.APIResponse("JWT token missing or invalid", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		response := writer.APIResponse("failed to cast claims as jwt.MapClaims", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	cronjobs, err := handler.cronjob.ListCronjob(ctx, int64(claims["id"].(float64)))
	if err != nil {
		logger.Error(ctx, nil, err, "handler.cronjob.ListCronjob() error - HandleGetCronjob")
		response := writer.APIResponse("Get Cronjobs Fail", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	cronjobResponse := []CronjobResponse{}
	for _, cronjob := range cronjobs {
		cronjobResponse = append(cronjobResponse, CronjobResponse{
			ID:       cronjob.ID,
			Name:     cronjob.Name,
			Schedule: cronjob.Schedule,
			Task:     cronjob.Task,
			Status:   cronjob.Status,
		})
	}

	response := writer.APIResponse("Get Cronjob Successfully", true, cronjobResponse)
	return c.JSON(http.StatusOK, response)
}

func (handler *Handler) HandleShowCronjob(c echo.Context) error {
	ctx := c.Request().Context()
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		response := writer.APIResponse("JWT token missing or invalid", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		response := writer.APIResponse("failed to cast claims as jwt.MapClaims", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	ID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error(ctx, nil, err, "strconv.ParseInt() error - HandleShowCronjob")
		response := writer.APIResponse("ID is not valid", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	cronjob, err := handler.cronjob.GetCronjob(ctx, ID, int64(claims["id"].(float64)))
	if err != nil {
		logger.Error(ctx, nil, err, "handler.cronjob.GetCronjob() error - HandleShowCronjob")
		response := writer.APIResponse("Get Cronjob Fail", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if cronjob.ID == 0 {
		response := writer.APIResponse("Cronjob Not Found", false, nil)
		return c.JSON(http.StatusNotFound, response)
	}

	cronjobResponse := CronjobResponse{
		ID:       cronjob.ID,
		Name:     cronjob.Name,
		Schedule: cronjob.Schedule,
		Task:     cronjob.Task,
		Status:   cronjob.Status,
	}

	response := writer.APIResponse("Show Cronjob Successfully", true, cronjobResponse)
	return c.JSON(http.StatusOK, response)
}

func (handler *Handler) HandleCreateCronjob(c echo.Context) error {
	ctx := c.Request().Context()

	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		response := writer.APIResponse("JWT token missing or invalid", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		response := writer.APIResponse("failed to cast claims as jwt.MapClaims", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	var input cronjob.CronjobRequest
	if err := c.Bind(&input); err != nil {
		logger.Error(ctx, nil, err, "c.Bind() error - HandleCreateCronjob")
		response := writer.APIResponse("Store Cronjob Fail", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(input); err != nil {
		logger.Error(ctx, nil, err, "c.Validate() error - HandleCreateCronjob")
		response := writer.APIResponse("Store Cronjob Fail", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	cronjob, err := handler.cronjob.AddCronjob(ctx, input, int64(claims["id"].(float64)))
	if err != nil {
		logger.Error(ctx, nil, err, "handler.cronjob.AddCronjob() error - HandleCreateCronjob")
		response := writer.APIResponse("Store Cronjob Fail", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	cronjobResponse := CronjobResponse{
		ID:       cronjob.ID,
		Name:     cronjob.Name,
		Schedule: cronjob.Schedule,
		Task:     cronjob.Task,
		Status:   cronjob.Status,
	}

	response := writer.APIResponse("Store Cronjob Successfully", true, cronjobResponse)
	return c.JSON(http.StatusOK, response)
}

func (handler *Handler) HandleEditCronjob(c echo.Context) error {
	ctx := c.Request().Context()

	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		response := writer.APIResponse("JWT token missing or invalid", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		response := writer.APIResponse("failed to cast claims as jwt.MapClaims", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	ID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error(ctx, nil, err, "strconv.ParseInt() error - HandleEditCronjob")
		response := writer.APIResponse("ID is not valid", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	var input cronjob.CronjobRequest
	if err := c.Bind(&input); err != nil {
		logger.Error(ctx, nil, err, "c.Bind() error - HandleEditCronjob")
		response := writer.APIResponse("Update Cronjob Fail", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(input); err != nil {
		logger.Error(ctx, nil, err, "c.Validate() error - HandleEditCronjob")
		response := writer.APIResponse("Update Cronjob Fail", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	cronjob, err := handler.cronjob.UpdateCronjob(ctx, ID, input, int64(claims["id"].(float64)))
	if err != nil {
		logger.Error(ctx, nil, err, "handler.cronjob.UpdateCronjob() error - HandleEditCronjob")
		response := writer.APIResponse("Update Cronjob Fail", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	cronjobResponse := CronjobResponse{
		ID:       cronjob.ID,
		Name:     cronjob.Name,
		Schedule: cronjob.Schedule,
		Task:     cronjob.Task,
		Status:   cronjob.Status,
	}

	response := writer.APIResponse("Update Cronjob Successfully", true, cronjobResponse)
	return c.JSON(http.StatusOK, response)
}

func (handler *Handler) HandleDeleteCronjob(c echo.Context) error {
	ctx := c.Request().Context()

	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		response := writer.APIResponse("JWT token missing or invalid", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		response := writer.APIResponse("failed to cast claims as jwt.MapClaims", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	ID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error(ctx, nil, err, "strconv.ParseInt() error - HandleDeleteCronjob")
		response := writer.APIResponse("ID is not valid", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = handler.cronjob.DeleteCronjob(ctx, ID, int64(claims["id"].(float64)))
	if err != nil {
		logger.Error(ctx, nil, err, "handler.cronjob.DeleteCronjob() error - HandleDeleteCronjob")
		response := writer.APIResponse("Delete Cronjob Fail", false, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	response := writer.APIResponse("Delete Cronjob Successfully", true, nil)
	return c.JSON(http.StatusOK, response)
}

func (handler *Handler) HandleRunAllCronjob(ctx context.Context) error {
	return handler.cronjob.RunAllCronjob(ctx)
}
