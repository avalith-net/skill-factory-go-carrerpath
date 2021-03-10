package database

import (
	"fmt"
	"testing"

	"github.com/avalith-net/skill-factory-go-carrerpath/models"
	"github.com/avalith-net/skill-factory-go-carrerpath/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	registersMock      registerMock
	insertRegisterMock func(user models.User) (string, bool, error)
)

type registerMock struct{}

func init() {
	RegisterService = &registerMock{}
}

func (mock *registerMock) InsertRegister(user models.User) (string, bool, error) {
	return insertRegisterMock(models.User{})
}

func TestInsertRegisterNoError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	var user models.User = models.User{
		Name:     "testing",
		Surname:  "testing",
		Email:    "test@registro.com",
		Password: "Test123",
	}

	//This way we don´t need to access the database
	insertRegisterMock = func(user models.User) (string, bool, error) {
		return "123456", true, nil
	}
	_, status, err := insertRegisterMock(user)
	assert.Nil(t, err)
	assert.EqualValues(t, true, status)
	err = utils.FormValidation(user)
	assert.Nil(t, err)
	assert.EqualValues(t, "testing", user.Name)
	assert.EqualValues(t, "testing", user.Surname)
	assert.EqualValues(t, "test@registro.com", user.Email)
	assert.EqualValues(t, "Test123", user.Password)
}
func TestInsertRegisterWithNameMissing(t *testing.T) {
	gin.SetMode(gin.TestMode)
	var user models.User = models.User{
		Name:     "",
		Surname:  "testing",
		Email:    "test@registro.com",
		Password: "Test123",
	}

	//This way we don´t need to access the database
	insertRegisterMock = func(user models.User) (string, bool, error) {
		return "", false, fmt.Errorf("wrong register")
	}
	_, status, err := insertRegisterMock(user)
	assert.NotNil(t, err)
	assert.EqualValues(t, false, status)
	err = utils.FormValidation(user)
	assert.NotNil(t, err)
	assert.EqualValues(t, "name is needed", err.Error())
	assert.EqualValues(t, "testing", user.Surname)
	assert.EqualValues(t, "test@registro.com", user.Email)
	assert.EqualValues(t, "Test123", user.Password)
}
