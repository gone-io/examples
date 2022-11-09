package main

import (
	"github.com/gone-io/examples/simple/student"
	"github.com/gone-io/examples/simple/user"
	"github.com/gone-io/gone"
	"github.com/gone-io/gone/goner/redis"
)

//⚠️ 代码要正常运行，需要配置在 `config/local.properties` 中配置 redis de 地址（`redis.server`） 和 密码（`redis.password`）

// 1. 增加 main 函数，调用 gone.Run
func main() {
	//2. 给 gone.Run 方法提供一个 `Priest` 函数
	gone.Run(Priest)
}

func Priest(cemetery gone.Cemetery) error {
	// 4. 调用redis 的 Priest，对redis的相关组件完成"安葬"
	_ = redis.Priest(cemetery)

	// 3. "安葬"
	cemetery.Bury(user.NewUserService())       // 3.1 在 `Priest` 函数中 "安葬" `user.NewUserService()`构造出来的 Goner
	cemetery.Bury(student.NewStudentService()) // 3.2 在 `Priest` 函数中 "安葬" `user.NewStudentService()`构造出来的 Goner
	return nil
}
