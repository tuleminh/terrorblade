package services

import "github.com/google/wire"

var ServiceSet = wire.NewSet(NewUserService)
