package delivery

import (
	"project/dashboard/features/class"
	"time"
)

type Response struct {
	ID        uint
	Name      string
	StartDate time.Time
	EndDate   time.Time
	UserID    uint
}

func CoreToResponse(data class.CoreClass) Response {
	return Response{
		ID:        data.ID,
		Name:      data.Name,
		StartDate: data.StartDate,
		EndDate:   data.EndDate,
		UserID:    data.UserID,
	}
}

func ListCoreResponse(data []class.CoreClass) []Response {
	var response []Response
	for _, v := range data {
		response = append(response, CoreToResponse(v))
	}

	return response
}
