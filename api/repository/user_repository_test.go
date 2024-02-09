package repository

import (
	"testing"

	"github.com/errorsolver/Todo-Golang/api/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var users []model.User = []model.User{
	{
		Name:     "Name1",
		Password: "Pass1",
		Todo:     []model.Todo{},
	},
	{
		Name:     "Name2",
		Password: "Pass2",
		Todo:     []model.Todo{},
	},
	{
		Name:     "Name3",
		Password: "Pass3",
		Todo:     []model.Todo{},
	},
}

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Find(dest interface{}) *MockDB {
	m.Called(dest)
	return m
}

func (m *MockDB) Error() error {
	args := m.Called()
	return args.Error(0)
}

func TestGetAllUsers(t *testing.T) {
	mockDB := new(MockDB)

	// Memberikan expectasi panggilan metode Find
	mockDB.On("Find", &[]model.User{}).Return(mockDB)

	// Tidak ada error yang akan dihasilkan
	mockDB.On("Error").Return(nil)

	repository.DB = mockDB

	expect := users
	actual, err := GetAllUsers()
	if ok := assert.Nil(t, err); !ok {
		t.Logf("Expect no error, but get %d", err)
	}
	if ok := assert.Equal(t, expect, actual); !ok {
		t.Logf("Expect %v, but get %v", expect, actual)

	}
}
