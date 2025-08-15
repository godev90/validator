package faults

import (
	_ "embed"
	"log"
)

//go:embed builtin_list.yaml
var builtinList []byte

func init() {
	// loadYamlFile("builtin_list.yaml")
	builtinYaml = NewYamlPackage()
	err := builtinYaml.LoadBytes(builtinList)

	if err != nil {
		log.Fatal("failed load file: ", err)
	}

	ErrBadRequest = builtin("err_bad_request")
	ErrUnauthorized = builtin("err_unauthorized")
	ErrPaymentRequired = builtin("err_payment_required")
	ErrForbidden = builtin("err_forbidden")
	ErrNotFound = builtin("err_not_found")
	ErrMethodNotAllowed = builtin("err_method_not_allowed")
	ErrNotAcceptable = builtin("err_not_acceptable")
	ErrProxyAuthRequired = builtin("err_proxy_auth_required")
	ErrRequestTimeout = builtin("err_request_timeout")
	ErrConflict = builtin("err_conflict")
	ErrGone = builtin("err_gone")
	ErrLengthRequired = builtin("err_length_required")
	ErrPreconditionFailed = builtin("err_precondition_failed")
	ErrPayloadTooLarge = builtin("err_payload_too_large")
	ErrURITooLong = builtin("err_uri_too_long")
	ErrUnsupportedMediaType = builtin("err_unsupported_media_type")
	ErrRangeNotSatisfiable = builtin("err_range_not_satisfiable")
	ErrExpectationFailed = builtin("err_expectation_failed")
	ErrIMTeapot = builtin("err_im_teapot")
	ErrMisdirectedRequest = builtin("err_misdirected_request")
	ErrUnprocessableEntity = builtin("err_unprocessable_entity")
	ErrLocked = builtin("err_locked")
	ErrFailedDependency = builtin("err_failed_dependency")
	ErrTooEarly = builtin("err_too_early")
	ErrUpgradeRequired = builtin("err_upgrade_required")
	ErrPreconditionRequired = builtin("err_precondition_required")
	ErrTooManyRequests = builtin("err_too_many_requests")
	ErrInternalServerError = builtin("err_internal_server_error")
	ErrNotImplemented = builtin("err_not_implemented")
	ErrBadGateway = builtin("err_bad_gateway")
	ErrBadGatewayf = builtin("err_bad_gateway_f")
	ErrServiceUnavailable = builtin("err_service_unavailable")
	ErrGatewayTimeout = builtin("err_gateway_timeout")
	ErrUnprocessable = builtin("err_unprocessable")
	ErrHTTPVersionNotSupported = builtin("err_http_version_not_supported")
	ErrInsufficientStorage = builtin("err_insufficient_storage")

	// Validator / input
	ErrRequired = builtin("err_required")
	ErrBelowMinimum = builtin("err_below_minimum")
	ErrAboveMaximum = builtin("err_above_maximum")
	ErrMustBeEmail = builtin("err_must_be_email")
	ErrMustBeDigit = builtin("err_must_be_digit")
	ErrMustBeAlphanum = builtin("err_must_be_alphanum")
	ErrMustBeAlphabet = builtin("err_must_be_alphabet")
	ErrLengthBelowMinimum = builtin("err_length_below_minimum")
	ErrLengthAboveMaximum = builtin("err_length_above_maximum")
	ErrInvalidDateFormat = builtin("err_invalid_date_format")
	ErrInvalidDatetimeFormat = builtin("err_invalid_datetime_format")
	ErrUnsupportedDataType = builtin("err_unsupported_data_type")
	ErrInvalidParameter = builtin("err_invalid_parameter")
	ErrMustBeName = builtin("err_must_be_name")
	ErrMustBeText = builtin("err_must_be_text")
	ErrMustBeOneOf = builtin("err_must_be_one_of")
	ErrUnsupportedContentType = builtin("err_unsupported_content_type")
	ErrCannotBeNull = builtin("err_cannot_be_null")
	ErrTypeMismatch = builtin("err_type_mismatch")
	ErrInvalidNumericFormat = builtin("err_invalid_numeric_format")
	ErrInvalidFloatNumber = builtin("err_invalid_float_number")
	ErrInvalidIntegerNumber = builtin("err_invalid_integer_number")
}

var (
	builtinYaml YamlPackage

	ErrBadRequest              Error
	ErrUnauthorized            Error
	ErrPaymentRequired         Error
	ErrForbidden               Error
	ErrNotFound                Error
	ErrMethodNotAllowed        Error
	ErrNotAcceptable           Error
	ErrProxyAuthRequired       Error
	ErrRequestTimeout          Error
	ErrConflict                Error
	ErrGone                    Error
	ErrLengthRequired          Error
	ErrPreconditionFailed      Error
	ErrPayloadTooLarge         Error
	ErrURITooLong              Error
	ErrUnsupportedMediaType    Error
	ErrRangeNotSatisfiable     Error
	ErrExpectationFailed       Error
	ErrIMTeapot                Error
	ErrMisdirectedRequest      Error
	ErrUnprocessableEntity     Error
	ErrLocked                  Error
	ErrFailedDependency        Error
	ErrTooEarly                Error
	ErrUpgradeRequired         Error
	ErrPreconditionRequired    Error
	ErrTooManyRequests         Error
	ErrInternalServerError     Error
	ErrNotImplemented          Error
	ErrBadGateway              Error
	ErrBadGatewayf             Error
	ErrServiceUnavailable      Error
	ErrGatewayTimeout          Error
	ErrUnprocessable           Error
	ErrHTTPVersionNotSupported Error
	ErrInsufficientStorage     Error

	ErrRequired               Error
	ErrBelowMinimum           Error
	ErrAboveMaximum           Error
	ErrMustBeEmail            Error
	ErrMustBeDigit            Error
	ErrMustBeAlphanum         Error
	ErrMustBeAlphabet         Error
	ErrLengthBelowMinimum     Error
	ErrLengthAboveMaximum     Error
	ErrInvalidDateFormat      Error
	ErrInvalidDatetimeFormat  Error
	ErrUnsupportedDataType    Error
	ErrInvalidParameter       Error
	ErrMustBeName             Error
	ErrMustBeText             Error
	ErrMustBeOneOf            Error
	ErrUnsupportedContentType Error
	ErrCannotBeNull           Error
	ErrTypeMismatch           Error
	ErrInvalidNumericFormat   Error
	ErrInvalidFloatNumber     Error
	ErrInvalidIntegerNumber   Error
)
