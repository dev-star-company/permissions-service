package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const rolehaspermissionNotFoundError = "rolehaspermission not found"

func RoleHasPermissionNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", rolehaspermissionNotFoundError, id))
}

func IsRoleHasPermissionNotFound(err error) bool {
	return strings.Contains(err.Error(), rolehaspermissionNotFoundError)
}
