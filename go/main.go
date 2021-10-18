package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb = initRedis()

func main() {
	router := gin.Default()
	router.GET("/sha", getHash)
	router.POST("/sha", setHash)

	router.Run("localhost:8080")
}

func getHash(c *gin.Context) {
	key := c.Query("key")

	if key == "" {
		c.JSON(http.StatusBadRequest, map[string]string{
			"detail": "key param is required",
		})
		return
	}

	value, err := getValueRedis(key)

	switch {
	case err == redis.Nil:
		c.JSON(http.StatusNotFound, gin.H{
			"detail": "key not found",
		})
		return
	case err != nil:
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"value": value,
	})
}

func setHash(c *gin.Context) {
	key := c.PostForm("key")
	value := c.PostForm("value")

	if key == "" || value == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": "key & value are required",
		})
		return
	}

	err := setValueRedis(key, value)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"detail": key + " set to " + value,
	})
}
