package schedule

import (
	"github.com/priteshgudge/couchbasegosample/domain"
)

type Category struct {
	LevelOne string `json:"levelOne,omitempty"`
	LevelTwo string `json:"levelTwo,omitempty"`
}

type Product struct {
	SKUCode     string
	ProductName string
}

type SOWINGDETAILS struct {
	SowingMethod struct {
		UUID      string        `json:"uuid"`
		Method    string        `json:"method"`
		Localised []interface{} `json:"localised"`
	} `json:"sowingMethod"`
	Season struct {
		UID       string        `json:"uuid"`
		Name      string        `json:"name"`
		Localised []interface{} `json:"localised"`
	} `json:"season"`
	SowingRange struct {
		StartDate int64 `json:"startDate"`
		EndDate   int64 `json:"endDate"`
	} `json:"sowingRange"`
}

type BASICDETAILS struct {
	IrrigationType []struct {
		UID       string        `json:"uuid"`
		Type      string        `json:"type"`
		Localised []interface{} `json:"localised"`
	} `json:"irrigationType"`
	KitType []struct {
		UID       string        `json:"uuid"`
		Type      string        `json:"type"`
		Localised []interface{} `json:"localised"`
		Products  []interface{} `json:"products"`
	} `json:"kitType"`
	SoilType []struct {
		UID       string        `json:"uuid"`
		Type      string        `json:"type"`
		Localised []interface{} `json:"localised"`
	} `json:"soilType"`
	Region struct {
		State    string `json:"state"`
		District string `json:"district"`
		Taluka   string `json:"taluka"`
	} `json:"region"`
	ScheduleDuration int `json:"scheduleDuration"`
}

type Schedule struct {
	UUID          domain.UID    `json:"uuid"`
	Name          string        `json:"name"`
	CreatedByName string        `json:"createdByName"`
	UpdatedByName string        `json:"updatedByName"`
	CreatedByID   int           `json:"createdById"`
	UpdatedByID   int           `json:"updatedById"`
	CreatedOn     int64         `json:"createdOn"`
	UpdatedOn     int64         `json:"updatedOn"`
	Status        string        `json:"status"`
	CropName      string        `json:"cropName"`
	SegmentName   string        `json:"segmentName"`
	Products      []Product     `json:"products"`
	DisplayImage  string        `json:"displayImage"`
	BasicDetails  BASICDETAILS  `json:"BASIC_DETAILS"`
	SowingDetails SOWINGDETAILS `json:"SOWING_DETAILS"`
}

type ScheduleResult struct {
	ScheduleDoc Schedule `json:"schedule-bucket"`
}
