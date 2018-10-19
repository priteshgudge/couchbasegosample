package dbadapter

import (
	"github.com/priteshgudge/couchbasegosample/domain/schedule"
)

type Reader interface {
	Get(uuid string) (*schedule.Schedule, error)
	GetByName(name string) ([]*schedule.Schedule, error)
	GetByCropName(name string) ([]*schedule.Schedule, error)
}

type Writer interface {
	Store(scheduleObj *schedule.Schedule) error
	Delete(scheduleId *string)
}

type Filter interface {
	FindByProductName(productName string) ([]*schedule.Schedule, error)
	FindByCategory(categoryName string) ([]*schedule.Schedule, error)
}

type Repository interface {
	Reader
	Writer
	Filter
}
