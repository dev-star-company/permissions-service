package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const loginattemptsNotFoundError = "loginattempts not found"

func LoginAttemptsNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", loginattemptsNotFoundError, id))
}

func IsLoginAttemptsNotFound(err error) bool {
	return strings.Contains(err.Error(), loginattemptsNotFoundError)
}
