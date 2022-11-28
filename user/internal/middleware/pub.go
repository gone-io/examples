package middleware

import (
	"github.com/gone-io/gone"
	"github.com/gone-io/gone/goner/gin"
	"github.com/gone-io/gone/goner/logrus"
	"os"
	"path"
	"strings"
)

//go:gone
func NewPlantMiddleware() gone.Goner {
	return &PubMiddleware{}
}

// PubMiddleware 公共中间件
type PubMiddleware struct {
	gone.Flag
	logrus.Logger `gone:"gone-logger"`
}

func (m *PubMiddleware) Next(ctx *gin.Context) (interface{}, error) {
	if err := servePublic(ctx); err != nil {
		return nil, err
	}
	return nil, nil
}

// server for public front static assets
func servePublic(ctx *gin.Context) error {
	url := ctx.Request.URL.String()
	if !strings.HasPrefix(url, "/api/") {
		filepath := path.Join("public", url)

		stat, err := os.Stat(filepath)
		if err != nil {
			if os.IsNotExist(err) {
				ctx.File("public/index.html")
				ctx.Abort()
				return nil
			} else {
				return gin.ToError(err)
			}
		}

		if stat.IsDir() {
			ctx.File(path.Join(filepath, "index.html"))
			ctx.Abort()
		} else {
			ctx.File(filepath)
			ctx.Abort()
		}
	}
	return nil
}
