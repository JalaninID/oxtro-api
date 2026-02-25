package middleware

import "connectrpc.com/connect"

func NewOrganizationInterceptor() connect.UnaryInterceptorFunc {
	return NewAuthInterceptor(nil)
}
