package handler

import (
	"basic_trade/admin"
	"basic_trade/common"
	"basic_trade/helper"
	"basic_trade/variant"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type variantHandler struct {
	service variant.Service
}

func NewVariantHandler(service variant.Service) *variantHandler {
	return &variantHandler{service}
}

func (h *variantHandler) GetVariants(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("search")

	variants, err := h.service.GetVariants(page, limit, search)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Error to get variants", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of variants", http.StatusOK, "success", common.FormatVariants(variants))
	c.JSON(http.StatusOK, response)
}

func (h *variantHandler) GetDetailVariant(c *gin.Context) {
	var input variant.GetVariantByUUIDInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Error to get detail variant", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if input.UUID == "" {
		response := helper.APIResponse("UUID cannot be empty", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	variantDetail, err := h.service.GetVariantByUUID(input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Error to get detail variant", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Detail of variant", http.StatusOK, "success", variantDetail)
	c.JSON(http.StatusOK, response)
}

func (h *variantHandler) CreateVariant(c *gin.Context) {
	var input variant.VariantInput

	input.VariantName = c.PostForm("variant_name")
	input.Quantity, _ = strconv.Atoi(c.PostForm("quantity"))
	input.ProductUUID = c.PostForm("product_id")

	currentAdmin, exists := c.MustGet("currentAdmin").(admin.Admin)
	input.Admin = currentAdmin

	if !exists {
		response := helper.APIResponse("You don't have permission", http.StatusUnauthorized, "error", nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	newVariant, err := h.service.CreateVariant(input)
	if err != nil {
		if err.Error() == "unauthorized" {
			response := helper.APIResponse("You don't have permission", http.StatusUnauthorized, "error", nil)
			c.JSON(http.StatusUnauthorized, response)
			return
		}

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create variant", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Variant has been created", http.StatusOK, "success", common.FormatVariant(*newVariant))
	c.JSON(http.StatusOK, response)
}

func (h *variantHandler) UpdateVariant(c *gin.Context) {
	var inputID variant.GetVariantByUUIDInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update variant", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var input variant.VariantInput

	input.VariantName = c.PostForm("variant_name")
	input.Quantity, _ = strconv.Atoi(c.PostForm("quantity"))

	currentAdmin, exists := c.MustGet("currentAdmin").(admin.Admin)
	input.Admin = currentAdmin

	if !exists {
		response := helper.APIResponse("You don't have permission", http.StatusUnauthorized, "error", nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	updatedVariant, err := h.service.UpdateVariant(inputID, input)
	if err != nil {
		if err.Error() == "unauthorized" {
			response := helper.APIResponse("You don't have permission", http.StatusUnauthorized, "error", nil)
			c.JSON(http.StatusUnauthorized, response)
			return
		}

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update variant", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Variant has been updated", http.StatusOK, "success", common.FormatVariant(*updatedVariant))
	c.JSON(http.StatusOK, response)
}

func (h *variantHandler) DeleteVariant(c *gin.Context) {
	var inputID variant.GetVariantByUUIDInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to parse UUID", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentAdmin, exists := c.MustGet("currentAdmin").(admin.Admin)
	if !exists {
		response := helper.APIResponse("You don't have permission", http.StatusUnauthorized, "error", nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	deletedVariant, err := h.service.DeleteVariant(inputID.UUID, currentAdmin.ID)
	if err != nil {
		if err.Error() == "unauthorized" {
			response := helper.APIResponse("You don't have permission", http.StatusUnauthorized, "error", nil)
			c.JSON(http.StatusUnauthorized, response)
			return
		}

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to delete variant", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Variant has been deleted", http.StatusOK, "success", common.FormatVariant(*deletedVariant))
	c.JSON(http.StatusOK, response)
}
