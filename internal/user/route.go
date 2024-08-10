package user

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ratheeshkumar/restaurant_gRPC_gatewayV1/middleware"
)

func NewUserRoutes(r *gin.Engine) {
	client, err := ClientDial()
	if err != nil {
		log.Fatalf("error not connected with gRPC server, %v", err.Error())
	}

	userHandler := &User{
		client: client,
	}

	apiUser := r.Group("/api/user")
	{
		apiUser.POST("/signup", userHandler.Signup)
		apiUser.POST("/login", userHandler.Login)
		apiUser.POST("/verify-otp", userHandler.VerifyOTP)
	}

	apiUserAuth := r.Group("/api/user/auth")
	apiUserAuth.Use(middleware.Authorization("user"))
	{
		apiUserAuth.GET("/getmenus", userHandler.FindAllMenu)
		apiUserAuth.GET("/getmenu", userHandler.FindMenu)
	}
}
