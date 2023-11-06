package response

import (
	"encoding/json"
	"fmt"
	custom "go-gc-community/pkg/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	ResponseCode	string	`json:"responseCode"`
	ResponseMessage	string	`json:"responseMessage"`
}

const (
	CONTENT_TYPE = "Content-Type"
	APPLICATION_JSON = "application/json"
	SUCCESS_DEFAULT = "Response has been successfully proceed."
)

func Error(w gin.ResponseWriter,statusCode int, featureType string, errorNumber string, err error) {
	bt, err := json.Marshal(Response{
		ResponseCode: fmt.Sprintf("%d%s%s", statusCode, featureType, errorNumber),
		ResponseMessage: err.Error(),
	})
	if err != nil {
		bt, _ = json.Marshal(Response{
			ResponseCode: "5000000",
			ResponseMessage: custom.INTERNAL_SERVER_ERROR.Message,
		})
		
		w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(bt)
		return
	}

	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.WriteHeader(statusCode)
	w.Write(bt)
	return
}

func Success(w gin.ResponseWriter,statusCode int, featureType string, successMessage string) {
	bt, err := json.Marshal(Response{
		ResponseCode: fmt.Sprintf("%d%s00", statusCode, featureType),
		ResponseMessage: successMessage,
	})
	if err != nil {
		bt, _ = json.Marshal(Response{
			ResponseCode: "5000000",
			ResponseMessage: custom.INTERNAL_SERVER_ERROR.Message,
		})
		
		w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(bt)
		return
	}

	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.WriteHeader(statusCode)
	w.Write(bt)
	return
}