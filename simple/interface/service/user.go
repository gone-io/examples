package service

type UserInfo struct {
	ID   int64
	Name string
}

type User interface {
	GetUserInfo(id int64) (*UserInfo, error)
}
