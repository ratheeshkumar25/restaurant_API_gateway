package di

import (
	"github.com/ratheeshkumar/restaurant_gRPC_gatewayV1/config"
	"github.com/ratheeshkumar/restaurant_gRPC_gatewayV1/internal/admin"
	"github.com/ratheeshkumar/restaurant_gRPC_gatewayV1/internal/server"
	"github.com/ratheeshkumar/restaurant_gRPC_gatewayV1/internal/user"
)

func Init() {
	server := server.Server()
	config.LoadConfig()
	user.NewUserRoutes(server.R)
	admin.NewAdminRoute(server.R)
	server.StartServer()
}
