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
	rest.RespondCreated(c, "user created successfully", newUser2)
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

	rest.RespondCreated(c, "user logged in successfully", gin.H{
		"token": token,
		"user":  user,
	})
}

func (h *Handler) profile(c *gin.Context) {
	user, e := c.Get("currentUser")
	if !e {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Missing or invalid JWT token",
		})
		c.Header("WWW-Authenticate", "Bearer")
		return
	}
	rest.RespondOK(c, "", user)
}

func (h *Handler) UpdProfile(c *gin.Context) {
	user, exists := c.Get("currentUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Missing or invalid JWT token",
		})
		c.Header("WWW-Authenticate", "Bearer")
		return
	}

	currentUser, ok := user.(*u.User)
	if !ok {
		rest.RespondError(c, http.StatusInternalServerError, "Invalid user data in context", nil)
		return
	}

	var updateData RegisterInput
	if err := c.ShouldBindJSON(&updateData); err != nil {
		rest.RespondError(c, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	updates := make(map[string]interface{})

	if updateData.FullName != "" {
		updates["full_name"] = updateData.FullName
	}

	if updateData.Phone != "" && updateData.Phone != currentUser.Phone {
		if exists, err := h.Service.PhoneExists(updateData.Phone); err != nil {
			rest.RespondError(c, http.StatusInternalServerError, "Error checking phone availability", err)
			return
		} else if exists {
			rest.RespondError(c, http.StatusConflict, "Phone number already in use", nil)
			return
		}
		updates["phone"] = updateData.Phone
	}

	if updateData.Email != "" && updateData.Email != currentUser.Email {
		if exists, err := h.Service.EmailExists(updateData.Email); err != nil {
			rest.RespondError(c, http.StatusInternalServerError, "Error checking email availability", err)
			return
		} else if exists {
			rest.RespondError(c, http.StatusConflict, "Email already in use", nil)
			return
		}
		updates["email"] = updateData.Email
	}

	if updateData.Password != "" {
		hashedPassword, err := u.HashPassword(updateData.Password)
		if err != nil {
			rest.RespondError(c, http.StatusInternalServerError, "Failed to process password", err)
			return
		}
		updates["password"] = hashedPassword
	}

	if updateData.Address != "" {
		updates["address"] = updateData.Address
	}

	if updateData.ProfileImageBase64 != "" {
		updates["profileImageBase64"] = updateData.ProfileImageBase64
	}

	if updateData.BankInfo != (u.BankInfo{}) {
		updates["bank_info"] = updateData.BankInfo
	}

	if err := h.Service.Update(currentUser.ID, updates); err != nil {
		rest.RespondError(c, http.StatusInternalServerError, "Failed to update profile", err)
		return
	}

	rest.RespondOK(c, "profile updated successfully", nil)
}
