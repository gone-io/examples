package controller

import (
	"github.com/gone-io/gone"
	"github.com/gone-io/gone/goner/gin"
)

//go:gone
func NewUserExtController() gone.Goner {
	return &userExtController{}
}

type userExtController struct {
	gone.Flag

	authRouter gin.IRouter `gone:"router-auth"`
	pubRouter  gin.IRouter `gone:"router-pub"`
}

func (ctr *userExtController) Mount() gin.MountError {

	ctr.
		authRouter.
		Group("/users/:userId/extends").
		POST("", ctr.createUserExtends).
		GET("", ctr.getUserExtends).
		DELETE("/:extId", ctr.deleteUserExtend).
		PUT("/:extId", ctr.updateUserExtend)

	return nil
}

func (ctr *userExtController) createUserExtends(ctx *gin.Context) (any, error) {
	return nil, nil
}

func (ctr *userExtController) getUserExtends(ctx *gin.Context) (any, error) {
	return nil, nil
}

func (ctr *userExtController) deleteUserExtend(ctx *gin.Context) (any, error) {
	return nil, nil
}

func (ctr *userExtController) updateUserExtend(ctx *gin.Context) (any, error) {
	return nil, nil
}
