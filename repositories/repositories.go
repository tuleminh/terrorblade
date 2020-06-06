package repositories

import (
	"github.com/google/wire"

	"terrorblade"
)

var RepositorySet = wire.NewSet(
	NewUserRepository,
	wire.Bind(new(terrorblade.UserRepository), new(*UserRepository)),
)
