package delivery

import "project/dashboard/features/logs"

type Request struct {
	Status   string `json:"status" form:"status"`
	Feedback string `json:"feedback" form:"feedback"`
	Url      string `json:"url" form:"url"`
}

func (req *Request) ReqToCore(userid, menteeid uint) logs.CoreLogs {
	return logs.CoreLogs{
		UserID:   userid,
		MenteeID: menteeid,
		Status:   req.Status,
		Feedback: req.Feedback,
		Url:      req.Url,
	}
}
