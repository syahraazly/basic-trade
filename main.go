package main

import (
	"basic_trade/admin"
	"basic_trade/auth"
	"basic_trade/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:password@tcp(localhost:3306)/basic_trade?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	adminRepository := admin.NewRepository(db)

	adminService := admin.NewService(adminRepository)
	authService := auth.NewService()

	adminHandler := handler.NewAdminHandler(adminService, authService)

	router := gin.Default()
	api := router.Group("/api")

	api.POST("/auth/register", adminHandler.Register)
	api.POST("/auth/login", adminHandler.Login)

	router.Run()
}
