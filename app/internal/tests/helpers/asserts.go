package helpers

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertExpected(t *testing.T, expected Responce, actual Responce, err error) {
    if assert.NoError(t, err) {
        assert.Equal(t, expected.Code, actual.Code, "Wrong status code")

        data, err := json.Marshal(expected.Body)
        assert.NoError(t, err)
        expected.Body = string(data)
        assert.Equal(t, expected.Body, actual.Body, "Wrong JSON answer")
    }
}
