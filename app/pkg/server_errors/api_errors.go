package server_errors

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiError struct {
    HttpCode    int             `json:"code"`
    Type        apiErrorType    `json:"type"`
    Context     map[string]any  `json:"context"`
}

func (this *ApiError) Error() string {
    return fmt.Sprintf("%v error with details:\n\t%v", this.Type, this.Context)
}

func (this *ApiError) Send(c *gin.Context) {
    c.JSON(this.HttpCode, this)
}

func newApiError(code int, name apiErrorType, d map[string]any) *ApiError {
    return &ApiError{code, name, d}
}



// ========================================================================
// |                           ApiError factory                           |
// ========================================================================

// If err is an ApiError, then only casting is performed, in the other case, casting err to ApiError
func UndefinedError(err error) *ApiError {
    var apiErr *ApiError
    if errors.As(err, &apiErr) {
        return apiErr
    }
    return newApiError(http.StatusBadGateway, ERR_UNDEFINED, map[string]any{"error": err.Error()})
}

func UserAlreadyExists(uname *string, email string) *ApiError {
    return newApiError(http.StatusBadRequest, ERR_USER_ALREADY_EXISTS, map[string]any{
        "username": uname,
        "email": email,
    })
}

func UserNotFoundError(id int) *ApiError {
    return newApiError(http.StatusBadRequest, ERR_USER_NOT_FOUND, map[string]any{"id": id})
}

func JsonCorruptedError() *ApiError {
    return newApiError(http.StatusBadRequest, ERR_JSON_CORRUPTED, nil)
}

func WrongValueError(val any) *ApiError {
    return newApiError(http.StatusBadRequest, ERR_WRONG_VALUE, map[string]any{"value": val})
}

func WrongKVError(key string, value any) *ApiError {
    return newApiError(http.StatusBadRequest, ERR_JSON_KV, map[string]any{key: value})
}
