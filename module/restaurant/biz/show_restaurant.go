package restaurantbiz

import "context"

type ShowRestaurantStore interface {
	FindDataWithCondition(context context.Context)
}
