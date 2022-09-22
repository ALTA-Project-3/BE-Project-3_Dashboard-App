package usecase

import (
	"errors"
	"project/dashboard/mocks"
	"testing"

	"project/dashboard/features/user"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {

	userMock := new(mocks.UserData)
	returnData := []user.Core{{Fullname: "amin", ID: 1, Email: "amin@a.com", Password: "123123", Team: "Academic", Role: "Default", Status: "Active"}}
	token := 1
	page := 1
	t.Run("Get All Success", func(t *testing.T) {

		userMock.On("SelectAll", token, page).Return(returnData, nil).Once()

		useCase := New(userMock)
		res, err := useCase.GetAll(page, token)
		assert.NoError(t, err)
		assert.Equal(t, returnData[0].ID, res[0].ID)
		userMock.AssertExpectations(t)

	})

	t.Run("Get All Failed", func(t *testing.T) {
		userMock.On("SelectAll", token, page).Return(nil, nil).Once()

		useCase := New(userMock)
		res, _ := useCase.GetAll(page, token)
		assert.Equal(t, 0, len(res))
		userMock.AssertExpectations(t)

	})

	t.Run("Get All Failed", func(t *testing.T) {
		userMock.On("SelectAll", token, page).Return(nil, errors.New("error")).Once()

		useCase := New(userMock)
		_, err := useCase.GetAll(page, token)
		assert.Error(t, err)
		userMock.AssertExpectations(t)

	})

}

func TestGetProfile(t *testing.T) {

	userMock := new(mocks.UserData)
	returnData := user.Core{Fullname: "amin", ID: 1, Email: "amin@a.com", Password: "123123", Team: "Academic", Role: "Default", Status: "Active"}
	returnData2 := user.DashBoard{Active: 10, Placement: 10, FeedBack: 10, ActiveInMonth: []user.MonthActive{}, GraduateInMonth: []user.MonthGraduate{}}

	token := 1

	t.Run("Get profile success", func(t *testing.T) {
		userMock.On("SelectProfile", token).Return(returnData, returnData2, nil).Once()

		useCase := New(userMock)
		res, res2, _ := useCase.GetProfile(token)
		assert.Equal(t, token, int(res.ID))
		assert.Equal(t, returnData2, res2)
		userMock.AssertExpectations(t)

	})

	t.Run("Get profile failed", func(t *testing.T) {

		userMock.On("SelectProfile", token).Return(user.Core{}, user.DashBoard{}, errors.New("error")).Once()

		useCase := New(userMock)
		param := 1
		res, _, err := useCase.GetProfile(token)
		assert.Error(t, err)
		assert.NotEqual(t, param, int(res.ID))
		userMock.AssertExpectations(t)

	})

}

func TestPut(t *testing.T) {

	userMock := new(mocks.UserData)
	input := user.Core{Fullname: "amin", ID: 1, Email: "amin@a.com", Team: "Academic", Role: "Default", Status: "Active"}

	t.Run("update succes", func(t *testing.T) {

		userMock.On("UpdateData", input).Return(1, nil).Once()

		useCase := New(userMock)
		res := useCase.PutData(input)
		assert.Equal(t, 1, res)
		userMock.AssertExpectations(t)

	})

	t.Run("update failed", func(t *testing.T) {

		userMock.On("UpdateData", input).Return(-1, errors.New("error")).Once()

		useCase := New(userMock)
		res := useCase.PutData(input)
		assert.Equal(t, -1, res)
		userMock.AssertExpectations(t)

	})

}

func TestPostData(t *testing.T) {

	userMock := new(mocks.UserData)
	input := user.Core{Fullname: "amin", ID: 1, Email: "amin@a.com", Password: "123123", Team: "Academic", Role: "Default", Status: "Active"}

	t.Run("create success", func(t *testing.T) {

		userMock.On("InsertData", mock.Anything).Return(1, nil).Once()

		useCase := New(userMock)
		res := useCase.PostData(input)
		assert.Equal(t, 1, res)
		userMock.AssertExpectations(t)
	})

	t.Run("create failed", func(t *testing.T) {

		userMock.On("InsertData", mock.Anything).Return(-1, errors.New("error")).Once()

		useCase := New(userMock)
		res := useCase.PostData(input)
		assert.Equal(t, -1, res)
		userMock.AssertExpectations(t)

	})

	t.Run("create failed", func(t *testing.T) {

		input.Fullname = ""
		input.Password = ""
		useCase := New(userMock)
		res := useCase.PostData(input)
		assert.Equal(t, -1, res)
		userMock.AssertExpectations(t)

	})

}

func TestDelete(t *testing.T) {

	userMock := new(mocks.UserData)
	param := 1

	t.Run("delete succes", func(t *testing.T) {

		userMock.On("DelData", param).Return(1, nil).Once()

		useCase := New(userMock)
		res := useCase.DeleteData(param)
		assert.Equal(t, 1, res)
		userMock.AssertExpectations(t)

	})

	t.Run("delete failed", func(t *testing.T) {

		userMock.On("DelData", param).Return(-1, errors.New("error")).Once()

		useCase := New(userMock)
		res := useCase.DeleteData(param)
		assert.Equal(t, -1, res)
		userMock.AssertExpectations(t)

	})

}
