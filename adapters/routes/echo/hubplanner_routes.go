package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
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

func (h *HubPlannerHandler) RegisterHubPlannerRoutes(g *echo.Group) {
	g.GET("/projects", h.Projects)
	g.GET("/categories", h.Categories)
	g.POST("/timeentry", h.TimeEntry)
	g.GET("/timeentry/submit/:timeEntryID", h.TimeEntrySubmit)
	g.GET("/timeentries", h.TimeEntries)
}

func (h *HubPlannerHandler) Projects(c echo.Context) error {
	response, err := h.hubPlannerUseCase.Projects()
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

func (h *HubPlannerHandler) TimeEntrySubmit(c echo.Context) error {
	resourceID := GetUserID(c)
	timeEntryID := c.Param("timeEntryID")
	response, err := h.hubPlannerUseCase.TimeEntrySubmit(timeEntryID, resourceID)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, response)
}

func (h *HubPlannerHandler) TimeEntries(c echo.Context) error {
	repositoryID := GetUserID(c)
	response, err := h.hubPlannerUseCase.TimeEntries(repositoryID)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, response)
}
