package restaurantbiz

import (
	"context"
	"errors"
	restaurantmodel "learn_go/module/restaurant/model"
	"log"
)

type CreateRestaurantStore interface {
	CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	log.Println("dkm")
	log.Println(data)
	log.Println(context)

	if data.Name == "" {
		return errors.New("name cannot be empty")
	}

	if err := biz.store.CreateRestaurant(context, data); err != nil {
		return err
	}

	return nil
}