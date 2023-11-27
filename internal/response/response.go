package response

import (
	"encoding/json"
	"fmt"
	custom "go-gc-community/pkg/errors"
	"go-gc-community/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

func Error(w gin.ResponseWriter, statusCode int, featureType string, errorNumber string, err error, path string) {
	code := fmt.Sprintf("%d%s%s", statusCode, featureType, errorNumber)
	msg := err.Error()
	bt, err := json.Marshal(Response{
		ResponseCode: code,
		ResponseMessage: msg,
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

	logger.Logger.Error(fmt.Sprintf("Error: %s", msg), zap.String("path", path), zap.Error(err), zap.String("code", code))
	return
}

func Default(w gin.ResponseWriter,statusCode int, featureType string, successMessage string, path string) {
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

	logger.Logger.Info(successMessage, zap.String("path", path))
	return
}

func Success(w gin.ResponseWriter, statusCode int, path string, response interface{}) {
	bt, err := json.Marshal(response)
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
	
	logger.Logger.Info(SUCCESS_DEFAULT, zap.String("path", path))
	return
}