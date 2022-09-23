package usecase

import (
	"errors"
	"project/dashboard/features/class"
	"project/dashboard/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostClass(t *testing.T) {
	classMock := new(mocks.ClassData)
	var layout2 = "2006-01-02"
	StartDate := "2022-08-01"
	EndDate := "2022-12-01"
	start, _ := time.Parse(layout2, StartDate)
	end, _ := time.Parse(layout2, EndDate)
	mocks := class.CoreClass{ID: 1, Name: "BE11", StartDate: start, EndDate: end, UserID: 2}

	t.Run("Yang Sukses", func(t *testing.T) {
		classMock.On("InsertClass", mock.Anything).Return("a string", nil).Once()
		classlogic := New(classMock)
		msg, err := classlogic.CreateClass(mocks)
		assert.Equal(t, msg, "a string")
		assert.NoError(t, err)
		classMock.AssertExpectations(t)
	})
	t.Run("Yang Sukses", func(t *testing.T) {
		classMock.On("InsertClass", mock.Anything).Return("Terjadi Kesalahan", errors.New("An Error")).Once()
		classlogic := New(classMock)
		msg, err := classlogic.CreateClass(mocks)
		assert.Equal(t, msg, "Terjadi Kesalahan")
		assert.Error(t, err)
		classMock.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	classMock := new(mocks.ClassData)
	var layout2 = "2006-01-02"
	StartDate := "2022-08-01"
	EndDate := "2022-12-01"
	start, _ := time.Parse(layout2, StartDate)
	end, _ := time.Parse(layout2, EndDate)
	mocks := []class.CoreClass{{ID: 1, Name: "BE11", StartDate: start, EndDate: end, UserID: 2}}

	t.Run("Yang Sukses", func(t *testing.T) {
		classMock.On("SelectAllClass", mock.Anything).Return(mocks, nil).Once()
		classlogic := New(classMock)
		sli, err := classlogic.GetAllClass(0)
		assert.Equal(t, sli[0], mocks[0])
		assert.NoError(t, err)
		classMock.AssertExpectations(t)
	})
	t.Run("Yang Gagal", func(t *testing.T) {
		classMock.On("SelectAllClass", mock.Anything).Return(nil, errors.New("An Error")).Once()
		classlogic := New(classMock)
		sli, err := classlogic.GetAllClass(0)
		assert.Nil(t, sli)
		assert.Error(t, err)
		assert.NotNil(t, err)
		classMock.AssertExpectations(t)
	})
}
