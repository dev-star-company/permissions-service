package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const permissionNotFoundError = "permission not found"

func PermissionNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", permissionNotFoundError, id))
}

func IsPermissionNotFound(err error) bool {
	return strings.Contains(err.Error(), permissionNotFoundError)
}
