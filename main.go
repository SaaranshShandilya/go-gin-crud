package main

import (
	"fmt"
	"log"

	"github.com/SaaranshShandilya/go-crud/controllers"
	"github.com/SaaranshShandilya/go-crud/initialise"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)


var DB * gorm.DB

func init(){
	err:=godotenv.Load()
	if err!=nil{
		log.Fatal(err)
	}

	DB = initialise.ConnectToDb()

    // Provide db variable to controllers
    
}




func main() {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		fmt.Println()
        c.Set("db", DB)
        c.Next()
    })

	r.POST("/create", controllers.PostsCreate)
	r.POST("/getPost", controllers.PostsIndex)
	r.GET("/getAll", controllers.GetAll)
	r.Run() // listen and serve on 0.0.0.0:8080
}