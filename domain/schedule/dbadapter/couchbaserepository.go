package dbadapter

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/priteshgudge/couchbasegosample/domain"
	"github.com/priteshgudge/couchbasegosample/domain/schedule"
	gocb "gopkg.in/couchbase/gocb.v1"
)

type CouchbaseRepository struct {
	Cl *gocb.Cluster
}

var BucketName string = "schedule-bucket"

func (r *CouchbaseRepository) Get(uid domain.UID) (*schedule.Schedule, error) {
	var schedule schedule.Schedule
	bucket, _ := r.Cl.OpenBucket(BucketName, "")
	_, err := bucket.Get(uid.String(), &schedule)
	if gocb.IsKeyNotFoundError(err) {
		return nil, domain.ErrorNotFound
	} else if err != nil {
		return nil, domain.ErrorProcessingDB
	}

	return &schedule, nil
}

func (r *CouchbaseRepository) GetByName(name string) ([]*schedule.Schedule, error) {

	bucket, _ := r.Cl.OpenBucket(BucketName, "")

	query := gocb.NewN1qlQuery("SELECT * from schedule-bucket where name=$1")

	rows, err := bucket.ExecuteN1qlQuery(query, []interface{}{name})
	if err != nil {
		return nil, domain.ErrorProcessingDB
	}
	var row schedule.Schedule
	scheduleObjects := []*schedule.Schedule{}
	for rows.Next(&row) {
		var rowObject schedule.Schedule
		copier.Copy(&rowObject, &row)
		scheduleObjects = append(scheduleObjects, &rowObject)
		fmt.Printf("Row %v\n", row)
	}

	return scheduleObjects, nil
}

func (r *CouchbaseRepository) GetByCropName(name string) ([]*schedule.Schedule, error) {
	bucket, err := r.Cl.OpenBucket(BucketName, "")
	if err != nil {
		return nil, domain.ErrorProcessingDB
	}

	query := gocb.NewN1qlQuery("SELECT * from `schedule-bucket` where cropName=$1")

	rows, err := bucket.ExecuteN1qlQuery(query, []interface{}{name})
	if err != nil {
		return nil, domain.ErrorProcessingDB
	}

	var row interface{}
	scheduleRes := schedule.ScheduleResult{}
	scheduleObjects := []*schedule.Schedule{}
	for rows.Next(&row) {
		fmt.Printf("Row %v\n", row)
		rowJson, err := json.Marshal(row)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(rowJson[:]))
		err = json.Unmarshal(rowJson, &scheduleRes)
		if err != nil {
			fmt.Println(err)
		}
		scheduleObjects = append(scheduleObjects, &scheduleRes.ScheduleDoc)
		scheduleRes = schedule.ScheduleResult{}

	}

	return scheduleObjects, nil
}

// type Writer interface {

// 	Delete(scheduleId *string)
// }

func (r *CouchbaseRepository) Store(scheduleObj *schedule.Schedule) (*domain.UID, error) {
	if domain.IsNil(scheduleObj.UUID) {
		scheduleObj.UUID = domain.NewUID()
	}
	bucket, err := r.Cl.OpenBucket(BucketName, "")
	if err != nil {
		return nil, domain.ErrorProcessingDB
	}
	keyStr := domain.ToString(scheduleObj.UUID)
	_, err = bucket.Upsert(keyStr, scheduleObj, 0)
	if err != nil {
		fmt.Println(err)
		return nil, domain.ErrorProcessingDB
	}
	return &scheduleObj.UUID, nil

}
func (r *CouchbaseRepository) Delete(uid domain.UID) error {
	bucket, err := r.Cl.OpenBucket(BucketName, "")
	if err != nil {
		return domain.ErrorProcessingDB
	}
	var sched *schedule.Schedule
	cas, err := bucket.Get(uid.String(), &sched)
	_, err = bucket.Remove(uid.String(), cas)
	if err != nil {
		return domain.ErrorProcessingDB
	}

	return nil

}
