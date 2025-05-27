package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const banNotFoundError = "ban not found"

func BanNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", banNotFoundError, id))
}

func IsBanNotFound(err error) bool {
	return strings.Contains(err.Error(), banNotFoundError)
}
