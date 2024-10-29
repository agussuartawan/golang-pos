package errors

import "errors"

var (
	ErrFormatInvalid = errors.New("format invalid")
	ErrNotFound       = errors.New("not found")
	ErrForbidden      = errors.New("forbidden")
	ErrConflict       = errors.New("conflict")
	ErrBadGateway     = errors.New("bad gateway")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrTooManyRequest = errors.New("too many request")

	// company errors
	ErrCompanyNotFound = errors.New("companyId not found")

	// warehouse errors
	ErrWarehouseNotFound = errors.New("warehouseId not found")
)