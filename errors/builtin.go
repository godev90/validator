package errors

import "strings"

var (
	ErrUnauthorized error

	ErrDuplicateEntry error

	ErrFieldRequired error

	ErrFieldBelowMinimum = func(min any) error {
		return registerBuiltinError("ErrFieldBelowMinimum", min)
	}

	ErrFieldAboveMaximum = func(max any) error {
		return registerBuiltinError("ErrFieldAboveMaximum", max)
	}

	ErrFieldMustBeEmail      error
	ErrFieldMustBeDigit      error
	ErrFieldMustBeAlphanum   error
	ErrFieldMustBeAlphabet   error
	ErrInvalidDateFormat     error
	ErrInvalidDatetimeFormat error

	ErrFieldLengthBelowMinimum = func(min int) error {
		return registerBuiltinError("ErrFieldLengthBelowMinimum", min)
	}

	ErrFieldLengthAboveMaximum = func(max int) error {
		return registerBuiltinError("ErrFieldLengthAboveMaximum", max)
	}

	ErrFieldUnsupportedDataType error

	ErrFieldInvalidParam = func(param any) error {
		return registerBuiltinError("ErrFieldInvalidParam", param)
	}

	ErrFieldMustBeName error

	ErrFieldMustBeText error

	ErrFieldMustBeOneOf = func(allowed []string) error {
		return registerBuiltinError("ErrFieldMustBeOneOf", strings.Join(allowed, ", "))
	}

	ErrUnsupportedContentType error
	ErrCannotBeNull           error
	ErrTypeMismatch           error
	ErrUnknownSession         error

	ErrForbidden            error
	ErrNotFound             error
	ErrTooManyRequest       error
	ErrNotAllowed           error
	ErrInvalidInput         error
	ErrInternalServer       error
	ErrBadRequest           error
	ErrRequestTimeout       error
	ErrUnprocessable        error
	ErrServiceUnavailable   error
	ErrBadGateway           error
	ErrInvalidNumericFormat error
	ErrInvalidFloatNumber   error
	ErrInvalidIntegerNumber error
)

func init() {
	loadYamlFile("builtin_list.yaml")

	ErrUnauthorized = registerBuiltinError("ErrUnauthorized")
	ErrDuplicateEntry = registerBuiltinError("ErrDuplicateEntry")
	ErrFieldRequired = registerBuiltinError("ErrFieldRequired")
	ErrFieldMustBeEmail = registerBuiltinError("ErrFieldMustBeEmail")
	ErrFieldMustBeDigit = registerBuiltinError("ErrFieldMustBeDigit")
	ErrFieldMustBeAlphanum = registerBuiltinError("ErrFieldMustBeAlphanum")
	ErrFieldMustBeAlphabet = registerBuiltinError("ErrFieldMustBeAlphabet")
	ErrInvalidDateFormat = registerBuiltinError("ErrInvalidDateFormat")
	ErrInvalidDatetimeFormat = registerBuiltinError("ErrInvalidDatetimeFormat")
	ErrFieldUnsupportedDataType = registerBuiltinError("ErrFieldUnsupportedDataType")
	ErrFieldMustBeName = registerBuiltinError("ErrFieldMustBeName")
	ErrFieldMustBeText = registerBuiltinError("ErrFieldMustBeText")
	ErrUnsupportedContentType = registerBuiltinError("ErrUnsupportedContentType")
	ErrCannotBeNull = registerBuiltinError("ErrCannotBeNull")
	ErrTypeMismatch = registerBuiltinError("ErrTypeMismatch")
	ErrUnknownSession = registerBuiltinError("ErrUnknownSession")
	ErrForbidden = registerBuiltinError("ErrForbidden")
	ErrNotFound = registerBuiltinError("ErrNotFound")
	ErrTooManyRequest = registerBuiltinError("ErrTooManyRequest")
	ErrNotAllowed = registerBuiltinError("ErrNotAllowed")
	ErrInvalidInput = registerBuiltinError("ErrInvalidInput")
	ErrInternalServer = registerBuiltinError("ErrInternalServer")
	ErrBadRequest = registerBuiltinError("ErrBadRequest")
	ErrRequestTimeout = registerBuiltinError("ErrRequestTimeout")
	ErrUnprocessable = registerBuiltinError("ErrUnprocessable")
	ErrServiceUnavailable = registerBuiltinError("ErrServiceUnavailable")
	ErrBadGateway = registerBuiltinError("ErrBadGateway")
	ErrInvalidNumericFormat = registerBuiltinError("ErrInvalidNumericFormat")
	ErrInvalidFloatNumber = registerBuiltinError("ErrInvalidFloatNumber")
	ErrInvalidIntegerNumber = registerBuiltinError("ErrInvalidFloatNumber")
}
