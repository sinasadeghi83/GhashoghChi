package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sinasadeghi83/ghashoghchi/internal/platform/rest"
	u "github.com/sinasadeghi83/ghashoghchi/internal/user"
)

type Handler struct {
	Service u.UserService
}

func NewHandler(svc u.UserService) *Handler {
	return &Handler{
		Service: svc,
	}
}

func (h *Handler) Register(c *gin.Context) {
	var regInput RegisterInput
	if err := c.ShouldBindJSON(&regInput); err != nil {
		rest.RespondError(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	newUser := u.User{
		FullName:           regInput.FullName,
		Phone:              regInput.Phone,
		Email:              regInput.Email,
		Password:           regInput.Password,
		Role:               regInput.Role,
		Address:            regInput.Address,
		ProfileImageBase64: regInput.ProfileImageBase64,
		BankInfo:           regInput.BankInfo,
	}

	newUser, err := h.Service.Register(newUser)
	if err != nil {
		rest.RespondError(c, http.StatusInternalServerError, "Failed to register user", err)
		return
	}

	rest.RespondCreated(c, newUser)
}
