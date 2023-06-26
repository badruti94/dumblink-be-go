package link

import (
	"dumblink-be-go/src/db"
	"dumblink-be-go/src/models/link"
	"dumblink-be-go/src/utils"
	"net/http"
	"os"
	"strconv"
	"time"

	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type UserLinkResult struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
type LinkResult struct {
	Id             int    `json:"id" form:"id"`
	Uniqid         string `json:"uniqid" form:"uniqid"`
	Type           string `json:"type" form:"type"`
	Title          string `json:"title" form:"title"`
	Description    string `json:"description" form:"description"`
	Vlog           string `json:"vlog" form:"vlog"`
	Galery         string `json:"galery" form:"galery"`
	Contact        string `json:"contact" form:"contact"`
	About          string `json:"about" form:"about"`
	Facebook       string `json:"facebook" form:"facebook"`
	Instagram      string `json:"instagram" form:"instagram"`
	Twitter        string `json:"twitter" form:"twitter"`
	Youtube        string `json:"youtube" form:"youtube"`
	Whatsapp       string `json:"whatsapp" form:"whatsapp"`
	Phone          string `json:"phone" form:"phone"`
	Web            string `json:"web" form:"web"`
	Lazada         string `json:"lazada" form:"lazada"`
	Photo          string `json:"photo" form:"photo"`
	View           int    `json:"view" form:"view"`
	UserId         int    `json:"userId" gorm:"column:userId"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	UserLinkResult `json:"user"`
}

func uploadFile(c *gin.Context) (string, error) {
	file, err := c.FormFile("photo")
	if err != nil {
		return "", err
	}
	fileContent, err := file.Open()
	if err != nil {
		return "", err
	}

	fileContent.Close()

	cloud_name := os.Getenv("CLOUDINARY_NAME")
	api_key := os.Getenv("CLOUDINARY_API_KEY")
	api_secret := os.Getenv("CLOUDINARY_API_SECRET")

	cld, _ := cloudinary.NewFromParams(cloud_name, api_key, api_secret)

	var ctx = context.Background()

	randomNumber := strconv.FormatInt(time.Now().UnixNano(), 24)

	resp, err := cld.Upload.Upload(ctx, fileContent, uploader.UploadParams{PublicID: randomNumber})
	if err != nil {
		return "", err
	}

	return resp.SecureURL, nil

}

func AddLink(c *gin.Context) {
	db := db.Connect()

	var link link.Link

	c.ShouldBind(&link)
	link.CreatedAt = utils.GetTime()
	link.UpdatedAt = utils.GetTime()

	link.UserId = c.MustGet("userId").(int)
	guid := xid.New()
	link.Uniqid = guid.String()
	photoLink, err := uploadFile(c)
	if err != nil {
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to upload file",
			})
			return
		}
	}
	link.Photo = photoLink

	db.Create(&link)

	c.JSON(http.StatusOK, gin.H{
		"message": "Link created successfully",
	})
}

func UpdateLink(c *gin.Context) {
	db := db.Connect()

	var link link.Link

	db.First(&link, c.Param("id"))
	c.ShouldBind(&link)
	link.UpdatedAt = utils.GetTime()

	_, err := c.FormFile("photo")

	fileExist := true
	if err != nil {
		if err == http.ErrMissingFile {
			fileExist = false
		}
	}
	if fileExist {
		photoLink, err := uploadFile(c)
		if err != nil {
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "failed to upload file",
				})
				return
			}
		}
		link.Photo = photoLink
	}

	db.Save(&link)

	c.JSON(http.StatusOK, gin.H{
		"message": "Link updated successfully",
	})
}

func GetLink(c *gin.Context) {
	db := db.Connect()

	Link := link.Link{}
	var link LinkResult

	db.Model(&Link).Select("links.*, users.email, users.name").Joins("join users on links.\"userId\" = users.id").First(&link, c.Param("id"))
	link.UserLinkResult.Name = link.Name
	link.UserLinkResult.Email = link.Email

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"link": link,
		},
	})
}

func GetLinkByUniqid(c *gin.Context) {
	db := db.Connect()

	Link := link.Link{}
	var link LinkResult

	db.Model(&Link).Select("links.*, users.email, users.name").Joins("join users on links.\"userId\" = users.id").Where("uniqid = ?", c.Param("uniqid")).First(&link)
	link.UserLinkResult.Name = link.Name
	link.UserLinkResult.Email = link.Email

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"link": link,
		},
	})
}

func GetLinks(c *gin.Context) {
	db := db.Connect()

	var links []link.Link

	userId := c.MustGet("userId").(int)
	db.Where("\"userId\" = ?", userId).Find(&links)

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"links": links,
		},
	})
}

func DeleteLink(c *gin.Context) {
	db := db.Connect()

	var link link.Link

	link.Id, _ = strconv.Atoi(c.Param("id"))
	db.Delete(&link, link)

	c.JSON(http.StatusOK, gin.H{
		"message": "Link deleted successfully",
	})
}

func CountLInk(c *gin.Context) {
	db := db.Connect()

	var link link.Link
	db.First(&link, c.Param("id"))
	link.View = link.View + 1
	db.Save(&link)

	c.JSON(http.StatusOK, gin.H{
		"message": "Add 1 view to link",
	})
}
