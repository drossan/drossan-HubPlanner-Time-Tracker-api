package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"hubplanner-proxy-api/domain/models"
	"hubplanner-proxy-api/usecases"
)

type LoginHandler struct {
	hubPlannerUseCase *usecases.LoginUseCase
}

func NewLoginHandler(uc *usecases.LoginUseCase) *LoginHandler {
	return &LoginHandler{
		hubPlannerUseCase: uc,
	}
}

func (h *LoginHandler) RegisterAuthLoginRoutes(g *echo.Group) {
	g.POST("/login", h.Login)
	g.GET("/login-google/:id-token", h.LoginGoogle)
}

func (h *LoginHandler) Login(c echo.Context) error {
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

func (h *LoginHandler) LoginGoogle(c echo.Context) error {
	idToken := c.Param("id-token")
	response, err := h.hubPlannerUseCase.LoginGoogle(idToken)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, response)
}
