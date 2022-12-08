package errors

import "errors"

var (
	MethodNotAllowed = errors.New("method not allowed")
	UnknownError = errors.New("unknown error")
	PlacementsRequestIncorrectParametersError = errors.New("placements request incorrect parameters")
	PlacementsResponseIncorrectParametersError = errors.New("placements response incorrect parameters")
	PlacementsRequestError = errors.New("placements request returned an error")
)
