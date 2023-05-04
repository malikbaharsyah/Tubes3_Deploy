package model

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type ChatGPT struct {
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
}

func Index(c *gin.Context) {

	var GPTs []ChatGPT

	DB.Find(&GPTs)
	c.JSON(http.StatusOK, gin.H{"GPTs": GPTs})

}

func Show(c *gin.Context) {
	var gpt ChatGPT
	id := c.Param("id")

	if err := DB.First(&gpt, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"gpt": gpt})
}

func Create(c *gin.Context) {

	var gpt ChatGPT

	if err := c.ShouldBindJSON(&gpt); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	DB.Create(&gpt)
	c.JSON(http.StatusOK, gin.H{"gpt": gpt})
}

func Update(c *gin.Context) {
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
