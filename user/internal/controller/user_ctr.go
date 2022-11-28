package controller

import (
	"github.com/gone-io/gone"
	"github.com/gone-io/gone/goner/gin"
	"net/http"
)

//go:gone
func NewUserController() gone.Goner {
	return &userController{}
}

type userController struct {
	gone.Flag

	authRouter gin.IRouter `gone:"router-auth"`
	pubRouter  gin.IRouter `gone:"router-pub"`
}

func (ctr *userController) Mount() gin.MountError {

	//不需要鉴权的路由分组
	ctr.
		pubRouter.
		Group("/users").
		POST("/sign-up", ctr.signUp).
		POST("/sign-in", ctr.signIn).
		GET("/:userId/page", ctr.getUserPage)

	//需要鉴权的路由分组
	ctr.
		authRouter.
		Group("/users").
		GET("/sign-out", ctr.signOut).
		POST("", ctr.createUser).
		GET("", ctr.getUserList).
		PUT("/:userId", ctr.updateUser).
		DELETE("/:userId", ctr.deleteUser).
		GET("/:userId", ctr.getUser)

	return nil
}

// 注册
func (ctr *userController) signUp(ctx *gin.Context) (any, error) {
	return nil, nil
}

// 登录
func (ctr *userController) signIn(ctx *gin.Context) (any, error) {
	return nil, nil
}

// 注销
func (ctr *userController) signOut(ctx *gin.Context) (any, error) {
	return nil, nil
}

func (ctr *userController) createUser(ctx *gin.Context) (any, error) {
	return nil, nil
}

func (ctr *userController) updateUser(ctx *gin.Context) (any, error) {
	return nil, nil
}

func (ctr *userController) deleteUser(ctx *gin.Context) (any, error) {
	return nil, nil
}

func (ctr *userController) getUser(ctx *gin.Context) (any, error) {
	return nil, nil
}

func (ctr *userController) getUserList(ctx *gin.Context) (any, error) {
	return nil, nil
}

// 获取某User的page（html 页面）
func (ctr *userController) getUserPage(ctx *gin.Context) (any, error) {
	ctx.HTML(http.StatusOK, "users/page.tmpl", gin.R{
		"title": "user pages",
	})
	return nil, nil
}
