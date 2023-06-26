package profile

import (
	"dumblink-be-go/src/db"
	"dumblink-be-go/src/models/user"
	"dumblink-be-go/src/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	db := db.Connect()

	var user user.User

	fmt.Println(c.MustGet("userId"))
	user.Id = c.MustGet("userId").(int)

	db.Select("email", "name").First(&user, user.Id)

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"profile": user,
		},
	})

}

func UpdateProfile(c *gin.Context) {
	db := db.Connect()

	var user user.User

	user.Id = c.MustGet("userId").(int)
	db.First(&user)
	c.Bind(&user)
	user.UpdatedAt = utils.GetTime()
	db.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile updated successfully",
	})
}

func DeleteProfile(c *gin.Context) {
	db := db.Connect()

	var user user.User
	user.Id = c.MustGet("userId").(int)
	db.Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile deleted successfully",
	})
}
