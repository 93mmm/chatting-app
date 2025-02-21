package helpers

type Request struct {
    Url         string      `json:"url"`
    Method      string      `json:"method"`
    ContentType string      `json:"content_type"`
    Body        any         `json:"body"`
}

type Responce struct {
    Code        int         `json:"code"`
    Body        any         `json:"body"`
}

type RRTest struct {
    Sent        Request      `json:"sent"`
    Expected    Responce     `json:"expected"`
}
