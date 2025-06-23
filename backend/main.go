package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sharukh010/url-shortner/route"
)


func main(){
	router := gin.New()
	route.RegisterRoutes(router)

	if err := http.ListenAndServe(":8000",router); err != nil {
		log.Panic(err.Error())
	}
}