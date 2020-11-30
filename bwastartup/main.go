package main

import (
	"bwastartup/auth"
	"bwastartup/handler"
	"bwastartup/user"
	"fmt"

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
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)
	token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxfQ.pR5mUyz1tm_Ni6-mCi-ankpmIwVifpJ0k_tNjbyp6p8")
	fmt.Println(token)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run(":3000")
}

/**
single test direct to service, repo
*/
// authService := auth.NewService()
// 	token, err := authService.GenerateToken(1001)

// 	fmt.Println(token)

// userService.SaveAvatar(1, "images/1-profie.png")

// user, err := userRepository.FindUserByID(1)
// fmt.Println(user.Name)

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
