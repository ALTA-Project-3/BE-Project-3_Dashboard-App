package delivery

import (
	"project/dashboard/features/mentee"
)

type ResponList struct {
	ID       uint
	Name     string
	Class    string
	Status   string
	Category string
	Gender   string
}

func toRespon(data mentee.Join) ResponList {
	return ResponList{
		ID:       data.ID,
		Name:     data.Name,
		Class:    data.Class,
		Status:   data.Status,
		Category: data.Category,
		Gender:   data.Gender,
	}
}

func toResponList(data []mentee.Join, class, category, status string) []ResponList {
	var resQuery []ResponList
	if category != "" && class != "" && status != "" {
		for key, v := range data {
			if category == v.Category && class == v.Class && status == v.Status {
				resQuery = append(resQuery, toRespon(data[key]))
			}
		}
		return resQuery
	} else if category != "" && class != "" && status == "" {
		for key, v := range data {
			if category == v.Category && class == v.Class {
				resQuery = append(resQuery, toRespon(data[key]))
			}
		}
		return resQuery
	} else if category != "" && class == "" && status != "" {
		for key, v := range data {
			if category == v.Category && status == v.Status {
				resQuery = append(resQuery, toRespon(data[key]))
			}
		}
		return resQuery
	} else if category == "" && class != "" && status != "" {
		for key, v := range data {
			if class == v.Class && status == v.Status {
				resQuery = append(resQuery, toRespon(data[key]))
			}
		}
		return resQuery
	} else if category == "" && class == "" && status != "" {
		for key, v := range data {
			if status == v.Status {
				resQuery = append(resQuery, toRespon(data[key]))
			}
		}
		return resQuery
	} else if category == "" && class != "" && status == "" {
		for key, v := range data {
			if class == v.Class {
				resQuery = append(resQuery, toRespon(data[key]))
			}
		}
		return resQuery
	} else if category != "" && class == "" && status == "" {
		for key, v := range data {
			if category == v.Category {
				resQuery = append(resQuery, toRespon(data[key]))
			}
		}
		return resQuery
	} else {
		for key := range data {
			resQuery = append(resQuery, toRespon(data[key]))
		}
		return resQuery
	}
}
