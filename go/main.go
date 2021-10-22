package main

import (
	"crypto/sha256"
	"encoding/hex"
	redisInterface "github.com/mohamadaminkarami/Web-HW1/redis"
	"github.com/mohamadaminkarami/Web-HW1/utils"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func main() {
	utils.LoadEnv()

	router := gin.Default()
	router.GET("/sha", getHash)
	router.POST("/sha", setHash)

	err := router.Run(utils.GetEnv("SERVER_ADDR", "localhost:8080"))
	if err != nil {
		panic(err)
	}
}

func getHash(c *gin.Context) {
	encoded := c.Query("encoded")

	if encoded == "" {
		c.JSON(http.StatusBadRequest, map[string]string{
			"detail": "encoded param is required",
		})
		return
	}

	value, err := redisInterface.GetValueRedis(encoded)

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
		"detail": value,
	})
}

func setHash(c *gin.Context) {
	decoded := c.PostForm("decoded")

	if len(decoded) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": "input must be at least 8 characters",
		})
		return
	}

	input := strings.NewReader(decoded)

	hash := sha256.New()
	if _, err := io.Copy(hash, input); err != nil {
		panic(err)
	}
	sum := hash.Sum(nil)
	encoded := hex.EncodeToString(sum)

	err := redisInterface.SetValueRedis(encoded, decoded)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"detail": encoded,
	})
}
