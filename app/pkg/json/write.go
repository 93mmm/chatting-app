package json_rw

import (
	"encoding/json"
	"os"
)

func Write(fn string, d any) error {
    data, err := json.Marshal(d)
    if err != nil {
        return err
    }

    err = os.WriteFile(fn, data, 0666)
    return err
}
