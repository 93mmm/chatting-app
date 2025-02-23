package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertExpected(t *testing.T, expected Responce, actual Responce, err error) {
    if assert.NoError(t, err) {
        assert.Equal(t, expected.Code, actual.Code, "Wrong status code")
        assert.Equal(t, expected.Body["id"], actual.Body["id"], "Wrong JSON answer")
        assert.Equal(t, expected.Body["type"], actual.Body["type"], "Wrong JSON answer")
    }
}
