package handler

import (
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

/**
1. tangkap input dari user
2. map input dari user ke struct RegisterUserInput
3. struct di atas kita passing sebagai parameter service
*/
func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	user, err := h.userService.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	c.JSON(http.StatusOK, user)
}
