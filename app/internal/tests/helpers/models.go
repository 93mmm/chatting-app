package helpers

type Request struct {
    Url         string          `json:"url"`
    Method      string          `json:"method"`
    ContentType string          `json:"content_type"`
    Body        map[string]any  `json:"body"`
}

type Responce struct {
    Code        int             `json:"code"`
    Body        map[string]any  `json:"body"`
}

type RRTest struct {
    Name        string          `json:"name"`
    Sent        Request         `json:"sent"`
    Expected    Responce        `json:"expected"`
}
