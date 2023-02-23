package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// TODO: Authentication things...
// TODO: This is going to be a register endpoint.
// TODO: Should we authenticate right after user registers, or redirect to login??

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Create User
//
//	@Summary	Create User
//	@Tags		User
//	@Id			create-user
//	@Param		credentials	body	CreateUserRequest	true	"New User Info"
//	@Produce	json
//	@Success	201
//	@Router		/users [post]
func (h *Handler) CreateUser(c echo.Context) error {
	var req CreateUserRequest
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid create user credentials")
	}

	err = h.svc.Create(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}
