package routes

import (
	"github.com/labstack/echo/v4"
	"hubplanner-proxy-api/infrastructure/repository"
	"hubplanner-proxy-api/usecases"
)

func LoadRoutes(
	groupAPIRestricted *echo.Group,
	groupAPIAccessible *echo.Group,
) {
	/*******************************************************************************************************************
	CERATE INSTANCES OF REPOSITORIES
	*******************************************************************************************************************/
	healthRepo := repository.NewHealthConnectionRepository()
	hubPlannerRepo := repository.NewHubPlannerConnectionRepository()

	/*******************************************************************************************************************
	INITIALISING USE CASES
	*******************************************************************************************************************/
	healthUseCase := usecases.NewUserUseCase(healthRepo)
	hubPlannerUseCase := usecases.NewHubPlannerUserUseCase(hubPlannerRepo)
	/*******************************************************************************************************************
	INITIALISING HANDLERS
	*******************************************************************************************************************/
	healthHandler := NewHealthHandler(healthUseCase)
	hubPlannerHandler := NewHubPlannerHandler(hubPlannerUseCase)

	/*******************************************************************************************************************
	REGISTER PUBLIC ROUTES
	*******************************************************************************************************************/
	healthHandler.RegisterRoutes(groupAPIAccessible)
	hubPlannerHandler.RegisterAuthHubPlannerRoutes(groupAPIAccessible)

	/*******************************************************************************************************************
	REGISTER RESTRICTED ROUTES
	*******************************************************************************************************************/
	hubPlannerHandler.RegisterHubPlannerRoutes(groupAPIRestricted)
}
