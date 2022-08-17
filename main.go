package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"learn_go/module/restaurant/transport/ginrestaurant"
)

func main() {

	dsn := "root:quachtinh95@tcp(127.0.0.1:3306)/learn_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	r := gin.Default()

	// POST /restaurants
	v1 := r.Group("/v1")
	v1.POST("/restaurants", ginrestaurant.CreateRestaurant(db))
	v1.GET("/:id", ginrestaurant.DeleteRestaurant(db))
	v1.DELETE("/:id", ginrestaurant.DeleteRestaurant(db))

	r.Run()
}
