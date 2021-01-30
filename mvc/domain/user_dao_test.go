package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUser(t *testing.T) {
	user, err := GetUser(0)
	assert.Nil(t, user, "we are not expecting a user with id 0")
	assert.NotNil(t, err, "expecting error when user id is 0")
	assert.Equal(t, http.StatusNotFound, err.StatusCode)
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.ID)
	assert.EqualValues(t, "pradeepdoit@gmail.com", user.Email)
}
