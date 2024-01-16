package controllers

import (
	"fmt"
	"github.com/SaaranshShandilya/go-crud/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	
)

func PostsCreate(c *gin.Context){

	var body struct{
		Body string
		Title string

	}

	c.Bind(&body)

	db := c.MustGet("db").(*gorm.DB)
	user := models.Post{Title: body.Title, Body: body.Body}
	result:=db.Create(&user)
	fmt.Println("the result is %v",result)

	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}

	c.JSON(200,gin.H{
		"post":user,
	})
}


func PostsIndex (c *gin.Context){
	var body struct{
		Id int
	}
	c.Bind(&body)
	db := c.MustGet("db").(*gorm.DB)
	fmt.Println(body.Id)
	d:=models.Post{}
	result:=db.First(&d, body.Id)

	if result.Error != nil{
		fmt.Println(result.Error)
		return
	}

	c.JSON(200,gin.H{
		"Single Post":d,
	})




}

func GetAll (c *gin.Context){
	
	db := c.MustGet("db").(*gorm.DB)
	d := []models.Post{}
	result:=db.Find(&d)

	if result.Error != nil{
		fmt.Println(result.Error)
		return
	}

	c.JSON(200,gin.H{
		"All Posts":d,
	})




}