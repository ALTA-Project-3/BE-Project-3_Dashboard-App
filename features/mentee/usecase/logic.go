package usecase

import "project/dashboard/features/mentee"

type menteeUsecase struct {
	menteeData mentee.DataInterface
}

func New(data mentee.DataInterface) mentee.UsecaseInterface {
	return &menteeUsecase{
		menteeData: data,
	}
}

func (usecase *menteeUsecase) PostData(data mentee.MenteeCore, idToken int) int {

	if data.Address == "" || data.Class == "" || data.Homeaddress == "" || data.Name == "" || data.Email == "" || data.Gender == "" || data.Phone == "" || data.EmergencyData.Name == "" || data.EmergencyData.Phone == "" || data.EmergencyData.Status == "" || data.Education.Graduate == "" || data.Education.Major == "" || data.Education.Type == "" {
		return -2
	} else if data.Gender != "Male" && data.Gender != "Female" {
		return -3
	}

	row := usecase.menteeData.InsertData(data, idToken)
	if row == -1 {
		return -1
	}

	return row

}

func (usecase *menteeUsecase) GetData(page, idToken int) ([]mentee.Join, error) {

	data, err := usecase.menteeData.SelectData(page, idToken)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (usecase *menteeUsecase) UpdateData(data mentee.MenteeCore, menteeid int) (string, error) {
	msg, err := usecase.menteeData.UpdateData(data, menteeid)
	if err != nil {
		return "", err
	}
	return msg, nil
}
