package teststore_test

import (
	"testing"

	"github.com/ekimovvd/http-rest-api/internal/app/model"
	"github.com/ekimovvd/http-rest-api/internal/app/store"
	"github.com/ekimovvd/http-rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

// TestUserRepository_Create
func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
}

// TestUserRepository_Find ...
func TestUserRepository_Find(t *testing.T) {
	s := teststore.New()

	u1 := model.TestUser(t)
	s.User().Create(u1)
	u2, err := s.User().Find(u1.ID)

	assert.NoError(t, err)
	assert.NotNil(t, u2)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
