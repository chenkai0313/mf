package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

var allowedHeaders = []string{
	"Accept-Language",
	"X-Request-Agent",
	"Wallet",
	"Network",
	"Content-Type",
	"Content-Length",
	"Accept-Encoding",
	"X-CSRF-Token",
	"Authorization",
	"accept",
	"origin",
	"Cache-Control",
	"X-Requested-With",
}

var allowedMethods = []string{
	"GET",
	"PUT",
	"POST",
	"OPTIONS",
	"HEAD",
	"DELETE",
	"PATCH",
}

// CORS middleware.
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", strings.Join(allowedHeaders, ", "))
		c.Writer.Header().Set("Access-Control-Allow-Methods", strings.Join(allowedMethods, ", "))

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
