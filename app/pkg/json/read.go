package json_rw

import (
	"encoding/json"
	"os"
)

func Read(fn string) (any, error) {
    data, err := os.ReadFile(fn)
    if err != nil {
        return nil, err
    }

    var json_data any
    err = json.Unmarshal(data, &json_data)
    if err != nil {
        return nil, err
    }

    return json_data, nil
}

