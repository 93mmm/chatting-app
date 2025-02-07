package json_handlers

import (
	"encoding/json"
	"os"
)

func ReadJSON(name string) (interface{}, error) {
    data, err := os.ReadFile(name)
    if err != nil {
        return nil, err
    }

    var json_data interface{}
    err = json.Unmarshal(data, &json_data)
    if err != nil {
        return nil, err
    }

    return json_data, nil
}
