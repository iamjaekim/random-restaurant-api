package handlers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/iamjaekim/random-restaurant-api/utils"
	"github.com/iamjaekim/random-restaurant-api/yelp"
)

func Index(c *gin.Context) {
	c.Redirect(http.StatusFound, "https://random-restaurant-zipcode.herokuapp.com/")
}

func GetRestaurants(c *gin.Context) {
	if !utils.ZipValidation(c.Param("zipCode")) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid zipcode"})
		return
	}
    zipCodeInt, err := strconv.Atoi(c.Param("zipCode"))
	businesses, err := yelp.SearchBusinesses(zipCodeInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Yelp API Error"})
		return
	}

	c.JSON(http.StatusOK, businesses)
}

func GetRestaurant(c *gin.Context) {
	storeID := c.Param("storeId")

	business, err := yelp.GetBusiness(storeID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Yelp API Error"})
		return
	}

	c.JSON(http.StatusOK, business)
}

