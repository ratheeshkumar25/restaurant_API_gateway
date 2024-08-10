// package middleware

// import (
// 	//"fmt"
// 	"net/http"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt"
// )

// // ClearCache to handle the session
// func ClearCache() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
// 		c.Header("Pragma", "no-cache")
// 		c.Header("Expires", "0")

// 		c.Next()
// 	}
// }

// // Authorization to handle  authorization through middleware
// func Authorization(role string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		tokenString := c.GetHeader("Authorization")

// 		if tokenString == "" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is missing"})
// 			c.Abort()
// 			return
// 		}

// 		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

// 		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
// 			return []byte("q3e67yajhsdb4"), nil
// 		})
// 		if err != nil || !token.Valid {
// 			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid token"})
// 			c.Abort()
// 			return
// 		}

// 		claims, ok := token.Claims.(jwt.MapClaims)
// 		if !ok {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
// 			c.Abort()
// 			return
// 		}

// 		username, ok := claims["username"].(string)
// 		if !ok {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Email missing in claims"})
// 			c.Abort()
// 			return
// 		}

// 		userIDf, ok := claims["userid"].(float64)
// 		if !ok {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID missing in claims"})
// 			c.Abort()
// 			return
// 		}

// 		ClaimRole, ok := claims["role"].(string)
// 		if !ok {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Role missing in claims"})
// 			c.Abort()
// 			return
// 		}
// 		if role != ClaimRole {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Don't have permissions to access"})
// 			c.Abort()
// 			return
// 		}

// 		userID := uint64(userIDf)

//			c.Set("username", username)
//			c.Set("user_id", userID)
//			c.Next()
//		}
//	}
package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// ClearCache to handle the session
func ClearCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")

		c.Next()
	}
}

// Authorization to handle authorization through middleware
func Authorization(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is missing"})
			c.Abort()
			return
		}

		tokenString = strings.TrimSpace(strings.Replace(tokenString, "Bearer ", "", 1))

		// token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// 	if _, ok := t.Method.(*jwt.SigningMethodHS256); !ok {
		// 		return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		// 	}
		// 	return []byte("q3e67yajhsdb4"), nil
		// })
		// if err != nil || !token.Valid {
		// 	fmt.Printf("Token Parsing Error: %v\n", err)
		// 	c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid token"})
		// 	c.Abort()
		// 	return
		// }

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte("q3e67yajhsdb4"), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			fmt.Println("Invalid token claims")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		fmt.Printf("Token Claims: %+v\n", claims)

		username, ok := claims["username"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Username missing in claims"})
			c.Abort()
			return
		}

		// Handle userid as either a float64 or string
		// var userID uint
		// fmt.Printf("datatype %T", claims["userid"])
		// switch v := claims["userid"].(type) {
		// case uint:
		// 	userID = v
		// case string:
		// 	fmt.Sscanf(v, "%d", &userID)
		// default:
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID missing or invalid in claims"})
		// 	c.Abort()
		// 	return
		// }

		ClaimRole, ok := claims["role"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Role missing in claims"})
			c.Abort()
			return
		}

		if role != ClaimRole {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Don't have permissions to access"})
			c.Abort()
			return
		}

		c.Set("username", username)
		// c.Set("user_id", userID)
		c.Next()
	}
}
