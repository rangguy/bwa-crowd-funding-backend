package main

import (
	"bwastartup/auth"
	"bwastartup/handler"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
)

func main() {
	dsn := "host=localhost user=postgres password=postgresql dbname=bwastartup port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Connection Success")

	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	err = router.Run()
	if err != nil {
		return
	}
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Bearer "token"
		tokenString := strings.Split(authHeader, " ")[1]

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}

// ambil nilai header Authorization, Bearer "token"
// ambil nilai "token" dari header
// memvalidasi token
// jika token valid maka mengambil nilai user_id
// ambil user dari db berdasarkan user_id lewat service
// jika user ada maka set context (sebuah tempat untuk menyimpan suatu nilai yang nantinya nilai bisa diambil oleh yang lain) yang isinya user
