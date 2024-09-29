package handler

import (
	"basic_trade/admin"
	"basic_trade/auth"
	"basic_trade/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type adminHandler struct {
	adminService admin.Service
	authService  auth.Service
}

func NewAdminHandler(adminService admin.Service, authService auth.Service) *adminHandler {
	return &adminHandler{adminService, authService}
}

func (h *adminHandler) Register(c *gin.Context) {
	var input admin.RegisterInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register Account Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newAdmin, err := h.adminService.Register(input)
	if err != nil {
		response := helper.APIResponse("Register Account Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newAdmin.ID)
	if err != nil {
		response := helper.APIResponse("Register Account Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := admin.FormatAdmin(&newAdmin)

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", gin.H{
		"admin": formatter,
		"token": token,
	})

	c.JSON(http.StatusOK, response)
}

func (h *adminHandler) Login(c *gin.Context) {
	var input admin.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loginAdmin, err := h.adminService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(loginAdmin.ID)
	if err != nil {
		response := helper.APIResponse("Login Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := admin.FormatAdmin(&loginAdmin)

	response := helper.APIResponse("Login Success", http.StatusOK, "success", gin.H{
		"admin": formatter,
		"token": token,
	})

	c.JSON(http.StatusOK, response)
}
