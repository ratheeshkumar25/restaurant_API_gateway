package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/ratheeshkumar/restaurant_gRPC_gatewayV1/internal/admin/handler"
	adminpb "github.com/ratheeshkumar/restaurant_gRPC_gatewayV1/internal/admin/pb"
)

type AdminRoutes struct {
	client adminpb.AdminServiceClient
}

func (a *AdminRoutes) Login(c *gin.Context) {
	handler.AdminLoginHandler(c, a.client)
}

func (a *AdminRoutes) CreateMenu(c *gin.Context) {
	handler.CreateBookHandler(c, a.client)
}
func (a *AdminRoutes) FindMenu(c *gin.Context) {
	handler.FindMenuHandler(c, a.client)
}
func (a *AdminRoutes) FindAllMenu(c *gin.Context) {
	handler.FetchAllMenuHandler(c, a.client)
}
