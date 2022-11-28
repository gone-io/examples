package user

import "github.com/gone-io/gone"

//go:gone
func NewUserService() gone.Goner {
	return &svc{}
}

type svc struct {
	gone.Flag
}
