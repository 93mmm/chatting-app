package helpers

type Request struct {
    Url         string
    JsonBody    string
}

type Responce struct {
    Code        int
    JsonBody    string
}

type ExpectedResponceToRequest struct {
    Sent        Request
    Expected    Responce
}
