package v1

import (
	"errors"
	"go-gc-community/internal/response"
	custom "go-gc-community/pkg/errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Decoded struct {
	UserId	int
	AccountNumber string
}

func (h *V1Handler) Retrieve(ctx *gin.Context) (*jwt.Token, error) {
	header := ctx.GetHeader("Authorization")
	if header == "" {
		return nil, errors.New("Empty Header")
	}

	token := strings.Split(header, " ")
	if len(token) != 2 || token[0] != "Bearer" {
		return nil, errors.New("Invalid Authorization")
	}

	if len(token[1]) == 0 {
		return nil, errors.New("Empty Authorization")
	}

	return h.authorization.Validate(token[1])
}

func (h *V1Handler) Get(encoded *jwt.Token) (Decoded, error) {
	var content Decoded
	decoded, ok := encoded.Claims.(jwt.MapClaims)
	if !ok && !encoded.Valid {
		return content, errors.New("invalid token")
	}

	content.UserId = int(decoded["id"].(float64))
	content.AccountNumber = decoded["accountNumber"].(string)

	return content, nil
}

func Token(ctx *gin.Context) string {
	token := ctx.Query("token")
	if token != "" {
		return token
	}

	bearerToken := ctx.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func (h *V1Handler) Authorize(ctx *gin.Context) {
	retrieved, err := h.Retrieve(ctx)
	if err != nil {
		//logger.Error(err)
		response.Error(ctx.Writer, http.StatusUnauthorized, "00", "01", custom.UNAUTHORIZED.Error, ctx.Request.URL.Path)
	}

	content, err := h.Get(retrieved)
	if err != nil {
		//logger.Error(err)
		response.Error(ctx.Writer, http.StatusUnauthorized, "00", "02", custom.UNAUTHORIZED.Error, ctx.Request.URL.Path)
	}

	ctx.Set("userId", content.UserId)
	ctx.Set("accountNumber", content.AccountNumber)
}