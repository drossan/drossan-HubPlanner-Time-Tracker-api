package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"hubplanner-proxy-api/usecases"
)

type HealthHandler struct {
	healthUseCase *usecases.HealthUseCase
}

func NewHealthHandler(uc *usecases.HealthUseCase) *HealthHandler {
	return &HealthHandler{healthUseCase: uc}
}

func (h *HealthHandler) RegisterRoutes(g *echo.Group) {
	g.GET("/health", h.CheckStatusAPI)
}

func (h *HealthHandler) CheckStatusAPI(c echo.Context) error {
	status, err := h.healthUseCase.CheckStatus()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, status)
}
