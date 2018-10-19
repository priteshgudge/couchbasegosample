package main

import (
	"fmt"

	"github.com/priteshgudge/couchbasegosample/domain/schedule/dbadapter"
	gocb "gopkg.in/couchbase/gocb.v1"
)

func main() {

	cbCluster, _ := gocb.Connect("couchbase://0.0.0.0")
	err := cbCluster.Authenticate(gocb.PasswordAuthenticator{"Administrator", "password"})
	if err != nil {
		fmt.Println(err)
	}
	bucket, err := cbCluster.OpenBucket(dbadapter.BucketName, "")
	if err != nil {
		fmt.Println(err)
	}
	bucket.Manager("", "").CreatePrimaryIndex("", true, false)

	cbRepo := dbadapter.CouchbaseRepository{Cl: cbCluster}
	// scheduleFactory := schedule.ScheduleFactory{}
	// schedule := scheduleFactory.NewSampleSchedule()
	// schedule2 := scheduleFactory.NewSampleSchedule2()

	// id1, err := cbRepo.Store(schedule)
	// id2, err := cbRepo.Store(schedule2)
	// fmt.Printf("%v %v", id1, id2)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// out, err := cbRepo.Get(*id1)
	// out2, err := cbRepo.Get(*id2)
	// fmt.Printf("%v %v", out, out2)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	out3, err := cbRepo.GetByCropName("Apple")
	out4, err := cbRepo.GetByCropName("Banana")
	fmt.Printf("%v %v", out3, out4)
	if err != nil {
		fmt.Println(err)
	}
}
