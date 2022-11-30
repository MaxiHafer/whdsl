package grpc

import (
	"github.com/bufbuild/connect-go"
)

func ErrNotFound(err error) *connect.Error {
	return connect.NewError(connect.CodeNotFound, err)
}

func ErrInternal(err error) *connect.Error {
	return connect.NewError(connect.CodeInternal, err)
}

func ErrExists(err error) *connect.Error {
	return connect.NewError(connect.CodeAlreadyExists, err)
}