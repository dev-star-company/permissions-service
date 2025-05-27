package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const serviceNotFoundError = "service not found"

func ServiceNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", serviceNotFoundError, id))
}

func IsServiceNotFound(err error) bool {
	return strings.Contains(err.Error(), serviceNotFoundError)
}
