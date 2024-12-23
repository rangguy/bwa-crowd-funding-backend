package main

import (
	"bwastartup/user"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	//dsn := "host=localhost user=postgres password=gansboy29 dbname=bwastartup port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//fmt.Println("Connection Success")
	//
	//var users []user.User
	//
	//db.Find(&users)
	//
	//for _, user := range users {
	//	fmt.Println(user.ID)
	//	fmt.Println(user.Name)
	//	fmt.Println(user.Occupation)
	//	fmt.Println(user.Email)
	//	fmt.Println(user.AvatarFileName)
	//	fmt.Println(user.Role)
	//	fmt.Println("================")
	//}

	router := gin.Default()
	router.GET("/handler", handler)
	err := router.Run()
	if err != nil {
		return
	}
}

func handler(c *gin.Context) {
	dsn := "host=localhost user=postgres password=gansboy29 dbname=bwastartup port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	var users []user.User
	db.Find(&users)

	c.JSON(http.StatusOK, users)
}
