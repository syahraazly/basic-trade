package handler

import (
	"basic_trade/admin"
	"basic_trade/cloudinary"
	"basic_trade/helper"
	"basic_trade/product"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	service product.Service
}

func NewProductHandler(service product.Service) *productHandler {
	return &productHandler{service}
}

var allowedFileTypes = []string{"image/jpeg", "image/png", "image/svg+xml"}

const maxFileSize = 5 * 1024 * 1024

func (h *productHandler) GetProducts(c *gin.Context) {
	adminID, _ := strconv.Atoi(c.Query("admin_id"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("search")

	products, err := h.service.GetProducts(adminID, page, limit, search)
	if err != nil {
		response := helper.APIResponse("Error to get products", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of products", http.StatusOK, "success", product.FormatProducts(products))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) GetDetailProduct(c *gin.Context) {
	var input product.GetProductByUUIDInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Error to get detail product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if input.UUID == "" {
		response := helper.APIResponse("UUID cannot be empty", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	productDetail, err := h.service.GetProductByUUID(input)
	if err != nil {
		response := helper.APIResponse("Error to get detail product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Detail of product", http.StatusOK, "success", product.FormatProductDetail(*productDetail))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) CreateProduct(c *gin.Context) {
	var input product.ProductInput

	input.Name = c.PostForm("name")
	if input.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	cloudinaryService, err := cloudinary.NewService()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	imageURL, err := cloudinaryService.UploadFile(file, fileHeader)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentAdmin := c.MustGet("currentAdmin").(admin.Admin)

	input.ImageURL = imageURL
	input.Admin = currentAdmin

	newProduct, err := h.service.CreateProduct(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := helper.APIResponse("Product has been created", http.StatusOK, "success", product.FormatProduct(*newProduct))
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) UpdateProduct(c *gin.Context) {
	var inputID product.GetProductByUUIDInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input product.ProductInput
	input.Name = c.PostForm("name")
	if input.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	file, fileHeader, err := c.Request.FormFile("file")
	if err == nil {
		defer file.Close()

		cloudinaryService, err := cloudinary.NewService()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		imageURL, err := cloudinaryService.UploadFile(file, fileHeader)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		input.ImageURL = imageURL
	}

	currentAdmin := c.MustGet("currentAdmin").(admin.Admin)
	input.Admin = currentAdmin

	updatedProduct, err := h.service.UpdateProduct(inputID, input)
	if err != nil {
		response := helper.APIResponse("Failed to update product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Product has been updated", http.StatusOK, "success", product.FormatProduct(*updatedProduct))
	c.JSON(http.StatusOK, response)
}

func isValidFileType(fileType string) bool {
	for _, allowedType := range allowedFileTypes {
		if allowedType == fileType {
			return true
		}
	}
	return false
}

func (h *productHandler) DeleteProduct(c *gin.Context) {
	var inputID product.GetProductByUUIDInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to parse UUID", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentAdmin := c.MustGet("currentAdmin").(admin.Admin)

	deletedProduct, err := h.service.DeleteProduct(inputID.UUID, currentAdmin.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Product has been deleted", http.StatusOK, "success", product.FormatProduct(*deletedProduct))
	c.JSON(http.StatusOK, response)
}
