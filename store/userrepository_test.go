package store_test

import (
	"github.com/OldGattsu/go-rest-api/app/model"
	"github.com/OldGattsu/go-rest-api/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "user@example.org",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
