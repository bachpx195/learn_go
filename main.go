package main

import (
	"learn_go/module/restaurant/transport/ginrestaurant"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:quachtinh95@tcp(127.0.0.1:3306)/learn_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	log.Println(err)

	r := gin.Default()

	// POST /restaurants
	v1 := r.Group("/v1")
	v1.POST("/restaurants", ginrestaurant.CreateRestaurant(db))

	r.Run()
}
