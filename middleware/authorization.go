package middleware

import (
	"app/constant"
	"app/pkg"
	"context"
	"strings"

	"connectrpc.com/connect"
	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			connect.NewError(connect.CodeUnauthenticated, constant.ErrAuthorization)
			return
		}

		tokenParts := strings.Split(token, " ")
		if len(tokenParts) != 2 {
			connect.NewError(connect.CodeUnauthenticated, constant.ErrAuthorization)
			return
		}
		tokenType := tokenParts[0]
		tokenValue := tokenParts[1]

		if tokenType != "Bearer" {
			connect.NewError(connect.CodeUnauthenticated, constant.ErrAuthorization)
			return
		}

		user, err := pkg.VerifyTokenHeader(tokenValue)
		if err != nil {
			connect.NewError(connect.CodeUnauthenticated, constant.ErrAuthorization)
			return
		}
		c.Set("user", user)
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, "user", user)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
