package usecase

import (
	"project/dashboard/features/user"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userData user.DataInterface
}

func New(data user.DataInterface) user.UsecaseInterface {
	return &userUsecase{
		userData: data,
	}
}

func (usecase *userUsecase) GetAll(page, token int) ([]user.Core, error) {

	data, err := usecase.userData.SelectAll(page, token)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (usecase *userUsecase) PutData(data user.Core) int {

	if data.Password != "" {
		hash, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		data.Password = string(hash)
	}

	row := usecase.userData.UpdateData(data)
	if row == -1 {
		return -1
	}

	return row

}

func (usecase *userUsecase) DeleteData(id int) int {

	row := usecase.userData.DelData(id)
	if row == -1 {
		return -1
	}

	return row
}

func (usecase *userUsecase) PostData(data user.Core) int {

	if data.Password == "" || data.Email == "" || data.Fullname == "" || data.Role == "" || data.Status == "" || data.Team == "" {
		return -1
	}

	if data.Role == "Admin" {
		data.Role = "admin"
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	data.Password = string(hash)

	row := usecase.userData.InsertData(data)
	if row == -1 {
		return -1
	}

	return row
}

func (usecase *userUsecase) GetProfile(id int) (user.Core, error) {

	data, err := usecase.userData.SelectProfile(id)
	if err != nil {
		return user.Core{}, err
	}

	return data, nil
}
