package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Restaurant struct {
	Id     int             `json:"id" gorm:"column:id;"`
	Name   string          `json:"name" gorm:"column:name;"`
	Addr   string          `json:"addr" gorm:"column:addr;"`
	Logo   json.RawMessage `json:"logo" gorm:"column:logo;"`
	Status int             `json:"status" gorm:"column:status;"`
}

func (Restaurant) TableName() string { return "restaurants" }

func main() {
	fmt.Println("bach")
	dsn := "root:quachtinh95@tcp(127.0.0.1:3306)/learn_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	log.Println(err)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.Run()

	newRestaurant := Restaurant{Name: "Tani", Addr: "Da nang", Status: 1}

	if err := db.Create(&newRestaurant).Error; err != nil {
		log.Println(err)
	}

	var myRestaurant Restaurant

	if err := db.Where("id = ?", newRestaurant.Id).First(&myRestaurant).Error; err != nil {
		log.Println(err)
	}

	log.Println(myRestaurant)

	myRestaurant.Name = "bach"
	if err := db.Where("id = ?", 2).Updates(&myRestaurant).Error; err != nil {
		log.Println(err)
	}

	log.Println(myRestaurant)
}
