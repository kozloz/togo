package main

import (
	"errors"

	customError "github.com/kozloz/togo/internal/errors"
	"github.com/kozloz/togo/internal/genproto"
)

// Converts the error to our JSON Error type
func CustomErrorToProto(err error) genproto.Error {
	var cerr customError.Error
	if errors.As(err, &cerr) {
		return genproto.Error{
			ErrorCode: int32(cerr.ErrorCode),
			ErrorDesc: cerr.ErrorDesc,
		}
	}
	return genproto.Error{
		ErrorCode: int32(customError.InternalError.ErrorCode),
		ErrorDesc: err.Error(),
	}
}
