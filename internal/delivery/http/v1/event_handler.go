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

func (h *V1Handler) eventRoutes(api *gin.RouterGroup) {
	event := api.Group("/event", h.Authorize)
	{
		event.GET("/list", h.List)
		event.GET("/:id/session", h.SessionList)
		event.POST("/register", h.Register)
	}
}

// @Summary Event List
// @Tags event-list
// @Description This is the endpoint retrieve event list
// @ModuleID Event
// @Accept  json
// @Produce  json
// @Success 200 {object} models.GetEventResponse "Response indicates that the request succeeded and user is logged in"
// @Failure 400 {object} response.Response "There is something wrong with how user input the data"
// @Router api/v1.0/user/callback [get] 
func (eh *V1Handler) List(ctx *gin.Context) {
	event, time, err := eh.usecase.Event.Events()
	if err != nil {
		logger.Error(err)
		response.Error(ctx.Writer, http.StatusBadRequest, "01", "01", err)
		return
	}

	count := len(event)

	list := make([]models.EventResponseDetail, len(event))
	for i, p := range event {
		list[i] = models.EventResponseDetail{
			EventID: p.ID,
			EventName: p.Name,
			EventDescription: p.Description,
			EventCode: p.Code,
			OpenRegistration: p.OpenRegistration,
			ClosedRegistration: p.ClosedRegistration,
			Status: p.Status,
		}
	}

	response.Success(ctx.Writer, http.StatusOK, models.GetEventResponse{
		ResponseCode: fmt.Sprintf("%d%s%s", http.StatusOK, "00", "00"),
		ResponseMessage: response.SUCCESS_DEFAULT,
		EventCount: count,
		CurrentTime: time,
		Events: list,
	})
}

// @Summary Session List
// @Tags session-list
// @Description This is the endpoint retrieve event list
// @ModuleID Event
// @Accept  json
// @Produce  json
// @Success 200 {object} models.GetSessionResponse "Response indicates that the request succeeded and user is logged in"
// @Failure 400 {object} response.Response "There is something wrong with how user input the data"
// @Router api/v1.0/event/:id/list [get]
func (eh *V1Handler) SessionList(ctx *gin.Context) {
	eventId := ctx.Param("id")

	session, event, time, err := eh.usecase.Event.Sessions(eventId)
	if err != nil {
		logger.Error(err)
		response.Error(ctx.Writer, http.StatusBadRequest, "01", "02", err)
		return
	}

	count := len(session)
	
	list := make([]models.SessionResponseDetail, len(session))
	for i, p := range session {
		list[i] = models.SessionResponseDetail{
			Name: p.Name,
			SessionID: p.ID,
			EventID: p.EventsId,
			Description: p.Description,
			Time: p.Time,
			MaxSeating: p.MaxSeating,
			Status: p.Status,
			OpenRegistration: p.OpenRegistration,
			ClosedRegistration: p.ClosedRegistration,
		}
	}

	response.Success(ctx.Writer, http.StatusOK, models.GetSessionResponse{
		ResponseCode: fmt.Sprintf("%d%s%s", http.StatusOK, "00", "00"),
		ResponseMessage: response.SUCCESS_DEFAULT,
		EventName: event.Name,
		EventId: event.ID,
		SessionCount: count,
		CurrentTime: time,
		Events: list,
	})
}

// @Summary Registration
// @Tags event-registration
// @Description This is the endpoint retrieve event list
// @ModuleID Event
// @Accept  json
// @Produce  json
// @Success 201 {object} models.RegistrationResponse "Response indicates that the request succeeded and user is logged in"
// @Failure 422 {object} response.Response "There is something wrong with how user input the data"
// @Failure 400 {object} response.Response "There is something wrong with how user input the data"
// @Router api/v1.0/event/register [post]
func (eh *V1Handler) Register(ctx *gin.Context) {
	var request models.RegistrationRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		logger.Error(err)
		response.Error(ctx.Writer, http.StatusUnprocessableEntity, "01", "03", errors.DATA_INVALID.Error)
		return
	}

	main, second, isValid, count, err := eh.usecase.Event.Register(&request)
	if err != nil {
		logger.Error(err)
		response.Error(ctx.Writer, http.StatusBadRequest, "01", "04", err)
		return
	}

	event, err := eh.usecase.Event.Event(request.EventID)
	if err != nil {
		logger.Error(err)
		response.Error(ctx.Writer, http.StatusBadRequest, "01", "05", err)
		return
	}

	list := make([]models.OtherBookingResponse, len(second))
	for i, p := range second {
		list[i] = models.OtherBookingResponse{
			Email: p.Email,
			Name: p.Name,
			BookingCode: p.Code,
		}
	}

	response.Success(ctx.Writer, http.StatusCreated, models.RegistrationResponse{
		ResponseCode: fmt.Sprintf("%d%s%s", http.StatusOK, "00", "00"),
		ResponseMessage: response.SUCCESS_DEFAULT,
		EventCode: event.Code,
		IsValid: isValid,
		SessionID: main.SessionsId,
		EventID: event.ID,
		RequestedSeats: count,
		MainEmail: main.Email,
		MainName: main.Name,
		MainAccountNumber: main.AccountNumber,
		MainBookingCode: main.Code,
		Others: list,
	})
}