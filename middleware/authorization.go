package middleware

import (
	"app/constant"
	"app/pkg"
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": constant.ErrAuthorization.Error()})
			return
		}

		tokenParts := strings.Split(token, " ")
		if len(tokenParts) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": constant.ErrAuthorization.Error()})
			return
		}
		tokenType := tokenParts[0]
		tokenValue := tokenParts[1]

		if tokenType != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": constant.ErrAuthorization.Error()})
			return
		}

		user, err := pkg.VerifyTokenHeader(tokenValue)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": constant.ErrAuthorization.Error()})
			return
		}
		c.Set("user", user)
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, "user", user)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
