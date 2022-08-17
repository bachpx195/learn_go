package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	restaurantbiz "learn_go/module/restaurant/biz"
	restaurantstorage "learn_go/module/restaurant/storage"
	"net/http"
	"strconv"
)

func ShowRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": 1,
		})
	}
}
