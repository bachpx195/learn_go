package ginrestaurant

import (
	restaurantmodel "learn_go/module/restaurant/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		log.Println("test")
		log.Println(data)
		//body, _ := ioutil.ReadAll(c.Request.Body)
		//log.Println(string(body))

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		//store := restaurantstorage.NewSQLStore(db)
		//
		//biz := restaurantbiz.NewCreateRestaurantBiz(store)
		//
		//if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
		//	c.JSON(http.StatusBadRequest, gin.H{
		//		"error": err.Error(),
		//	})
		//}
		//
		//c.JSON(http.StatusOK, gin.H{
		//	"data": data,
		//})
	}
}
