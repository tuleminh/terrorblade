package handlers

import (
	"github.com/google/wire"

	_commonHandlers "terrorblade/handlers"
)

var HandlerSet = wire.NewSet(
	_commonHandlers.NewHealthCheckHandler,
	NewUserHandler,
	NewHandlers,
)

// NewHandlers returns a new instance of Handlers.
func NewHandlers(
	healthCheck *_commonHandlers.HealthCheckHandler,
	user *UserHandler,
) *Handlers {
	return &Handlers{
		HealthCheck: healthCheck,
		User:        user,
	}
}

// Handlers contains all HTTP handlers.
type Handlers struct {
	HealthCheck *_commonHandlers.HealthCheckHandler
	User        *UserHandler
}
