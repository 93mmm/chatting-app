package server_errors

type apiErrorType string

// Database
const (
    ERR_DATABASE_UNKNOWN    apiErrorType = "UnknownDbError"

    ERR_USER_NOT_FOUND      apiErrorType = "UserNotFound"
    ERR_USER_ALREADY_EXISTS apiErrorType = "UserAlreadyExists"

    ERR_CHAT_NOT_FOUND      apiErrorType = "ChatNotFound"
)

// Request Errors
const (
    ERR_JSON_KV             apiErrorType = "IncorrectJsonField"
    ERR_JSON_CORRUPTED      apiErrorType = "CorruptedJson"
    
    ERR_WRONG_VALUE         apiErrorType = "WrongValue"
)

// Undefined
const (
    ERR_UNDEFINED           apiErrorType = "UndefinedError"
)
