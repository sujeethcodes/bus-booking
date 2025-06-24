package constant

// Status Codes

const (
	SUCCESS            = 200
	ACCEPTED           = 202
	BAD_REQUEST        = 400
	UNAUTHORIZED       = 401
	METHOD_NOT_ALLOWED = 405
	INTERNAL_ERROR     = 500
)

// Success Message
const (
	SUCCESS_MESSAGE        = "success"
	DELETE_SUCCESS_MESSAGE = "user deleted successfully"
)

// Error Message
const (
	BAD_REQUEST_MESSAGE       = "Invalid request"
	INVAILD_USER_TYPE_MESSAGE = "Invalid user type request"
	TOKEN_GENERATE_FAILED     = "Token generation failed"
	EMAIL_EXIST               = "Email already have a account"
	DB_USER_UPDATE_FAIL       = "user update failed"
	DB_USER_DELETE_FAIL       = "user delete failed"
	INVAILD_USER_ID           = "invaild user id"
	USER_ACCOUNT              = "this is user account"
)
