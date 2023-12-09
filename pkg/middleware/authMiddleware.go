package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Joybaruah/ssl-tracker/pkg/inits"
	models "github.com/Joybaruah/ssl-tracker/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var openRoutes = []string{"/api/user/register", "/api/user/login"}

func isOpenRoute(target string) bool {
	// Create a comma-separated string of array elements
	joinedString := strings.Join(openRoutes, ",")

	// Check if the target string is a substring of the joined string
	return strings.Contains(joinedString, target)
}

func RequireApiAuth(ctx *gin.Context) {

	url := ctx.Request.URL.Path

	if isOpenRoute := isOpenRoute(url); isOpenRoute {
		ctx.Next()
	}

	tokenString := ctx.Request.Header.Get("Authorization")
	if tokenString == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.JSON(401, gin.H{"error": "unauthorized"})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var user models.User
		inits.DB.First(&user, int(claims["id"].(float64)))
		if user.Id == 0 {
			ctx.JSON(401, gin.H{"error": "unauthorized"})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("user", user)
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	ctx.Next()

}

// func isOpenRoute(target string, arr []string) bool {
// 	// Create a comma-separated string of array elements
// 	joinedString := strings.Join(arr, ",")
// 	logrus.Info(joinedString)

// 	// Check if the target string is a substring of the joined string
// 	return strings.Contains(joinedString, target)
// }

// openRoutes := []string{"/api/user/register", "/api/user/login"}
// url := ctx.Request.URL
// 	isApiRoute := isApiRoute(url.Path)

// 	if isApiRoute {

// 	}

// 	tokenString, err := ctx.Cookie("Authorization")

// 	logrus.Info(tokenString)

// 	if err != nil {
// 		ctx.JSON(401, gin.H{"error": "unauthorized"})
// 		ctx.AbortWithStatus(http.StatusUnauthorized)
// 		return
// 	}
// 	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}

// 		return []byte(os.Getenv("SECRET")), nil
// 	})

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		if float64(time.Now().Unix()) > claims["exp"].(float64) {
// 			ctx.JSON(401, gin.H{"error": "unauthorized"})
// 			ctx.AbortWithStatus(http.StatusUnauthorized)
// 			return
// 		}

// 		var user models.User
// 		inits.DB.First(&user, int(claims["id"].(float64)))
// 		if user.Id == 0 {
// 			ctx.JSON(401, gin.H{"error": "unauthorized"})
// 			ctx.AbortWithStatus(http.StatusUnauthorized)
// 			return
// 		}

// 		ctx.Set("user", user)
// 		fmt.Println(claims["foo"], claims["nbf"])
// 	} else {
// 		ctx.AbortWithStatus(http.StatusUnauthorized)
// 	}
// 	ctx.Next()
