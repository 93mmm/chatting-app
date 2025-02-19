package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func AssertExpected(t *testing.T, exp Responce, res Responce, err error) {
    if assert.NoError(t, err) {
        assert.Equal(t, exp.Code, res.Code, "Wrong status code")
        assert.Equal(t, exp.JsonBody, res.JsonBody, "Wrong JSON answer")
    }
}
