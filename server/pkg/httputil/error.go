package httputil

// TODO: Are we going to use this, sure! But how we want to use it?

type HTTPError struct {
	Code    int    `json:"code" example:"404"`
	Message string `json:"message" example:"needle not found in haystack"`
	ErrID   string `json:"errID" example:"ERR_USER_NOTFOUND"`
	Err     error  `json:"-"`
}

func NewHTTPError(err error, id string, code int) *HTTPError {
	return &HTTPError{
		Code:    code,
		Message: err.Error(),
		ErrID:   id,
	}
}
