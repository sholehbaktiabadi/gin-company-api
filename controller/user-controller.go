package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sholehbaktiabadi/go-api/dto"
	"github.com/sholehbaktiabadi/go-api/helper"
	"github.com/sholehbaktiabadi/go-api/service"
)

type UserController interface {
	Update(ctx *gin.Context)
	Profile(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (controller *userController) Update(context *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO

	errDTO := context.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to proccess request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := controller.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	println(id)
	if err != nil {
		panic(err.Error())
	}
	userUpdateDTO.ID = id
	user := controller.userService.Update(userUpdateDTO)
	res := helper.BuildResponse(true, "OK", user)
	context.JSON(http.StatusOK, res)
}

func (controller *userController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := controller.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	user := controller.userService.Profile(fmt.Sprintf("%v", claims["user_id"]))
	res := helper.BuildResponse(true, "OK", user)
	context.JSON(http.StatusOK, res)
}
