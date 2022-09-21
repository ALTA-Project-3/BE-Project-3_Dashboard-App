package delivery

import "project/dashboard/features/mentee"

type ReqCore struct {
	Name          string       `json:"name" form:"name"`
	Address       string       `json:"address" form:"address"`
	Homeaddress   string       `json:"homeaddress" form:"homeaddress"`
	Email         string       `json:"email" form:"email"`
	Gender        string       `json:"gender" form:"gender"`
	Telegram      string       `json:"telegram" form:"telegram"`
	Phone         uint64       `json:"phone" form:"phone"`
	EmergencyData ReqEmergency `json:"emergencydata" form:"emergencydata"`
	Education     ReqEducation `json:"educationdata" form:"educationdata"`
	Class         string       `json:"class" form:"class"`
}

type ReqEmergency struct {
	Name   string `json:"name" form:"name"`
	Phone  uint64 `json:"phone" form:"phone"`
	Status string `json:"status" form:"status"`
}

type ReqEducation struct {
	Type     string `json:"type" form:"type"`
	Major    string `json:"major" form:"major"`
	Graduate string `json:"graduate" form:"graduate"`
}

func (req *ReqCore) toCore() mentee.MenteeCore {
	emergency := mentee.EmergencyCore{
		Name:   req.EmergencyData.Name,
		Phone:  req.EmergencyData.Phone,
		Status: req.EmergencyData.Status,
	}

	education := mentee.EducationCore{
		Type:     req.Education.Type,
		Major:    req.Education.Major,
		Graduate: req.Education.Graduate,
	}

	return mentee.MenteeCore{
		Name:          req.Name,
		Address:       req.Address,
		Homeaddress:   req.Homeaddress,
		Email:         req.Email,
		Gender:        req.Gender,
		Telegram:      req.Telegram,
		Phone:         req.Phone,
		EmergencyData: emergency,
		Education:     education,
		Class:         req.Class,
	}
}
