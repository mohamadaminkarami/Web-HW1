package main

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"net/http"
	"strings"

	redisInterface "github.com/mohamadaminkarami/Web-HW1/redis"
	"github.com/mohamadaminkarami/Web-HW1/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func main() {
	redisInterface.PingRedis()

	router := gin.Default()
	router.GET("sha", getHash)
	router.POST("sha", setHash)

	err := router.Run(utils.GetEnv("SERVER_ADDR", "0.0.0.0:8080"))
	if err != nil {
		log.Panic(err)
	}
}

func getHash(c *gin.Context) {
	encoded := c.Query("encoded")

	if encoded == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": []string{"encoded param is required"},
		})
		return
	}

	value, err := redisInterface.GetValueRedis(encoded)

	switch {
	case err == redis.Nil:
		c.JSON(http.StatusNotFound, gin.H{
			"errors": []string{"sha256 hash not found!"},
		})
		return
	case err != nil:
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"raw_string": value,
	})
}

type HashString struct {
	RawString string `form:"raw_string" json:"raw_string" xml:"raw_string"  binding:"required"`
}

func setHash(c *gin.Context) {
	var json HashString
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": []string{"raw_staring must be at least 8 characters"},
		})
		return
	}
	raw_string := json.RawString

	if len(raw_string) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": []string{"raw_staring must be at least 8 characters"},
		})
		return
	}

	input := strings.NewReader(raw_string)

	hash := sha256.New()
	if _, err := io.Copy(hash, input); err != nil {
		panic(err)
	}
	sum := hash.Sum(nil)
	encoded := hex.EncodeToString(sum)

	err := redisInterface.SetValueRedis(encoded, raw_string)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"encoded": encoded,
	})
}
