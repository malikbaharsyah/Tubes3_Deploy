package main

import (
	// "fmt"
	// "backend/algorithm"
	"backend/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	model.ConnectDatabase()
	r.Use(cors.Default())

	r.GET("/api/gpts", model.Index)
	r.GET("/api/gpt/*pertanyaan", model.Show)
	r.GET("/api/history/:id", model.ShowHistory)
	r.POST("/api/gpt", model.Create)
	r.POST("/api/history/:id", model.AddHistory) 
	r.PUT("/api/gpt/:id", model.Update)
	r.DELETE("/api/gpt", model.Delete)
	r.DELETE("/api/history/:id", model.DeleteHistory)

	r.Run(":8000")

}
