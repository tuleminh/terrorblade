package handlers

import "github.com/google/wire"

var HandlerSet = wire.NewSet(NewUserHandler)
