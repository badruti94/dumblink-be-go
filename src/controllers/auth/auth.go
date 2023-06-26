package auth

import (
	"dumblink-be-go/src/db"
	"dumblink-be-go/src/models/user"
	"dumblink-be-go/src/utils"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user user.User

	c.Bind(&user)
	user.CreatedAt = utils.GetTime()
	user.UpdatedAt = utils.GetTime()
	passwordByte, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		panic(err)
	}
	user.Password = string(passwordByte)

	// db := d
	db := db.Connect()
	db.Create(&user)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.Id,
	})
	hmacSampleSecret := []byte("my_secret_key")
	tokenString, _ := token.SignedString(hmacSampleSecret)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"user": gin.H{
				"name":  user.Name,
				"email": user.Email,
			},
			"token": tokenString,
		},
	})

}

func Login(c *gin.Context) {
	var user user.User

	c.Bind(&user)
	passwordString := user.Password

	db := db.Connect()
	db.Where("email = ?", user.Email).First(&user)

	if user.Id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "user not found",
		})
		return
	}

	fmt.Println(passwordString)
	fmt.Println(user.Password)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordString))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "credential is invalid",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.Id,
	})
	hmacSampleSecret := []byte("my_secret_key")
	tokenString, err := token.SignedString(hmacSampleSecret)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"user": gin.H{
				"name":  user.Name,
				"email": user.Email,
			},
			"token": tokenString,
		},
	})

}
