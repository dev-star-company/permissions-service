package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const firstloginNotFoundError = "firstlogin not found"

func FirstLoginNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", firstloginNotFoundError, id))
}

func IsFirstLoginNotFound(err error) bool {
	return strings.Contains(err.Error(), firstloginNotFoundError)
}
