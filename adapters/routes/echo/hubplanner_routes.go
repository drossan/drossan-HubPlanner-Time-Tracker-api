package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"hubplanner-proxy-api/domain/models"
	"hubplanner-proxy-api/domain/models/HubPlanner"
	"hubplanner-proxy-api/usecases"
)

type HubPlannerHandler struct {
	hubPlannerUseCase *usecases.HubPlannerUseCase
}

func NewHubPlannerHandler(uc *usecases.HubPlannerUseCase) *HubPlannerHandler {
	return &HubPlannerHandler{
		hubPlannerUseCase: uc,
	}
}

func (h *HubPlannerHandler) RegisterAuthHubPlannerRoutes(g *echo.Group) {
	g.POST("/login", h.Login)
}

func (h *HubPlannerHandler) RegisterHubPlannerRoutes(g *echo.Group) {
	g.GET("/projects", h.Projects)
	g.GET("/categories", h.Categories)
	g.POST("/timeentry", h.TimeEntry)
}

func (h *HubPlannerHandler) Login(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Email or password is required"})
	}

	response, err := h.hubPlannerUseCase.Login(user.Username, user.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, response)
}

func (h *HubPlannerHandler) Projects(c echo.Context) error {
	resourceID := GetUserID(c)
	response, err := h.hubPlannerUseCase.Projects(resourceID)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, response)
}

func (h *HubPlannerHandler) Categories(c echo.Context) error {
	response, err := h.hubPlannerUseCase.Categories()
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, response)
}

func (h *HubPlannerHandler) TimeEntry(c echo.Context) error {
	timeEntry := new(HubPlanner.TimeEntry)
	if err := c.Bind(timeEntry); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	timeEntry.Resource = GetUserID(c)
	response, err := h.hubPlannerUseCase.TimeEntry(timeEntry)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, response)
}
