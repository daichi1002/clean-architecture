package validator

import (
	"context"
	"log"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ValidateStructWithContext validation.ValidateStructWithContext wrapper.
// And return gRPC invalid argument error with errdetails.BadRequest.
func ValidateStructWithContext(ctx context.Context, structPtr interface{}, fields ...*validation.FieldRules) error {
	if err := validation.ValidateStructWithContext(ctx, structPtr, fields...); err != nil {
		if validationErrors, ok := err.(validation.Errors); ok {
			sts := status.New(codes.InvalidArgument, codes.InvalidArgument.String())
			violations := make([]*errdetails.BadRequest_FieldViolation, 0)
			for key, value := range validationErrors {
				violations = append(violations, &errdetails.BadRequest_FieldViolation{
					Field:       key,
					Description: value.Error(),
				})
			}
			log.Printf("validation error, error: %v", validationErrors)
			stsWithDetails, _ := sts.WithDetails(&errdetails.BadRequest{FieldViolations: violations})
			return stsWithDetails.Err()
		}
		return err
	}
	return nil
}
