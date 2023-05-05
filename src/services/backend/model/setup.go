package model

import (
	"backend/algorithm"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var questions []string
var answers []string

type ChatGPT struct {
	Id         int64  `gorm:"primaryKey" json:"id"`
	Pertanyaan string `gorm:"type:text" json:"pertanyaan"`
	Jawaban    string `gorm:"type:text" json:"jawaban"`
}

type History1 struct {
	Id         int64  `gorm:"primaryKey" json:"id"`
	Pertanyaan string `gorm:"type:text" json:"pertanyaan"`
	Jawaban    string `gorm:"type:text" json:"jawaban"`
}

type History2 struct {
	Id         int64  `gorm:"primaryKey" json:"id"`
	Pertanyaan string `gorm:"type:text" json:"pertanyaan"`
	Jawaban    string `gorm:"type:text" json:"jawaban"`
}

type History3 struct {
	Id         int64  `gorm:"primaryKey" json:"id"`
	Pertanyaan string `gorm:"type:text" json:"pertanyaan"`
	Jawaban    string `gorm:"type:text" json:"jawaban"`
}

type History4 struct {
	Id         int64  `gorm:"primaryKey" json:"id"`
	Pertanyaan string `gorm:"type:text" json:"pertanyaan"`
	Jawaban    string `gorm:"type:text" json:"jawaban"`
}

type History5 struct {
	Id         int64  `gorm:"primaryKey" json:"id"`
	Pertanyaan string `gorm:"type:text" json:"pertanyaan"`
	Jawaban    string `gorm:"type:text" json:"jawaban"`
}


func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/gpt"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&ChatGPT{})

	DB = database
	// get all Pertanyaan inside chat_gpt table and store it in questions
	DB.Model(&ChatGPT{}).Select("pertanyaan").Find(&questions)
	// get all Jawaban inside chat_gpt table and store it in answers
	DB.Model(&ChatGPT{}).Select("jawaban").Find(&answers)
}

func Index(c *gin.Context) {
	MigrateToGPT()
	var GPTs []ChatGPT

	DB.Find(&GPTs)
	c.JSON(http.StatusOK, gin.H{"GPTs": GPTs})

}

func Show(c *gin.Context) {
	MigrateToGPT()
	pertanyaan := c.Param("pertanyaan")
	pertanyaan = pertanyaan[1:]
	fmt.Println(pertanyaan)
	response := algorithm.ParseInput(pertanyaan, questions, answers)
	fmt.Println(response)

	c.JSON(http.StatusOK, gin.H{"answer": response})

}

func Create(c *gin.Context) {
	MigrateToGPT()
	var gpt ChatGPT

	if err := c.ShouldBindJSON(&gpt); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	DB.Create(&gpt)
	c.JSON(http.StatusOK, gin.H{"gpt": gpt})
}

func Update(c *gin.Context) {
	MigrateToGPT()
	var gpt ChatGPT
	id := c.Param("id")

	if err := c.ShouldBindJSON(&gpt); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if DB.Model(&gpt).Where("id = ?", id).Updates(&gpt).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate gpt"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})

}

func Delete(c *gin.Context) {
	MigrateToGPT()
	var gpt ChatGPT

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if DB.Delete(&gpt, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus gpt"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}

func ShowHistory(c *gin.Context) {
	id := c.Param("id")
	var history []History1
	MigrateToHistory(id)
	DB.Find(&history)
	c.JSON(http.StatusOK, gin.H{"history": history})
}


func MigrateToHistory(id string) {
	if (id == "1") {
		DB.AutoMigrate(&History1{})
	} else if (id == "2") {
		DB.AutoMigrate(&History2{})
	} else if (id == "3") {
		DB.AutoMigrate(&History3{})
	} else if (id == "4") {
		DB.AutoMigrate(&History4{})
	} else if (id == "5") {
		DB.AutoMigrate(&History5{})
	}
}

func MigrateToGPT() {
	DB.AutoMigrate(&ChatGPT{})
}

func AddHistory(c *gin.Context) {
	id := c.Param("id")
	MigrateToHistory(id)
	var history History1
	if err := c.ShouldBindJSON(&history); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	DB.Create(&history)
	c.JSON(http.StatusOK, gin.H{"history": history})
}

func DeleteHistory(c *gin.Context) {
	id := c.Param("id")
	MigrateToHistory(id)
	// clear all data inside history with id table
	DB.Exec("DELETE FROM history" + id)
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}