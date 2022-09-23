package class

import "time"

type CoreClass struct {
	ID        uint
	Name      string
	StartDate time.Time
	EndDate   time.Time
	UserID    uint
}

type DataInterface interface {
	InsertClass(core CoreClass) (string, error)
	SelectAllClass(page int) ([]CoreClass, error)
	UpdateAClass(core CoreClass, idclass uint) (string, error)
	DeleteAClass(classid uint) (string, error)
	GetClassUserid(classid uint) uint
	// GetToken(email, password string) (string, error)
}

type UsecaseInterface interface {
	CreateClass(core CoreClass) (string, error)
	GetAllClass(page int) ([]CoreClass, error)
	EditAClass(core CoreClass, idclass uint) (string, error)
	DeleteAClass(userid, classid uint) (string, error)
	// GetToken(email, password string) (string, error)
}
