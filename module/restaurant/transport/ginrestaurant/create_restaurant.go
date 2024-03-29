package ginrestaurant

import (
	restaurantbiz "learn_go/module/restaurant/biz"
	restaurantmodel "learn_go/module/restaurant/model"
	restaurantstorage "learn_go/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := restaurantstorage.NewSQLStore(db)

		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.Create(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
