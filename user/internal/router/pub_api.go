package router

import (
	"github.com/gone-io/gone"
	"github.com/gone-io/gone/goner/gin"
	"user-center-service/internal/middleware"
)

const IdPubApiRouter = "router-pub"

//go:gone
func NewOpenApi() (gone.Goner, gone.GonerId) {
	return &pubRouter{}, IdPubApiRouter
}

// 不鉴权分组路由
type pubRouter struct {
	gone.Flag
	gin.IRouter
	root gin.IRouter `gone:"gone-gin-router"`

	auth *middleware.AuthorizeMiddleware `gone:"*"`
	pub  *middleware.PubMiddleware       `gone:"*"`
}

func (r *pubRouter) AfterRevive() gone.AfterReviveError {
	r.IRouter = r.root.Group("/api")
	return nil
}
