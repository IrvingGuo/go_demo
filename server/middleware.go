package server

import (
	"net/http"
	"resource-plan-improvement/api"
	"resource-plan-improvement/config"
	"resource-plan-improvement/service"
	"time"

	"github.com/gin-gonic/gin"
)

// cors handle cors request
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token, X-Token,X-User-Id")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		// allow all OPTIONS
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

// Logger instances a Logger middleware for Gin.
// copy from photoprism
func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		// clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		if raw != "" {
			path = path + "?" + raw
		}

		// Use debug level to keep production logs clean.
		log.Debugf("http: %s %s (%3d) [%v]",
			method,
			path,
			statusCode,
			latency,
		)
	}
}

func TokenVerifier() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Infof("%s %s", c.Request.Method, c.Request.URL)
		header := c.Request.Header
		token, ok := header["Authorization"]
		if !ok {
			api.MiddlewareAbortWithJson(c, "no token provided in the request")
			return
		}
		if userId, err := service.ValidateToken(token[0]); err != nil {
			api.MiddlewareAbortWithJson(c, "invalid token, please login again")
		} else {
			c.Set(config.CTX_KEY_USER_ID, userId)
		}
	}
}
