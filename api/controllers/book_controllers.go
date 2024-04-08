package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"goapi.railway.app/api/database"
	"goapi.railway.app/api/model"
)

func CreateBook(c *gin.Context) {
	var book model.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message" : err.Error()})
		return
	}

	database.DB.Create(&book)
	c.JSON(http.StatusOK, gin.H{"success": true, "message" : book})
}

func GetAll (c *gin.Context){
	var book []model.Book

	database.DB.Find(&book)
	c.JSON(http.StatusOK, gin.H{"success" : true , "message" : book})
}


func GetOne (c *gin.Context){
	var book []model.Book

	if err := database.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message" : "Record not Found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success" : true , "message" : book})
}


func UpdateBook (c *gin.Context){
	var book model.Book

	if err := database.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message" : "Record not Found!"})
		return
	}
	

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message" : err.Error()})
		return
	}

	database.DB.Save(&book)
	c.JSON(http.StatusOK, gin.H{"success": true, "message" : book})
}

func PatchBook (c *gin.Context){
	var book model.Book

	if err := database.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message" : "Record not Found!"})
		return
	}
	


	var updadeBook model.Book
	if err := c.BindJSON(&updadeBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message" : err.Error()})
		return
	}



	database.DB.Model(&book).Updates(&updadeBook)
	c.JSON(http.StatusOK, gin.H{"success": true, "message" : book})
}


func DeleteBook (c *gin.Context){
	var book model.Book

	if err := database.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message" : "Record not Found!"})
		return
	}
	

	database.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"success": true, "message" : book})
}