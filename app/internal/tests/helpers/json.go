package helpers

import (
	"encoding/json"
	"log"
	"os"
)

func ReadJsonTests(path string) []RRTest {
    var r []RRTest

    data, err := os.ReadFile(GetTestJsonPath(path))
    if err != nil {
        log.Fatal(err)
    }

    err = json.Unmarshal(data, &r)
    if err != nil {
        log.Fatal(err)
    }
    return r
}
