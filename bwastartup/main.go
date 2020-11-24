package main

import (
	"bwastartup/user"
	"log"

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

	userInput := user.RegisterUserInput{}
	userInput.Name = "test simpan dari service"
	userInput.Email = "contoh@gmail.com"
	userInput.Occupation = "petanikode"
	userInput.Password = "passw0rd"

	userService.RegisterUser(userInput)

	// user := user.User{
	// 	Name: "Test Simpan",
	// }

	// userRepository.Save(user)
}
