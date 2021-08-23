package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"api/api/api_errors"
	"api/app"
	"api/config"
)

func RequestRate() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.ClientIP()
		redis := app.RedisDB
		if redis.LLen(key).Val() >= config.GetApiRequestRate().Rate {
			c.JSON(http.StatusOK, api_errors.ErrorTooManyRequests())
			c.Abort()
			return
		}

		rateDuringTime := config.GetApiRequestRate().DuringTime * time.Second
		if v := redis.Exists(key).Val(); v == 0 {
			pipe := redis.TxPipeline()
			pipe.RPush(key, key)
			pipe.Expire(key, rateDuringTime)
			if _, err := pipe.Exec(); err != nil {
				app.ZapLog.Error("api rate", fmt.Sprintf("request rate redis pipe error:%v", err))
			}
		} else {
			redis.RPushX(key, key)
		}

		c.Next()
	}
}
