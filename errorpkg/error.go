package errorpkg

import "github.com/dangviethung096/core"

// Internal Server Error
var (
	ERROR_INTERNAL_SERVER            = core.NewError(1000, "Internal server error")
	ERROR_ACCOUNT_EXISTED            = core.NewError(1001, "Account existed")
	ERROR_ACCOUNT_NOT_EXISTED        = core.NewError(1002, "Account not existed")
	ERROR_WRONG_PASSWORD             = core.NewError(1003, "Wrong password")
	ERROR_WRONG_USERNAME_OR_PASSWORD = core.NewError(1004, "Username or password is incorrect")
	ERROR_CANNOT_LOGOUT              = core.NewError(1005, "Cannot logout")
	ERROR_CUSTOMER_EXISTED           = core.NewError(1006, "Customer existed")
	ERROR_BAD_REQUEST                = core.NewError(1007, "Bad request")
	ERROR_NOT_EXISTED                = core.NewError(1010, "not existed")
	ERROR_NOT_FOUND_PRODUCT          = core.NewError(1011, "not found product")
	ERROR_NOT_FOUND_RESPONSE_MESSAGE = core.NewError(1012, "not found response message")
	ERROR_NOT_FOUND_SUBSCRIBER       = core.NewError(1013, "not found subscriber")
	ERROR_FILE_NOT_FOUND             = core.NewError(1014, "file not found")
	ERROR_NOT_FOUND_IN_DB            = core.NewError(1015, "not found in db")
)
