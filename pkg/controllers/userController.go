package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/Joybaruah/ssl-tracker/pkg/inits"
	models "github.com/Joybaruah/ssl-tracker/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {

	var body struct {
		Name     string
		Email    string
		Password string
	}

	if ctx.BindJSON(&body) != nil {
		ctx.JSON(400, gin.H{"error": "bad request"})
		return
	}

	// Check if Email Exists
	count := int64(0)
	emailCheck := inits.DB.Model(&models.User{}).Where("email = ?", body.Email).Count(&count)

	if emailCheck.Error != nil {
		ctx.JSON(500, gin.H{"error": emailCheck.Error})
		return
	}

	exists := count > 0
	if exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User with Email " + body.Email + " already exists"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err})
		return
	}

	user := models.User{Name: body.Name, Email: body.Email, Password: string(hash)}
	result := inits.DB.Create(&user)

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, user)

}

func Login(ctx *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if ctx.BindJSON(&body) != nil {
		ctx.JSON(400, gin.H{"error": "bad request"})
		return
	}

	var user models.User

	result := inits.DB.Where("email = ?", body.Email).First(&user)

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": "User with Email " + body.Email + " not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect Password"})
		return
	}

	// generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		ctx.JSON(500, gin.H{"error": "error signing token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})

	// ctx.SetSameSite(http.SameSiteLaxMode)
	// ctx.SetCookie("Authorization", tokenString, 3600*24*30, "", "localhost", false, true)
}

func GetUsers(ctx *gin.Context) {

	var user []models.User
	inits.DB.Find(&user)

	ctx.JSON(200, user)

}
