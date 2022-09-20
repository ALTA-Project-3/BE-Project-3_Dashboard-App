package delivery

type Request struct {
	ID       uint   `json:"id" form:"id"`
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Team     string `json:"team" form:"team"`
	Role     string `json:"role" form:"role"`
	Status   string `json:"status" form:"status"`
}
