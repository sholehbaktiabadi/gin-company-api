package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sholehbaktiabadi/go-api/config"
	"github.com/sholehbaktiabadi/go-api/controller"
	"github.com/sholehbaktiabadi/go-api/middleware"
	"github.com/sholehbaktiabadi/go-api/repository"
	"github.com/sholehbaktiabadi/go-api/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.ConnectDatabase()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	userService    service.UserService       = service.NewUserService(userRepository)
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
)

func main() {
	defer config.CloneConnection(db)
	r := gin.Default()
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("login", authController.Login)
		authRoutes.POST("register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJwt(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
