package main

import (
	user "bwastartup/user/users"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=postgres password=gansboy29 dbname=bwastartup port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connection Success")

	var users []user.User
	length := len(users)

	fmt.Println(length)

	db.Find(&users)

	length = len(users)
	fmt.Println(length)

	for _, user := range users {
		fmt.Println(user.ID)
		fmt.Println(user.Name)
		fmt.Println(user.Occupation)
		fmt.Println(user.Email)
		fmt.Println(user.AvatarFileName)
		fmt.Println(user.Role)
		fmt.Println("================")
	}
}
