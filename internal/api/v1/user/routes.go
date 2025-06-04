package user

import (
	"github.com/gin-gonic/gin"
	u "github.com/sinasadeghi83/ghashoghchi/internal/user"
	"gorm.io/gorm"
)

func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := u.NewGormUserRepo(db)
	svc := u.NewGormUserService(repo)
	handler := NewHandler(svc)

	rg.POST("/register", handler.Register)
	rg.POST("/login", handler.Login)

	rg.GET("/profile", handler.CheckAuth, handler.profile)
}
