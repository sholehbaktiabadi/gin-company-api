package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sholehbaktiabadi/go-api/config"
	"github.com/sholehbaktiabadi/go-api/controller"
	"github.com/sholehbaktiabadi/go-api/repository"
	"github.com/sholehbaktiabadi/go-api/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.ConnectDatabase()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
)

func main() {
	defer config.CloneConnection(db)
	r := gin.Default()
	authRoutes := r.Group("user/")
	{
		authRoutes.POST("login", authController.Login)
		authRoutes.POST("register", authController.Register)
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
