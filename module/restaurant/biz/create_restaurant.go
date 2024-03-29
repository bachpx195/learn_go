package restaurantbiz

import (
	"context"
	"errors"
	restaurantmodel "learn_go/module/restaurant/model"
)

type CreateRestaurantStore interface {
	Create(context context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) Create(context context.Context, data *restaurantmodel.RestaurantCreate) error {

	if data.Name == "" {
		return errors.New("name cannot be empty")
	}
	if err := biz.store.Create(context, data); err != nil {
		return err
	}

	return nil
}
