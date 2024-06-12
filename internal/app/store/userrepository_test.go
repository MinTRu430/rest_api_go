package store_test

import (
	"rest_api_go/internal/app/model"
	"rest_api_go/internal/app/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "user@exapmle.org",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "user@exapmle.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	u, err := s.User().Create(&model.User{
		Email: "user@exapmle.org",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
