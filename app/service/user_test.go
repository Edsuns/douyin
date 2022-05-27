package service

import (
	"douyin/app/config"
	"douyin/app/dao"
	"douyin/pkg/security"
	"github.com/stretchr/testify/assert"
	"testing"
)

func prepareDao() func() {
	config.Load("../../test/")
	dao.Setup()
	dao.TruncateAllTables()
	return dao.TruncateAllTables
}

func TestRegister(t *testing.T) {
	// prepare
	restore := prepareDao()
	defer restore()

	var (
		user *dao.User
		err  error
	)

	// action register
	expectedUsr := "test"
	expectedPwd := "test_pwd"
	user, err = Register(expectedUsr, expectedPwd)
	// check register
	assert.Nil(t, err)
	assert.True(t, user.ID > 0)
	assert.Equal(t, expectedUsr, user.Username)
	assert.True(t, security.VerifyPassword(expectedPwd, user.Password))

	// action register with empty values
	user, err = Register("", "pwd")
	// check register with empty values
	assert.Nil(t, user)
	assert.NotNil(t, err)
}

func TestLogin(t *testing.T) {
	// prepare
	restore := prepareDao()
	defer restore()

	var (
		usr   = "test"
		pwd   = "test_pwd"
		user  *dao.User
		token *string
		err   error
	)

	// action register
	user, err = Register(usr, pwd)
	// check register
	assert.Nil(t, err)
	assert.NotNil(t, user)

	// action login
	user, token = Login(usr, pwd)
	// check login
	assert.NotNil(t, user)
	assert.NotNil(t, token)
	assert.True(t, user.ID > 0)
	assert.Equal(t, usr, user.Username)
	assert.True(t, security.VerifyPassword(pwd, user.Password))

	// action login with invalid password
	user, token = Login("test", "test_pwd_")
	// check login with invalid password
	assert.Nil(t, user)
	assert.Nil(t, token)

	// action login with invalid username
	user, token = Login("test_", "test_pwd")
	// check login with invalid username
	assert.Nil(t, user)
	assert.Nil(t, token)
}
