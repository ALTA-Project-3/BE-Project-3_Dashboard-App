package usecase

import (
	"project/dashboard/features/logs"
)

type Service struct {
	do logs.DataInterface
}

func New(data logs.DataInterface) logs.UsecaseInterface {
	return &Service{
		do: data,
	}
}

func (service *Service) AddLog(core logs.CoreLogs) (string, error) {
	msg, err := service.do.InsertLog(core)
	return msg, err
}

func (service *Service) DetailMentee(menteeid uint) ([]logs.CoreLogs, []logs.User, logs.Mentee, logs.Education, string, error) {
	listlogs, msg1, err1 := service.do.SelectAMenteeLogs(menteeid)
	if err1 != nil {
		return nil, nil, logs.Mentee{}, logs.Education{}, msg1, err1
	}

	// fmt.Println(listlogs)

	listuser, msg2, err2 := service.do.SelectUserList(listlogs)
	if err2 != nil {
		return nil, nil, logs.Mentee{}, logs.Education{}, msg2, err2
	}

	// fmt.Println(listuser)

	datamentee, msg3, err3 := service.do.SelectAMenteePrivateData(menteeid)
	if err3 != nil {
		return nil, nil, logs.Mentee{}, logs.Education{}, msg3, err3
	}

	// fmt.Println(datamentee)

	menteeeducation, msg4, err4 := service.do.SelectAMenteeEducation(menteeid)
	if err4 != nil {
		return nil, nil, logs.Mentee{}, logs.Education{}, msg4, err4
	}

	// fmt.Println(menteeeducation)

	return listlogs, listuser, datamentee, menteeeducation, "Sukses Mengakses Semua Data", nil
}

// SelectAMenteeLogs(menteeid uint) ([]CoreLogs, error)
// SelectUserList(data []CoreLogs) ([]User, error)
// SelectAMenteePrivateData(menteeid uint) (Mentee, error)
// SelectAMenteeEducation(menteeid uint) (Education, error)
// SelectAMenteeClass(menteeid uint) (Class, error)
