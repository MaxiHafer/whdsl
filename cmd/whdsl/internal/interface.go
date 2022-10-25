package internal

import (
	"context"
)

type Server interface {
	Run(ctx context.Context) error
}
