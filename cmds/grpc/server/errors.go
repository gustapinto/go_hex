package server

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrEmptyRequest() error {
	return status.Error(codes.InvalidArgument, "error.request.empty.or.nil")
}

func ErrInternal(err error) error {
	if err == nil {
		return nil
	}

	return status.Error(codes.Internal, err.Error())
}
