package usecase

import "project/dashboard/features/user"

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
