package v1

import (
	"fmt"
	"go-gc-community/internal/models"
	"go-gc-community/internal/response"
	"go-gc-community/pkg/errors"
	"go-gc-community/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *V1Handler) userRoutes(api *gin.RouterGroup) {
	user := api.Group("/user")
	{
		/*google := user.Group("/google")
		{
			google.GET("/login", h.Login)
			google.GET("callback", h.Callback)
		}*/
		user.GET("/login", h.Login)
		user.GET("/callback", h.Callback)

		authorized := user.Group("/", h.Authorize)
		{
			authorized.POST("/inquiry", h.Inquiry)
			authorized.GET("/inquire", h.Inquire)
		}
	}
}

// @Summary Google Login
// @Tags user-login
// @Description This is the endpoint to redirect user to Google Oauth Consent Screen
// @ModuleID User
// @Accept  json
// @Produce  json
// @Success 307 {object} "Response indicates that the request succeeded and the resources has been fetched and transmitted in the message body"
// @Router api/v1.0/user/login [get] 
func (uh *V1Handler) Login(ctx *gin.Context) {
	url := uh.usecase.User.Redirect()
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}

// @Summary Google Callback
// @Tags user-callback
// @Description This is the endpoint register and login user to the database via Google API
// @ModuleID User
// @Accept  json
// @Produce  json
// @Success 201 {object} models.UserLoginResponse "Response indicates that the request succeeded and user account is created"
// @Success 200 {object} models.UserLoginResponse "Response indicates that the request succeeded and user is logged in"
// @Failure 422 {object} response.Response "There is something wrong with how user input the data"
// @Router api/v1.0/user/callback [get] 
func (uh *V1Handler) Callback(ctx *gin.Context) {
	code := ctx.Query("code")
	state := ctx.Query("state")
	user, appToken, statusCode, err := uh.usecase.User.Account(state, code)
	if err != nil {
		logger.Error(err.Error())
		response.Error(ctx.Writer, http.StatusUnprocessableEntity, "00", "01", errors.UNAUTHORIZED.Error)
	}

	if statusCode == 201 {
		ctx.JSON(http.StatusCreated, models.UserLoginResponse{
			ResponseCode: fmt.Sprintf("%d%s%s", http.StatusCreated, "00", "00"),
			ResponseMessage: "Response has been successfully proceeded.",
			AccountNumber: user.AccountNumber,
			UserID: user.ID,
			Token: appToken,
		})
		return
	}


	ctx.JSON(http.StatusOK, models.UserLoginResponse{
		ResponseCode: fmt.Sprintf("%d%s%s", http.StatusOK, "00", "00"),
		ResponseMessage: "Response has been successfully proceeded.",
		AccountNumber: user.AccountNumber,
		UserID: user.ID,
		Token: appToken,
	})
}

// @Summary Inquire User
// @Tags user-inquire
// @Description This is the endpoint retrieve user data
// @ModuleID User
// @Accept  json
// @Produce  json
// @Success 201 {object} models.UserLoginResponse "Response indicates that the request succeeded and user account is created"
// @Success 200 {object} models.UserLoginResponse "Response indicates that the request succeeded and user is logged in"
// @Failure 422 {object} response.Response "There is something wrong with how user input the data"
// @Router api/v1.0/user/inquiry [get] 
func (uh *V1Handler) Inquiry(ctx *gin.Context) {
	var request models.InquiryUserRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		logger.Error(err)
		response.Error(ctx.Writer, http.StatusUnprocessableEntity, "00", "02", errors.DATA_INVALID.Error)
	}

	userData, err := uh.usecase.User.Inquire(&request)
	if err != nil {
		logger.Error(err)
		response.Error(ctx.Writer, http.StatusBadRequest, "00", "03", err)
	}

	response.Success(ctx.Writer, http.StatusOK, models.InquiryUserResponse{
		ResponseCode: fmt.Sprintf("%d%s%s", http.StatusOK, "00", "00"),
		ResponseMessage: response.SUCCESS_DEFAULT,
		AccountNumber: userData.AccountNumber,
		Name: userData.Name,
		State: userData.State,
		Role: userData.RoleId,
		Email: userData.Email,
	})
}

func (uh *V1Handler) Inquire(ctx *gin.Context) {
	accountNumber, ok := ctx.Get("accountNumber")
	if !ok {
		response.Error(ctx.Writer, http.StatusUnprocessableEntity, "00", "04", errors.DATA_INVALID.Error)
	}
	
	var request models.InquiryUserRequest
	request.AccountNumber = accountNumber.(string)

	userData, err := uh.usecase.User.Inquire(&request)
	if err != nil {
		logger.Error(err)
		response.Error(ctx.Writer, http.StatusBadRequest, "00", "05", err)
	}

	response.Success(ctx.Writer, http.StatusOK, models.InquiryUserResponse{
		ResponseCode: fmt.Sprintf("%d%s%s", http.StatusOK, "00", "00"),
		ResponseMessage: "Response has been successfully proceeded.",
		AccountNumber: userData.AccountNumber,
		Name: userData.Name,
		State: userData.State,
		Role: userData.RoleId,
		Email: userData.Email,
	})
}