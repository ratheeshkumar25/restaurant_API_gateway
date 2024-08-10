package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	adminpb "github.com/ratheeshkumar/restaurant_gRPC_gatewayV1/internal/admin/pb"
	"github.com/ratheeshkumar/restaurant_gRPC_gatewayV1/internal/model"
)

func CreateBookHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	ctxt, cancel := context.WithTimeout(c, time.Second*2000)
	defer cancel()

	var menu model.Menu

	if err := c.ShouldBindJSON(&menu); err != nil {
		log.Printf("error binding json :%v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	response, err := client.CreateMenu(ctxt, &adminpb.AMenuItem{
		Category:  menu.Category,
		Name:      menu.Name,
		Price:     menu.Price,
		Foodimage: menu.FoodImage,
		Duration:  menu.Duration,
	})

	if err != nil {
		log.Printf("error creating menuItems %v err: %v", menu.Name, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusAccepted,
		"message": fmt.Sprintf("%v created successfully", menu.Name),
		"data":    response,
	})
}

func FetchAllMenuHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	ctxt, cancel := context.WithTimeout(c, time.Second*2000)
	defer cancel()

	response, err := client.FetchMenus(ctxt, &adminpb.AdminNoParam{})
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

func FindMenuHandler(c *gin.Context, client adminpb.AdminServiceClient) {
	ctxt, cancel := context.WithTimeout(c, time.Second*2000)
	defer cancel()

	id := c.Query("id")
	name := c.Query("name")

	response := &adminpb.AMenuItem{}

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
		response, err = client.FetchByMenuID(ctxt, &adminpb.AMenuBYId{Id: uint32(id)})
		if err != nil {
			log.Printf("error finding  menu by id err: %v", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  err.Error(),
			})
			return
		}
	} else if name != "" {
		response, err = client.FetchByName(ctxt, &adminpb.AMenuBYName{Name: name})
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
