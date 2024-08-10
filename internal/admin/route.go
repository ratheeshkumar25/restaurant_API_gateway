package admin

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ratheeshkumar/restaurant_gRPC_gatewayV1/middleware"
)

func NewAdminRoute(r *gin.Engine) {
	client, err := ClientDial()

	if err != nil {
		log.Fatalf("error Not connected with gRPC server, %v", err.Error())
	}

	adminHandler := &AdminRoutes{
		client: client,
	}

	apiAdmin := r.Group("/api/admin")

	{
		apiAdmin.POST("/login", adminHandler.Login)
	}
	apiAuthAdmin := r.Group("/api/admin/auth")
	apiAuthAdmin.Use(middleware.Authorization("admin"))
	{
		apiAuthAdmin.POST("/menu", adminHandler.CreateMenu)
		apiAuthAdmin.GET("/menus", adminHandler.FindAllMenu)
		apiAuthAdmin.GET("/menu", adminHandler.FindMenu)
	}
}
