package v1

import (
	"fmt"
	"go-gc-community/internal/models"
	"go-gc-community/internal/response"
	"go-gc-community/pkg/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *V1Handler) internalRoutes(api *gin.RouterGroup) {
	internal := api.Group("/internal", h.Authorize)
	{
		internal.POST("/role", h.Set)
		event := internal.Group("/event")
		{
			register := event.Group("/register")
			{
				register.GET("/list", h.RegistrationList)
				register.POST("/verify", h.Verify)
			}
		}
	}
}

// @Summary Role Set
// @Tags internal-role
// @Description Nanti
// @ModuleID Internal
// @Accept  json
// @Produce  json
// @Success 200 {object} models.SetRoleResponse "Response indicates that the request succeeded and user is logged in"
// @Failure 400 {object} response.Response "There is something wrong with how user input the data"
// @Router api/v1.0/internal/role [post]
func (ih *V1Handler) Set(ctx *gin.Context) {
	var request models.SetRoleRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		//logger.Error(err)
		response.Error(ctx.Writer, http.StatusUnprocessableEntity, "02", "01", errors.DATA_INVALID.Error, ctx.Request.URL.Path)
		return
	}

	user, err := ih.usecase.Internal.SetRole(&request)
	if err != nil {
		//logger.Error(err)
		response.Error(ctx.Writer, http.StatusBadRequest, "02", "02", err, ctx.Request.URL.Path)
		return
	}

	response.Success(ctx.Writer, http.StatusOK, ctx.Request.URL.Path, models.SetRoleResponse{
		ResponseCode: fmt.Sprintf("%d%s%s", http.StatusOK, "00", "00"),
		ResponseMessage: response.SUCCESS_DEFAULT,
		Name: user.Name,
		Email: user.Email,
		AccountNumber: user.AccountNumber,
		RoleId: user.RoleId,
	})
} 

func (ih *V1Handler) RegistrationList(ctx *gin.Context) {
	page := ctx.Query("page")
	limit := ctx.Query("limit")
    sort := ctx.Query("sort")
    filter := ctx.Query("filter")

	accountNumber, ok := ctx.Get("accountNumber")
	if !ok {
		response.Error(ctx.Writer, http.StatusConflict, "02", "03", errors.DATA_INVALID.Error, ctx.Request.URL.Path)
		return
	}

	result, err := ih.usecase.Internal.List(page, limit, sort, filter, accountNumber.(string))
	if err != nil {
		//logger.Error(err)
		response.Error(ctx.Writer, http.StatusBadRequest, "02", "04", err, ctx.Request.URL.Path)
		return
	}

	number, _ := strconv.Atoi(page)
	limits, _ := strconv.Atoi(limit)

	list := make([]models.RegistrationListDetail, len(result))
	for i, p := range result {
		list[i] = models.RegistrationListDetail{
			Name: p.Name,
			Email: p.Email,
			AccountNumber: p.AccountNumber,
			Code: p.Code,
			EventsId: p.EventsId,
			SessionsId: p.SessionsId,
			Status: p.Status,
			BookedBy: p.BookedBy,
		}
	}

	response.Success(ctx.Writer, http.StatusOK, ctx.Request.URL.Path, models.RegistrationListResponse{
		ResponseCode: fmt.Sprintf("%d%s%s", http.StatusOK, "00", "00"),
		ResponseMessage: response.SUCCESS_DEFAULT,
		Number: number,
		Limit: limits,
		Sort: sort,
		Data: list,
	})
}

// @Summary Role Set
// @Tags internal-role
// @Description Nanti
// @ModuleID Internal
// @Accept  json
// @Produce  json
// @Success 200 {object} models.VerifyRegistrationResponse "Response indicates that the request succeeded and user is logged in"
// @Failure 400 {object} response.Response "There is something wrong with how user input the data"
// @Router api/v1.0/internal/register/verify [post]
func (ih *V1Handler) Verify(ctx *gin.Context) {
	var request models.VerifyRegistrationRequest

	accountNumber, ok := ctx.Get("accountNumber")
	if !ok {
		response.Error(ctx.Writer, http.StatusConflict, "02", "05", errors.DATA_INVALID.Error, ctx.Request.URL.Path)
		return
	}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		//logger.Error(err)
		response.Error(ctx.Writer, http.StatusUnprocessableEntity, "02", "06", errors.DATA_INVALID.Error, ctx.Request.URL.Path)
		return
	}

	reg, event, session, err := ih.usecase.Internal.Verify(request.Code, accountNumber.(string))
	if err != nil {
		//logger.Error(err)
		
		response.Error(ctx.Writer, http.StatusBadRequest, "02", "07", err, ctx.Request.URL.Path)
		return
	}

	info := models.EventInformationDetail {
		EventsId: event.ID,
		EventName: event.Name,
		SessionsId: session.ID,
		SessionName: session.Name,
	}

	stats := models.EventStatisticsDetail {
		AvailableSeats: session.AvailableSeats,
		BookedSeats: session.BookedSeats,
		ScannedSeats: session.ScannedSeats,
		UnscannedSeats: session.UnscannedSeats,
		TotalRegistration: session.TotalRegistration,
	}

	response.Success(ctx.Writer, http.StatusOK, ctx.Request.URL.Path, models.VerifyRegistrationResponse{
		ResponseCode: fmt.Sprintf("%d%s%s", http.StatusOK, "00", "00"),
		ResponseMessage: response.SUCCESS_DEFAULT,
		Name: reg.Name,
		Email: reg.Email,
		Status: reg.Status,
		Information: info,
		Statistics: stats,
	})
}