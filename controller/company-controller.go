package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sholehbaktiabadi/go-api/dto"
	"github.com/sholehbaktiabadi/go-api/entity"
	"github.com/sholehbaktiabadi/go-api/helper"
	"github.com/sholehbaktiabadi/go-api/service"
)

type Companycontroller interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type companyController struct {
	companyService service.CompanyService
	jwtService     service.JWTService
}

func NewCompanyController(companyService service.CompanyService, jwtService service.JWTService) Companycontroller {
	return &companyController{
		companyService: companyService,
		jwtService:     jwtService,
	}
}

func (controller *companyController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("id params not found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var company entity.Company = controller.companyService.FindByID(id)
	if (company == entity.Company{}) {
		res := helper.BuildErrorResponse("data not found", "no data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", company)
		context.JSON(http.StatusOK, res)
	}
}

func (controller *companyController) All(context *gin.Context) {
	var company []entity.Company = controller.companyService.AllCompany()
	res := helper.BuildResponse(true, "all books", company)
	context.JSON(http.StatusOK, res)
}

func (controller *companyController) Insert(context *gin.Context) {
	var companyCreateDTO dto.CompanyCreateDTO
	errDTO := context.ShouldBind(&companyCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to proccess request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		auth := context.GetHeader("Authorization")
		userID := controller.getUserIdByToken(auth)
		convertAuthToID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			companyCreateDTO.UserID = convertAuthToID
		}
		result := controller.companyService.Insert(companyCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (controller *companyController) Update(context *gin.Context) {
	var companyUpdateDTO dto.CompanyUpdateDTO
	errDTO := context.ShouldBind(&companyUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	auth := context.GetHeader("Authorization")
	token, errToken := controller.jwtService.ValidateToken(auth)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	fmt.Println(userID)
	fmt.Println(companyUpdateDTO.ID)
	if controller.companyService.IsAllowedToEdit(userID, companyUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			companyUpdateDTO.UserID = id
		}
		result := controller.companyService.Update(companyUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (controller *companyController) Delete(context *gin.Context) {
	var company entity.Company
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "no params id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	company.ID = id
	auth := context.GetHeader("Authorization")
	token, errToken := controller.jwtService.ValidateToken(auth)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if controller.companyService.IsAllowedToEdit(userID, company.ID) {
		controller.companyService.Delete(company)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		res := helper.BuildErrorResponse("you dont have permission", "youre not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, res)
	}
}

func (controller *companyController) getUserIdByToken(token string) string {
	inputToken, err := controller.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := inputToken.Claims.(jwt.MapClaims)
	return fmt.Sprintf("%v", claims["user_id"])
}
