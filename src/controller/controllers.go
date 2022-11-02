package controller

import (
	"log"

	"example.com/main/src/service"
	"github.com/gin-gonic/gin"
)

var sv service.Service

func Find(c *gin.Context) {
	result, err := sv.Find()

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		c.JSON(400, gin.H{
			"status": "failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"data":   result,
		"status": "success",
	})
}

func Post(c *gin.Context) {

}

func FindOne(c *gin.Context) {
	// id := c.Param("id")


}
