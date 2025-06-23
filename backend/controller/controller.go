package controller

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sharukh010/url-shortner/db"
	"github.com/sharukh010/url-shortner/repository"
	"github.com/sharukh010/url-shortner/service"
)
var baseURL string

func init(){
	if err := godotenv.Load(); err != nil {
		log.Panic("Error occured while loading .env file")
	}
	baseURL = os.Getenv("BASE_URL")
}

func TestAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"message":"API working..."})
	}
}

func CreateLink() gin.HandlerFunc{

	return func(c *gin.Context){
		var newLink db.Url
		if err := c.BindJSON(&newLink); err != nil {
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return 
		}
		if newLink.Validity == 0 {
			newLink.Validity = 30
		}
		newLink.ExpireAt = time.Now().Add(time.Minute * time.Duration(newLink.Validity))

		if newLink.ShortCode == "" {
			newLink.ShortCode = service.GenerateShortCode()
		}
		newLink.ShortLink = baseURL+newLink.ShortCode;
		if err := repository.CreateLink(&newLink); err != nil {
			log.Println("Error while creating new link:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create short link"})
			return
		}
		c.JSON(http.StatusCreated,gin.H{"shortLink":newLink.ShortLink,"expiry":newLink.ExpireAt})
	}


}
