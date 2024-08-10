package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	userpb "github.com/ratheeshkumar/restaurant_gRPC_gatewayV1/internal/user/pb"
)

func FetchAllMenuHandler(c *gin.Context, client userpb.UserServicesClient) {
	ctxt, cancel := context.WithTimeout(c, time.Second*2000)
	defer cancel()

	response, err := client.UserMenuList(ctxt, &userpb.RNoparam{})
	if err != nil {
		log.Printf("error finding all menulist  err: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": "fetched all books successfully",
		"data":    response,
	})

}

func FindMenuHandler(c *gin.Context, client userpb.UserServicesClient) {
	ctxt, cancel := context.WithTimeout(c, time.Second*2000)
	defer cancel()

	id := c.Query("id")
	name := c.Query("name")

	response := &userpb.MenuItem{}

	var err error

	if id == "" && name == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "invalid query",
		})
		return
	} else if id != "" {
		id, err := strconv.Atoi(id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  "invalid id",
			})
			return
		}
		response, err = client.UserMenuByID(ctxt, &userpb.MenuID{Id: uint32(id)})
		if err != nil {
			log.Printf("error finding  menu by id err: %v", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  err.Error(),
			})
			return
		}
	} else if name != "" {
		response, err = client.UserFoodByName(ctxt, &userpb.FoodByName{Name: name})
		if err != nil {
			log.Printf("error finding  menu by name err: %v", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": "fetched all menuItems successfully",
		"data":    response,
	})

}
