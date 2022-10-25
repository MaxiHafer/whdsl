package emitter

import (
	"context"

	"whdsl/cmd/whdsl/internal"
)

var _ internal.Server = &Server{}

type Server struct {
}

func (s Server) Run(ctx context.Context) error {
	// TODO implement me
	panic("implement me")
}

func (s Server) Close(ctx context.Context) error {
	// TODO implement me
	panic("implement me")
}
