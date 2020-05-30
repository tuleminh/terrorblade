package repositories

import "github.com/google/wire"

var RepositorySet = wire.NewSet(NewUserRepository)
