package router

import (
	"github.com/gone-io/gone"
	"github.com/gone-io/gone/goner/gin"
	"user-center-service/internal/middleware"
)

const IdRouterPub = "router-root"

//go:gone
func NewRootRouter() (gone.Goner, gone.GonerId) {
	return &rootRouter{}, IdRouterPub
}

// 根路由分组
type rootRouter struct {
	gone.Flag
	gin.IRouter `gone:"gone-gin-router"`
	pub         *middleware.PubMiddleware `gone:"*"`
}

func (r *rootRouter) AfterRevive() gone.AfterReviveError {
	r.IRouter.Use(r.pub.Next)
	return nil
}
