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

type DataInterface interface {
	SelectAll(page, token int) ([]Core, error)
	UpdateData(data Core) int
}

type UsecaseInterface interface {
	GetAll(page, token int) ([]Core, error)
	PutData(data Core) int
}
