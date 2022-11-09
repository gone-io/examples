package student

import (
	"github.com/gone-io/examples/simple/interface/service"
	"github.com/gone-io/gone"
	"github.com/gone-io/gone/goner/logrus"
)

// 1. 定义 Goner：studentService
type studentService struct {
	gone.Flag                  // 2.  聚合 gone.Flag，使其实现gone.Goner接口成为一个Goner
	service.User `gone:"*"`    //4. 聚合 service.User，这里的 `gone:"*"` 表示 `按类型注入` 一个Goner
	log          logrus.Logger `gone:"gone-logger"` //6. 注入一个用于日志打印的Goner
}

func (s *studentService) GetStudentInfo(id int64) (*service.UserInfo, error) {
	return s.GetUserInfo(id) //5. 调用 User 的 `GetUserInfo` 来实现 `GetStudentInfo`方法
}

// AfterRevive 6.该方法会在 studentService 属性被注入完成后自动运行
func (s *studentService) AfterRevive() gone.AfterReviveError {
	s.log.Infof("======> Student Service is running ======>")

	info, err := s.GetUserInfo(100)
	if err != nil {
		s.log.Errorf("get info err:%v", err) //调用日志Goner，打印错误日志
	} else {
		s.log.Infof("student info:%v", info) //调用日志Goner，打印student info
	}
	return nil
}

// NewStudentService 3. 定义 `studentService` 构造函数
func NewStudentService() gone.Goner {
	return &studentService{}
}
