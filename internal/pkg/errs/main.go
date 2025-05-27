package errs

import (
	"errors"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	requesterIdRequiredError = "requester id is required"
	invalidForeignKeyError   = "invalid foreign key"
	listingError             = "error listing"
	startProducts            = "starting transaction"
	commitProducts           = "commiting transaction"
	errorDelete              = "error deleting"
	errorSave                = "error saving"
	errorCreate              = "error creating"
	invalidOrderByValue      = "invalid order by value"
)

func RequesterIdRequired() error {
	return errors.New(requesterIdRequiredError)
}

func CreateError(entity string, err error) error {
	return status.Errorf(codes.Internal, "%s: %v", errorCreate, err)
}

func SavingError(entity string, err error) error {
	return status.Errorf(codes.Internal, "%s: %v", errorSave, err)
}

func IsRequesterIdRequired(err error) bool {
	return strings.Contains(err.Error(), requesterIdRequiredError)
}

func DeleteError(entity string, err error) error {
	return status.Errorf(codes.Internal, "%s: %v", errorDelete, err)
}

func CommitProductsError(err error) error {
	return status.Error(codes.Internal, fmt.Sprintf("%s: %v", commitProducts, err))
}

func StartProductsError(err error) error {
	return status.Error(codes.Internal, fmt.Sprintf("%s: %v", startProducts, err))
}

func ListingError(entity string, err error) error {
	return status.Error(codes.Internal, fmt.Sprintf("%s: %v", listingError, err))
}

func InvalidForeignKey(err error) error {
	return status.Error(codes.InvalidArgument, fmt.Sprintf("%s: %v", invalidForeignKeyError, err))
}

func InvalidOrderByValue(err error) error {
	return status.Error(codes.InvalidArgument, fmt.Sprintf("%s: %v", invalidOrderByValue, err))
}
