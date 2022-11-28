package user

import "github.com/gone-io/gone"

//go:gone
func NewUserExtendService() gone.Goner {
	return &extSvc{}
}

type extSvc struct {
	gone.Flag
}
