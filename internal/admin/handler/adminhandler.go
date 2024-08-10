package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	adminpb "github.com/ratheeshkumar/restaurant_gRPC_gatewayV1/internal/admin/pb"
	"github.com/ratheeshkumar/restaurant_gRPC_gatewayV1/internal/model"
)

func AdminLoginHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	var admin model.AdminModel

	if err := c.ShouldBindJSON(&admin); err != nil {
		log.Printf("error binding JSON")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	ctx := context.Background()

	response, err := client.AdminLogin(ctx, &adminpb.AdminRequest{
		Username: admin.Username,
		Password: admin.Password,
	})
	if err != nil {
		log.Printf("error logging in admin %v err: %v", admin.Username, err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v logged in succesfully", admin.Username),
		"data":    response,
	})
}
