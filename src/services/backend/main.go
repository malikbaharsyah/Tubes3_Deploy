package main

import (
	"backend/model"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	model.ConnectDatabase()

	r.GET("/api/products", model.Index)
	r.GET("/api/product/:id", model.Show)
	r.POST("/api/product", model.Create)
	r.PUT("/api/product/:id", model.Update)
	r.DELETE("/api/product", model.Delete)

	// Serve the HTML form
	r.GET("/form", func(c *gin.Context) {
		c.File("tes.html")
	})

	//add new data nama product dan deskripsi laptop
	// laptop := model.Product{NamaProduct: "Laptop", Deskripsi: "Laptop adalah komputer jinjing"}


	r.Run()
}