package delivery

import "project/dashboard/features/class"

type Response struct {
	ID        uint
	Name      string
	StartDate string
	EndDate   string
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
