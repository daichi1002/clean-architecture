package validator

import (
	"clean-architecture/pb"
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type userValidator struct{}

func NewUserValidator() *userValidator {
	return &userValidator{}
}

func (u *userValidator) Validate(ctx context.Context, req interface{}) (bool, error) {
	switch request := req.(type) {
	case *pb.ListUsersRequest:
		return true, u.validateListUsers(ctx, request)
	}
	return false, nil
}

func (u *userValidator) validateListUsers(ctx context.Context, req *pb.ListUsersRequest) error {
	rules := make([]*validation.FieldRules, 0)
	rules = append(rules, validation.Field(&req.PageSize, validation.Min(1), validation.Max(100)))
	return ValidateStructWithContext(ctx, req, rules...)
}
