package app

import (
	"log"

	"github.com/gin-gonic/gin"
)

var (
	r = gin.Default()
)

func StartApplication() {
	mapUrls()
	log.Printf("Creating your time card application")

	r.Run(":8080")

}
