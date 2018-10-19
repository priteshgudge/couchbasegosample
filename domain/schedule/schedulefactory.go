package schedule

import (
	"time"

	"github.com/priteshgudge/couchbasegosample/domain"
)

func getUtcTime() int64 {
	return int64(time.Now().Nanosecond() / 1000000.0)
}

type ScheduleFactory struct {
}

func (p *ScheduleFactory) NewSampleSchedule() *Schedule {
	schedule := Schedule{
		UUID:          domain.NewUID(),
		Name:          "Apple Schedule",
		CreatedByName: "Clark",
		UpdatedByName: "Clark",
		CreatedOn:     getUtcTime(),
		UpdatedOn:     getUtcTime(),
		CropName:      "Apple",
		SegmentName:   "Fuji",
	}
	product1 := Product{
		ProductName: "Fertilizer",
		SKUCode:     "FERTI-1",
	}
	product2 := Product{
		ProductName: "Pesticide",
		SKUCode:     "PESTI-2",
	}
	schedule.BasicDetails = BASICDETAILS{}
	schedule.SowingDetails = SOWINGDETAILS{}
	schedule.Products = []Product{product1, product2}
	schedule.Status = "CREATED"

	return &schedule
}

func (p *ScheduleFactory) NewSampleSchedule2() *Schedule {
	schedule := Schedule{
		UUID:          domain.NewUID(),
		Name:          "Apple Schedule",
		CreatedByName: "Jason",
		UpdatedByName: "Jason",
		CreatedOn:     getUtcTime(),
		UpdatedOn:     getUtcTime(),
		CropName:      "Banana",
		SegmentName:   "Small",
	}
	product1 := Product{
		ProductName: "Fertilizer",
		SKUCode:     "FERTI-19",
	}
	product2 := Product{
		ProductName: "Pesticide",
		SKUCode:     "PESTI-29",
	}
	schedule.BasicDetails = BASICDETAILS{}
	schedule.SowingDetails = SOWINGDETAILS{}
	schedule.Products = []Product{product1, product2}
	schedule.Status = "RELEASED"

	return &schedule
}
