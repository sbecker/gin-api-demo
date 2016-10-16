package models

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	u := NewUser()
	assert.Equal(t, "models.User", reflect.TypeOf(u).String())
	assert.Equal(t, "", u.Email, "email should be blank")
}
