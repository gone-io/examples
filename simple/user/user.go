package user

import (
	"fmt"
	"github.com/gone-io/examples/simple/interface/service"
	"github.com/gone-io/gone"
	"github.com/gone-io/gone/goner/redis"
)

// 1. 定义 Goner：userService
type userService struct {
	gone.Flag             //2. 聚合 gone.Flag，使其实现gone.Goner接口成为一个Goner
	cache     redis.Cache `gone:"gone-redis-cache"` //4. 标记需要注入的依赖，这里表示在`cache`属性上注入一个ID=`gone-redis-cache`的 Goner 组件
}

func (s *userService) GetUserInfo(id int64) (*service.UserInfo, error) {
	info := new(service.UserInfo)
	key := fmt.Sprintf("user-%d", id)
	return info, s.cache.Get(key, info) //5.使用注入的依赖完成业务
}

// NewUserService 3. 定义 `userService` 构造函数
func NewUserService() gone.Goner {
	return &userService{}
}
