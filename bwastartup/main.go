package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:password@tcp(127.0.0.1:3308)/bwatstartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)

	router.Run(":3000")
}

/**
single test direct to service, repo
*/

// var email string = "contoh@gmail.com"
// loginUserInput := user.LoginUserInput{}
// loginUserInput.Email = email
// loginUserInput.Password = ""
// user, err := userService.LoginUser(loginUserInput)
// fmt.Println("User Service Find by email ", user)

// user, err := userRepository.FindByEmail(email)
// fmt.Println(user.Name)

// userInput := user.RegisterUserInput{}
// userInput.Name = "test simpan dari service"
// userInput.Email = "contoh@gmail.com"
// userInput.Occupation = "petanikode"
// userInput.Password = "passw0rd"

// userService.RegisterUser(userInput)

// user := user.User{
// 	Name: "Test Simpan",
// }

// userRepository.Save(user)
