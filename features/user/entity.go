package user

type Core struct {
	ID       uint
	Fullname string
	Email    string
	Password string
	Team     string
	Role     string
	Status   string
}

type DashBoard struct {
	Active          int64
	Placement       int64
	FeedBack        int64
	ActiveInMonth   []MonthActive
	GraduateInMonth []MonthGraduate
}

type MonthActive struct {
	Month int64
	Count int64
}

type MonthGraduate struct {
	Month int64
	Count int64
}

type DataInterface interface {
	SelectProfile(id int) (Core, DashBoard, error)
	SelectAll(page, token int) ([]Core, error)
	UpdateData(data Core) int
	DelData(id int) int
	InsertData(data Core) int
}

type UsecaseInterface interface {
	GetProfile(id int) (Core, DashBoard, error)
	GetAll(page, token int) ([]Core, error)
	PutData(data Core) int
	DeleteData(id int) int
	PostData(data Core) int
}
