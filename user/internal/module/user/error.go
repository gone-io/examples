package user

import "user-center-service/internal/pkg/utils"

const ModuleId = 1

const (
	Error1 = utils.AppId*utils.AppModuleNumber*utils.ModuleErrNumber + ModuleId*utils.ModuleErrNumber + iota
	Error2
	Error3
)
