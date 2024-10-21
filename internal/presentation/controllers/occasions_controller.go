package controllers

import (
	"errors"
	"net/http"

	"github.com/foxinuni/quickpass-backend/internal/domain/services"
	"github.com/labstack/echo/v4"
)

type OccasionsController struct {
	importService services.ImportService
}

func NewOccasionsController(importService services.ImportService) *OccasionsController {
	return &OccasionsController{
		importService: importService,
	}
}

func (oc *OccasionsController) Create(c echo.Context) error {
	// Get excel file from request
	file, err := c.FormFile("file")
	if err != nil {
		if errors.Is(err, http.ErrMissingFile) {
			return echo.NewHTTPError(http.StatusBadRequest, "excel file is required")
		}

		return err
	}

	// Read the file
	reader, err := file.Open()
	if err != nil {
		return err
	}

	// Import the occasions from the file
	count, err := oc.importService.ImportFromFile(reader)
	if err != nil {
		return err
	}
	// Return the number of occasions imported
	return c.JSON(http.StatusOK, map[string]interface{}{
		"number": count,
	})
}
