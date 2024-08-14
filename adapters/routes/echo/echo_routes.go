package echo

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
	loginRepo := repository.NewLoginConnectionRepository()
	hubPlannerRepo := repository.NewHubPlannerConnectionRepository()

	/*******************************************************************************************************************
	INITIALISING USE CASES
	*******************************************************************************************************************/
	healthUseCase := usecases.NewHealthUseCase(healthRepo)
	loginUseCase := usecases.NewLoginUserUseCase(loginRepo)
	hubPlannerUseCase := usecases.NewHubPlannerUserUseCase(hubPlannerRepo)
	/*******************************************************************************************************************
	INITIALISING HANDLERS
	*******************************************************************************************************************/
	healthHandler := NewHealthHandler(healthUseCase)
	loginHandler := NewLoginHandler(loginUseCase)
	hubPlannerHandler := NewHubPlannerHandler(hubPlannerUseCase)

	/*******************************************************************************************************************
	REGISTER PUBLIC ROUTES
	*******************************************************************************************************************/
	healthHandler.RegisterRoutes(groupAPIAccessible)
	loginHandler.RegisterAuthLoginRoutes(groupAPIAccessible)

	/*******************************************************************************************************************
	REGISTER RESTRICTED ROUTES
	*******************************************************************************************************************/
	hubPlannerHandler.RegisterHubPlannerRoutes(groupAPIRestricted)
}
