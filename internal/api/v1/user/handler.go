package user

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

	newUser2, err := h.Service.Register(newUser)
	if err != nil {
		rest.RespondError(c, http.StatusInternalServerError, "Failed to register user", err)
		return
	}

	// fix
	rest.RespondCreated(c, newUser2)
}

func (h *Handler) Login(c *gin.Context) {
	var loginInfo loginCredentials
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		rest.RespondError(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	user, err := h.Service.Login(loginInfo.Phone, loginInfo.Password)
	if err != nil {
		rest.RespondError(c, http.StatusInternalServerError, "Failed to login", err)
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
	}

	c.Header("tokent", token)

	rest.RespondCreated(c, user)
}
