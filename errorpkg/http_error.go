package errorpkg

import (
	"github.com/dangviethung096/core"
)

var (
	HTTP_ERROR_INTERNAL_SERVER    = core.NewHttpError(500, 1000, "Internal server error", nil)
	HTTP_ERROR_BAD_REQUEST        = core.NewHttpError(500, 1001, "Bad Request", nil)
	HTTP_ERROR_NOT_FOUND_CUSTOMER = core.NewHttpError(200, 1002, "Not found customer", nil)
)
