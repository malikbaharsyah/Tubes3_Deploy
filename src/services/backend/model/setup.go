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
	database, err := gorm.Open(mysql.Open("root:opEiqwNZ5ljqbstec2Pu@tcp(containers-us-west-72.railway.app:6044)/railway"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&ChatGPT{})

	DB = database
	// get all Pertanyaan inside chat_gpt table and store it in questions
	DB.Model(&ChatGPT{}).Select("pertanyaan").Find(&questions)
}

func Index(c *gin.Context) {
	MigrateToGPT()
	var GPTs []ChatGPT

	DB.Find(&GPTs)
	c.JSON(http.StatusOK, gin.H{"GPTs": GPTs})

}

func Show(c *gin.Context) {
	MigrateToGPT()
	getQuestions()
	pertanyaan := c.Param("pertanyaan")
	pertanyaan = pertanyaan[1:]
	fmt.Println(pertanyaan)
	param := c.Param("param")
	// set param to int
	var param_int int
	if param == "0" {
		param_int = 0
	} else {
		param_int = 1
	}
	response := algorithm.ParseInput(pertanyaan, questions, param_int)
	fmt.Println(response)
	if response[0] == "kalender" || response[0] == "kalkulator" {
		answer := []string{}
		answer = append(answer, response[1])
		c.JSON(http.StatusOK, gin.H{"answer": answer})
	} else if response[0] == "tambah" {
		answer := []string{}
		var gpt ChatGPT
		gpt.Pertanyaan = response[1]
		gpt.Jawaban = response[2]
		// if pertanyaan exists in database, update
		// else create new pertanyaan
		if DB.Where("pertanyaan = ?", response[1]).First(&gpt).RowsAffected != 0 {
			DB.Model(&gpt).Where("pertanyaan = ?", response[1]).Update("jawaban", response[2])
			answer = append(answer, "Pertanyaan "+response[1]+" sudah ada ! jawaban diupdate ke "+response[2])
			c.JSON(http.StatusOK, gin.H{"answer": answer})
		} else {
			DB.Create(&gpt)
			answer = append(answer, "Pertanyaan "+response[1]+" telah ditambahkan")
			c.JSON(http.StatusOK, gin.H{"answer": answer})
		}
	} else if response[0] == "hapus" {
		answer := []string{}
		var gpt ChatGPT
		// if pertanyaan exists, delete
		// else c.JSON(http.StatusOK, gin.H{"answer": "Tidak ada pertanyaan " + response[1] + " pada database!"})
		if DB.Where("pertanyaan = ?", response[1]).First(&gpt).RowsAffected != 0 {
			DB.Delete(&gpt)
			answer = append(answer, "Pertanyaan "+response[1]+" telah dihapus")
			c.JSON(http.StatusOK, gin.H{"answer": answer})
		} else {
			answer = append(answer, "Tidak ada pertanyaan "+response[1]+" pada database!")
			c.JSON(http.StatusOK, gin.H{"answer": answer})
		}
	} else if response[0] == "jawaban" {
		// get jawaban from database
		answer := []string{}
		var gpt ChatGPT
		if DB.Where("pertanyaan = ?", response[1]).First(&gpt).RowsAffected != 0 {
			answer = append(answer, gpt.Jawaban)
		} else {
			answer = append(answer, "Tidak ada pertanyaan "+response[1]+" pada database!")
		}

		c.JSON(http.StatusOK, gin.H{"answer": answer})
	} else if response[0] == "rekomendasi" {
		answer := []string{}
		answer = append(answer, response[1])
		c.JSON(http.StatusOK, gin.H{"answer": answer})
	}

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
	MigrateToHistory(id)
	switch id {
	case "1":
		var history []History1
		DB.Find(&history)
		c.JSON(http.StatusOK, gin.H{"history": history})
	case "2":
		var history []History2
		DB.Find(&history)
		c.JSON(http.StatusOK, gin.H{"history": history})
	case "3":
		var history []History3
		DB.Find(&history)
		c.JSON(http.StatusOK, gin.H{"history": history})
	case "4":
		var history []History4
		DB.Find(&history)
		c.JSON(http.StatusOK, gin.H{"history": history})
	case "5":
		var history []History5
		DB.Find(&history)
		c.JSON(http.StatusOK, gin.H{"history": history})
	}
}

func MigrateToHistory(id string) {
	if id == "1" {
		DB.AutoMigrate(&History1{})
	} else if id == "2" {
		DB.AutoMigrate(&History2{})
	} else if id == "3" {
		DB.AutoMigrate(&History3{})
	} else if id == "4" {
		DB.AutoMigrate(&History4{})
	} else if id == "5" {
		DB.AutoMigrate(&History5{})
	}
}

func MigrateToGPT() {
	DB.AutoMigrate(&ChatGPT{})
}

func AddHistory(c *gin.Context) {
	id := c.Param("id")
	MigrateToHistory(id)
	switch id {
	case "1":
		var history History1
		if err := c.ShouldBindJSON(&history); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		DB.Create(&history)
		c.JSON(http.StatusOK, gin.H{"history": history})
	case "2":
		var history History2
		if err := c.ShouldBindJSON(&history); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		DB.Create(&history)
		c.JSON(http.StatusOK, gin.H{"history": history})
	case "3":
		var history History3
		if err := c.ShouldBindJSON(&history); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		DB.Create(&history)
		c.JSON(http.StatusOK, gin.H{"history": history})
	case "4":
		var history History4
		if err := c.ShouldBindJSON(&history); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		DB.Create(&history)
		c.JSON(http.StatusOK, gin.H{"history": history})
	case "5":
		var history History5
		if err := c.ShouldBindJSON(&history); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		DB.Create(&history)
		c.JSON(http.StatusOK, gin.H{"history": history})
	}
}

func DeleteHistory(c *gin.Context) {
	id := c.Param("id")
	MigrateToHistory(id)
	DB.Exec("DELETE FROM history" + id)
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}

func getQuestions() {
	MigrateToGPT()
	DB.Model(&ChatGPT{}).Select("pertanyaan").Find(&questions)
}
