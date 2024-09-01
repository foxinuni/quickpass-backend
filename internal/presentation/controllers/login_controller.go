package controllers

import (
	"net/http"

	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
	"github.com/foxinuni/quickpass-backend/internal/domain/services"
	"github.com/foxinuni/quickpass-backend/internal/presentation/dtos"
	"github.com/labstack/echo/v4"
)

type LoginController struct {
	authService services.AuthService
}

func NewLoginController(authService services.AuthService) *LoginController {
	return &LoginController{
		authService: authService,
	}
}

func (lc *LoginController) Login(c echo.Context) error {
	// Bind the request body to the LoginDTO
	var login dtos.LoginDTO
	if err := c.Bind(&login); err != nil {
		return err
	}

	// Validate the LoginDTO
	if err := c.Validate(&login); err != nil {
		return err
	}

	// Call the login service
	session, err := lc.authService.Login(login.Email, login.Number, login.PhoneModel, login.PhoneEMEI)
	if err != nil {
		if err == services.ErrInvalidCredentials {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid credentials")
		}

		return err
	}

	// Return the session
	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": session.Token,
	})
}

func (lc *LoginController) Logout(c echo.Context) error {
	// Get session from context
	session, ok := c.Get("session").(*entities.Session)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "session required but not found")
	}

	// Call the logout service
	if err := lc.authService.Logout(session); err != nil {
		return err
	}

	// Return success
	return c.NoContent(http.StatusOK)
}
