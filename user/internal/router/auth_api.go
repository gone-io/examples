package router

import (
	"github.com/gone-io/gone"
	"github.com/gone-io/gone/goner/gin"
	"user-center-service/internal/middleware"
)

const IdAuthApiRouter = "router-auth"

//go:gone
func NewAuthApi() (gone.Goner, gone.GonerId) {
	return &authApiRouter{}, IdAuthApiRouter
}

// 鉴权路由分组
type authApiRouter struct {
	gone.Flag
	gin.IRouter
	root gin.IRouter `gone:"gone-gin-router"`

	auth *middleware.AuthorizeMiddleware `gone:"*"`
	pub  *middleware.PubMiddleware       `gone:"*"`
}

func (r *authApiRouter) AfterRevive() gone.AfterReviveError {
	r.IRouter = r.root.Group("/api", r.auth.Next)
	return nil
}
