package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ratheeshkumar/restaurant_gRPC_gatewayV1/internal/model"
	userpb "github.com/ratheeshkumar/restaurant_gRPC_gatewayV1/internal/user/pb"
)

func UserSignupHandler(c *gin.Context, client userpb.UserServicesClient) {
	ctxt, cancel := context.WithTimeout(c, time.Second*2000)
	defer cancel()

	var user model.UserModel

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("error binding json :%v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	response, err := client.Signup(ctxt, &userpb.SignupRequest{
		Phone: user.Phone,
	})

	if err != nil {
		log.Printf("error signing up user %v err: %v", user.Phone, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v otp generated successfully", user.Phone),
		"data":    response,
	})
}

func UserVerifyOTPHandler(c *gin.Context, client userpb.UserServicesClient) {
	ctxt, cancel := context.WithTimeout(c, time.Second*2000)
	defer cancel()

	var request *model.VerifyOTP

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("error binding json :%v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	response, err := client.VerifyOTP(ctxt, &userpb.VerifyOTPRequest{
		Phone: request.Phone,
		Otp:   request.Otp,
	})

	if err != nil {
		log.Printf("error verifying OTP for user %v err: %v", request.Phone, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": "OTP verified successfully",
		"token":   response.Token,
	})
}

func UserLoginHandler(c *gin.Context, client userpb.UserServicesClient) {
	ctxt, cancel := context.WithTimeout(c, time.Second*2000)
	defer cancel()

	var user model.UserModel

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("error binding json :%v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	response, err := client.Login(ctxt, &userpb.LoginRequest{
		Phone: user.Phone,
	})

	if err != nil {
		log.Printf("error signing up user %v err: %v", user.Phone, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v login  successfully", user.Phone),
		"data":    response,
		"token":   response.Token,
	})

}
