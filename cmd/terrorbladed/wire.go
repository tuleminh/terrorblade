//+build wireinject

package main

import (
	"github.com/google/wire"

	"terrorblade/cmd/terrorbladed/internal/handlers"
	"terrorblade/cmd/terrorbladed/internal/services"
	"terrorblade/repositories"
)

// InitializeMain returns a new instance of main application.
func InitializeMain() *App {
	wire.Build(MainSet, KitSet, handlers.HandlerSet, services.ServiceSet, repositories.RepositorySet)
	return &App{}
}
