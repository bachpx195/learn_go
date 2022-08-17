package restaurantstorage

import (
	"context"
	restaurantmodel "learn_go/module/restaurant/model"
)

func (s *sqlStore) FindDataWithCondition(
	context context.Context,
	condition map[string]interface{},
	moreKey ...string,
) (*restaurantmodel.Restaurant, error) {
	var data restaurantmodel.Restaurant

	if err := s.db.Create(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
