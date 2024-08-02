package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"hubplanner-proxy-api/adapters/routes"
	"hubplanner-proxy-api/config"
	"hubplanner-proxy-api/infrastructure/router"
)

func main() {
	/*******************************************************************************************************************
	LOAD DEFAULT CONFIGURATION
	*******************************************************************************************************************/
	cfg := config.LoadConfig()

	/*******************************************************************************************************************
	INIT ROUTER
	*******************************************************************************************************************/
	e, groupAPIRestricted, groupAPIAccessible, _ := router.NewEchoRouter(cfg.Server.JWTSecret)

	/*******************************************************************************************************************
	LOAD ROUTES
	*******************************************************************************************************************/
	routes.LoadRoutes(groupAPIRestricted, groupAPIAccessible)

	/*******************************************************************************************************************
	LAUNCH LAMBDA
	*******************************************************************************************************************/
	lambda.Start(echoadapter.New(e).ProxyWithContext)
}
