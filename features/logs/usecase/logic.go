package usecase

import "project/dashboard/features/logs"

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
