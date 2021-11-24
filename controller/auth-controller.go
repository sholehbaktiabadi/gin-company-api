package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sholehbaktiabadi/go-api/dto"
	"github.com/sholehbaktiabadi/go-api/entity"
	"github.com/sholehbaktiabadi/go-api/helper"
	"github.com/sholehbaktiabadi/go-api/service"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDto dto.LoginDto
	errDTO := ctx.ShouldBind(&loginDto)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDto.Email, loginDto.Password)
	if v, ok := authResult.(entity.User); ok {
		generateToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generateToken
		res := helper.BuildResponse(true, "OK", v)
		ctx.JSON(http.StatusOK, res)
		return
	} else {
		res := helper.BuildErrorResponse("Please check aain your credential", "invalid credential", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
	}
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDto
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		res := helper.BuildErrorResponse("Failed tp process request", "duplicated email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, res)
	} else {
		createuser := c.authService.CreateUser(registerDTO)
		// token := c.jwtService.GenerateToken(strconv.FormatUint(createuser.ID, 10))
		// createuser.Token = token
		res := helper.BuildResponse(true, "OK", createuser)
		ctx.JSON(http.StatusCreated, res)
	}
}
