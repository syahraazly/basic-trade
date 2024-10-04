package main

import (
	"basic_trade/admin"
	"basic_trade/auth"
	"basic_trade/handler"
	"basic_trade/helper"
	"basic_trade/product"
	"basic_trade/variant"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	dsn := "root:password@tcp(localhost:3306)/basic_trade?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	adminRepository := admin.NewRepository(db)
	productRepository := product.NewRepository(db)
	variantRepository := variant.NewRepository(db)

	adminService := admin.NewService(adminRepository)
	authService := auth.NewService()
	productService := product.NewService(productRepository)
	variantService := variant.NewService(variantRepository, productRepository)

	adminHandler := handler.NewAdminHandler(adminService, authService)
	productHandler := handler.NewProductHandler(productService)
	variantHandler := handler.NewVariantHandler(variantService)

	router := gin.Default()
	api := router.Group("/api")

	api.POST("/auth/register", adminHandler.Register)
	api.POST("/auth/login", adminHandler.Login)

	api.GET("/products", productHandler.GetProducts)
	api.GET("/products/:uuid", productHandler.GetDetailProduct)
	api.POST("/products", authMiddleware(authService, adminService), productHandler.CreateProduct)
	api.PUT("/products/:uuid", authMiddleware(authService, adminService), productHandler.UpdateProduct)
	api.DELETE("/products/:uuid", authMiddleware(authService, adminService), productHandler.DeleteProduct)

	api.GET("/products/variants", variantHandler.GetVariants)
	api.GET("/products/variants/:uuid", variantHandler.GetDetailVariant)
	api.POST("/products/variants", authMiddleware(authService, adminService), variantHandler.CreateVariant)
	api.PUT("/products/variants/:uuid", authMiddleware(authService, adminService), variantHandler.UpdateVariant)
	api.DELETE("/products/variants/:uuid", authMiddleware(authService, adminService), variantHandler.DeleteVariant)

	router.Run()
}

func authMiddleware(authService auth.Service, adminService admin.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		adminID := int(claim["admin_id"].(float64)) // Ambil admin_id dari token

		admin, err := adminService.GetByID(adminID) // Ambil admin berdasarkan ID
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentAdmin", admin) // Simpan admin ke context
		c.Next()
	}
}
