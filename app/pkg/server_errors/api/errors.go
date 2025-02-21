package api

type ApiError struct {
    Code    int     `json:"code"`
    Msg     string  `json:"msg"`
    Details map[string]string `json:"details,omitempty"`
}

func (this ApiError) Error() string {
    return "some api error happened"
}

type DbError struct {
    Code    int     `json:"code"`
    Msg     string  `json:"msg"`
    Details map[string]string `json:"details,omitempty"`
}

func (this DbError) Error() string {
    return "some db error happened"
}

func NewApiError(code int, details map[string]string) error {
    return ApiError{
        Code: code,
        Msg: "ErrorHappened",
        Details: details,
    }
}
