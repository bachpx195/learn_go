package restaurantmodel

type Restaurant struct {
	Id     int    `json:"id" gorm:"column:id;"`
	Name   string `json:"name" gorm:"column:name;"`
	Addr   string `json:"addr" gorm:"column:addr;"`
	Status int    `json:"status" gorm:"column:status"`
}

func (Restaurant) TableName() string { return "restaurants" }

type RestaurantCreate struct {
	Id     int    `json:"id" gorm:"column:id;"`
	Name   string `form:"name" json:"name" gorm:"column:name;"`
	Addr   string `form:"addr" json:"addr" gorm:"column:addr;"`
	Status int    `form:"status" json:"status" gorm:"column:status"`
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"address" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }
