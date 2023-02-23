package user

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type GetUserResponse struct {
	ID       uuid.UUID
	Username string
	Email    string
	Roles    []Role
}

func NewGetUserResponse(u *User) *GetUserResponse {
	return &GetUserResponse{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Roles:    u.Roles,
	}
}

// Get User something... Godoc!
//
//	@Summary	Get User
//	@Tags		User
//	@Id			get-user
//	@Produce	json
//	@Param		username	path		string	true	"Username"
//	@Success	200			{object}	GetUserResponse
//	@Router		/users/{username} [get]
func (h *Handler) GetUser(c echo.Context) error {
	username := c.Param("username")

	user, err := h.svc.GetByUsername(username)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}

	res := NewGetUserResponse(user)

	return c.JSON(http.StatusOK, res)
}
