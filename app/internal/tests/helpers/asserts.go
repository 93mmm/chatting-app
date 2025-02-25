package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertExpected(t *testing.T, expected Responce, actual Responce, err error, fields []string) {
    if assert.NoError(t, err) {
        assert.Equal(t, expected.Code, actual.Code, "Wrong status code")

        for _, f := range fields {
            assert.Equal(t, expected.Body[f], actual.Body[f], fmt.Sprintf("Wrong data on field %v", f))
        }
        assert.Equal(t, expected.Body["type"], actual.Body["type"], "Wrong error type")
    }
}
