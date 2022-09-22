package usecase

import (
	"errors"
	"project/dashboard/features/mentee"
	"project/dashboard/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostData(t *testing.T) {

	menteeMock := new(mocks.MenteeData)
	input := mentee.MenteeCore{ID: 1, Name: "Amin", Address: "Jakarta", Homeaddress: "Jakarta", Email: "Amin@mail.com", Telegram: "amin", Phone: "08381231212", Gender: "Male", EmergencyData: mentee.EmergencyCore{Name: "saya", Phone: "0812361243", Status: "Ayah"}, Education: mentee.EducationCore{Type: "Informatics", Major: "SMK", Graduate: "SMK 1"}, Class: "BE11", StatusMentee: "Active"}
	token := 1
	t.Run("succes Post", func(t *testing.T) {
		menteeMock.On("InsertData", mock.Anything, mock.Anything).Return(1, nil).Once()

		useCase := New(menteeMock)
		res := useCase.PostData(input, token)
		assert.Equal(t, 1, res)
		menteeMock.AssertExpectations(t)
	})

	t.Run("failed Post", func(t *testing.T) {
		menteeMock.On("InsertData", mock.Anything, mock.Anything).Return(-1, nil).Once()

		idToken := 10
		useCase := New(menteeMock)
		res := useCase.PostData(input, idToken)
		assert.NotEqual(t, 1, res)
		menteeMock.AssertExpectations(t)
	})

	t.Run("failed Post", func(t *testing.T) {

		input.Gender = "male"
		useCase := New(menteeMock)
		res := useCase.PostData(input, token)
		assert.NotEqual(t, 1, res)
		menteeMock.AssertExpectations(t)
	})

	t.Run("failed Post", func(t *testing.T) {

		input.Gender = ""
		useCase := New(menteeMock)
		res := useCase.PostData(input, token)
		assert.NotEqual(t, 1, res)
		menteeMock.AssertExpectations(t)
	})

}

func TestGetMentee(t *testing.T) {

	menteeMock := new(mocks.MenteeData)
	returnData := []mentee.Join{{ID: 1, Name: "Amin", Class: "BE11", Status: "Active", Gender: "Male", Category: "Non-Informatics"}}
	page := 1

	t.Run("Succes GetData", func(t *testing.T) {
		menteeMock.On("SelectData", mock.Anything, mock.Anything).Return(returnData, nil).Once()

		useCase := New(menteeMock)
		res, err := useCase.GetData(page, 1)
		assert.Equal(t, res, returnData)
		assert.NoError(t, err)
		menteeMock.AssertExpectations(t)
	})

	t.Run("Failed GetData", func(t *testing.T) {
		menteeMock.On("SelectData", mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()

		useCase := New(menteeMock)
		res, err := useCase.GetData(page, 10)
		assert.Nil(t, res)
		assert.NotNil(t, err)
		menteeMock.AssertExpectations(t)
	})

}

func TestPutMentee(t *testing.T) {

	menteeMock := new(mocks.MenteeData)
	input := mentee.MenteeCore{StatusMentee: "Active"}

	t.Run("Succes UpdateData", func(t *testing.T) {
		menteeMock.On("UpdateData", mock.Anything, mock.Anything).Return("Sukses Mengupdate Semua Data", nil).Once()

		useCase := New(menteeMock)
		res, err := useCase.UpdateData(input, 1)
		assert.Equal(t, res, "Sukses Mengupdate Semua Data")
		assert.NoError(t, err)
		menteeMock.AssertExpectations(t)
	})

	t.Run("Failed UpdateData", func(t *testing.T) {
		menteeMock.On("UpdateData", mock.Anything, mock.Anything).Return("Gagal Update Pada Education", errors.New("error")).Once()

		useCase := New(menteeMock)
		_, err := useCase.UpdateData(input, 10)
		assert.NotNil(t, err)
		menteeMock.AssertExpectations(t)
	})

}
