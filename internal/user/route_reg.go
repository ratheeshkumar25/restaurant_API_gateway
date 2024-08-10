package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ratheeshkumar/restaurant_gRPC_gatewayV1/internal/user/handler"
	userpb "github.com/ratheeshkumar/restaurant_gRPC_gatewayV1/internal/user/pb"
)

type User struct {
	client userpb.UserServicesClient
}

func (u *User) Signup(c *gin.Context) {
	handler.UserSignupHandler(c, u.client)
}
func (u *User) Login(c *gin.Context) {
	handler.UserLoginHandler(c, u.client)
}
func (u *User) VerifyOTP(c *gin.Context) {
	handler.UserVerifyOTPHandler(c, u.client)
}
func (u *User) FindMenu(c *gin.Context) {
	handler.FindMenuHandler(c, u.client)
}
func (u *User) FindAllMenu(c *gin.Context) {
	handler.FetchAllMenuHandler(c, u.client)
}
