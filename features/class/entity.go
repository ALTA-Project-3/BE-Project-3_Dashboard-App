package class

type CoreClass struct {
	ID        uint
	Name      string
	StartDate string
	EndDate   string
	UserID    uint
}

type DataInterface interface {
	InsertClass(core CoreClass) (string, error)
	SelectAllClass() ([]CoreClass, error)
	UpdateAClass(core CoreClass, idclass uint) (string, error)
	DeleteAClass(classid uint) (string, error)
	GetClassUserid(classid uint) uint
	// GetToken(email, password string) (string, error)
}

type UsecaseInterface interface {
	CreateClass(core CoreClass) (string, error)
	GetAllClass() ([]CoreClass, error)
	EditAClass(core CoreClass, idclass uint) (string, error)
	DeleteAClass(userid, classid uint) (string, error)
	// GetToken(email, password string) (string, error)
}
