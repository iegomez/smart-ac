package helpers

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/iegomez/smart-ac/internal/storage"
)

var errToCode = map[error]codes.Code{
	storage.ErrAlreadyExists:             codes.AlreadyExists,
	storage.ErrDoesNotExist:              codes.NotFound,
	storage.ErrUsedByOtherObjects:        codes.FailedPrecondition,
	storage.ErrUserInvalidUsername:       codes.InvalidArgument,
	storage.ErrUserPasswordLength:        codes.InvalidArgument,
	storage.ErrInvalidUsernameOrPassword: codes.Unauthenticated,
}

func ErrToRPCError(err error) error {
	cause := errors.Cause(err)

	// if the err has already a gRPC status (it is a gRPC error), just
	// return the error.
	if code := status.Code(cause); code != codes.Unknown {
		return cause
	}

	code, ok := errToCode[cause]
	if !ok {
		code = codes.Unknown
	}
	return grpc.Errorf(code, cause.Error())
}
