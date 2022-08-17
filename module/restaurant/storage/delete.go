package restaurantstorage

import (
	"context"
	restaurantmodel "learn_go/module/restaurant/model"
)

func (s *sqlStore) Delete(
	context context.Context,
	id int,
) error {
	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).Error; err != nil {
		return err
	}

	return nil
}
