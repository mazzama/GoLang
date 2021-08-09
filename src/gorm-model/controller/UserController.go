package controller

import (
	"net/http"

	dto "gorm-model/model/dto"
	"gorm-model/service"

	"github.com/bytedance/go-tagexpr/validator"
	"github.com/gin-gonic/gin"
)

// Login method:POST, endpoint:/login
func Login(context *gin.Context) {
	var auth dto.Auth

	if err := context.ShouldBindJSON(&auth); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "Invalid json provided"})
		return
	}

	service.Login(context, auth)
}

// CreateUser method:POST, endpoint:/user
func CreateUser(context *gin.Context) {
	// Validate input
	var input dto.Request
	var vd = validator.New("vd")

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  http.StatusUnprocessableEntity,
			"message": "Invalid json provided"})
		return
	}

	type RequesrValidation struct {
		UserName    string `vd:"len($)>3"`
		CompanyName string `vd:"len($)>3"`
	}

	requesrValidation := &RequesrValidation{UserName: input.UserName, CompanyName: input.CompanyName}

	if err := vd.Validate(requesrValidation); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error()})
		return
	}

	service.CreateUserService(context, input)
}

// GetUsers method:GET, endpoint:/users
func GetUsers(context *gin.Context) {
	service.GetUsersService(context)
}

// GetCompanies method:GET, endpoint:/companies
func GetCompanies(context *gin.Context) {
	service.GetCompaniesService(context)
}

// GetUserCompanies method:GET, endpoint:/user-companies
func GetUserCompanies(context *gin.Context) {
	service.GetUserCompaniesService(context)
}

// GetUserCompany method:GET, endpoint:/user-company/:id
func GetUserCompany(context *gin.Context) {
	service.GetUserCompanyService(context)
}
