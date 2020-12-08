package main

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/handler"
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"

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
	campaignRepository := campaign.NewRepository(db)
	campaignService := campaign.NewService(campaignRepository)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	api.GET("/campaigns", campaignHandler.GetCampaigns)

	router.Run(":3000")
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) // stop proses
			return
		}

		// Bearer token

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)

		if err != nil {
			response := helper.APIResponse("Unauthorized 1", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) // stop proses
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized 2", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) // stop proses
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized 3", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) // stop proses
			return
		}

		c.Set("currentUser", user)

	}
}

/**


	campaigns, err := campaignRepository.FindAll()
	fmt.Println(campaigns, err)

	campaigns, err = campaignRepository.FindByUserID(1)
	for _, campaign := range campaigns {

		if len(campaign.CampaignImages) > 0 {
			fmt.Println(campaign.CampaignImages[0].FileName)
		}
	}

	fmt.Println(campaigns[0].CampaignImages, err)

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
