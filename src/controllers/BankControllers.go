package controllers

import (
	"backend/data"
	"backend/models"
	"github.com/gin-gonic/gin"
	"time"
)

func PostBank(c *gin.Context) {
	c.Bind(&models.Bank{})

	var post models.Bank
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()
	result := data.ConnectDatabase().Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetAll(c *gin.Context) {
	var banks []models.Bank
	db := data.ConnectDatabase()
	db.Find(&banks)

	c.JSON(200, gin.H{
		"banks": banks,
	})
}
