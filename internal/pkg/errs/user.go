package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const userNotFoundError = "user not found"

func UserNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", userNotFoundError, id))
}

func IsUserNotFound(err error) bool {
	return strings.Contains(err.Error(), userNotFoundError)
}
