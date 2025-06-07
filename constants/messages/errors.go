package messages

const (
	ErrInternalServer = "sorry, something went wrong"
	ErrInvalidJSON    = "invalid JSON body"
	ErrInvalidMethod  = "invalid request method"
	ErrInvalidUserID  = "invalid user id"
	ErrLogin          = "wrong email or password"
	ErrNoCookie       = "cookie not found"
	ErrInvalidToken   = "invalid token"
	ErrNoToken        = "authorization is required"
	ErrForbidden      = "access denied"
	ErrDuplicateEmail = "this email already been used"
	ErrNotFound       = "resource not found"
	ErrInvalidBody    = "invalid request body"
)