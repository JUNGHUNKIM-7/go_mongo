package main

import (
	"context"
	"log"

	"example.com/main/src/controller"
	"example.com/main/src/repository"
	"example.com/main/src/types"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = repository.Connnect()
	if err != nil {
		panic("connection failed")
	}
}

func main() {
	r := gin.Default()
	cli := types.MongoClient

	userGroup := r.Group("/pokemons")
	{
		userGroup.GET("/", controller.Find)
		userGroup.GET("/:id", controller.FindOne)
		userGroup.POST("/", controller.Post)
	}

	defer func() {
		if err := cli.C.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	r.Run(":3000")
}
