package main

import (
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"

	"github.com/lapitskyss/go_backend_2/model"
	"github.com/lapitskyss/go_backend_2/sharding"
	"github.com/lapitskyss/go_backend_2/store/activitystr"
	"github.com/lapitskyss/go_backend_2/store/userstr"
)

func main() {
	p := sharding.NewPool()
	m := sharding.NewManager(10)
	s := sharding.NewSharding(m, p)

	m.Add(&sharding.Shard{
		Address: "port=8100 user=test password=test dbname=test sslmode=disable",
		Slaves:  []string{"port=8101 user=test password=test dbname=test sslmode=disable"},
		Number:  0,
	})
	m.Add(&sharding.Shard{
		Address: "port=8110 user=test password=test dbname=test sslmode=disable",
		Slaves:  []string{"port=8111 user=test password=test dbname=test sslmode=disable"},
		Number:  1,
	})
	m.Add(&sharding.Shard{
		Address: "port=8120 user=test password=test dbname=test sslmode=disable",
		Number:  2,
	})

	userStore := userstr.NewUserStore(s)
	activityStore := activitystr.NewActivityStore(s)

	//createUsers(userStore)
	getUsers(userStore)
	//createActivities(activityStore)
	getActivities(activityStore)
}

func createUsers(us *userstr.UserStore) {
	uu := []*model.User{
		{1, "Joe Biden", 78, 10},
		{10, "Jill Biden", 69, 1},
		{13, "Donald Trump", 74, 25},
		{25, "Melania Trump", 78, 13},
	}
	for _, u := range uu {
		err := us.Create(u)
		if err != nil {
			fmt.Println(fmt.Errorf("error on create user %v: %w", u, err))
		}
	}
}

func getUsers(us *userstr.UserStore) {
	u, err := us.Read(10)
	if err != nil {
		log.Fatalf("Err to get user, %v", err)
	}

	log.Println(u)
}

func createActivities(as *activitystr.ActivityStore) {
	aa := []*model.Activity{
		{1, "Download", time.Now()},
		{10, "Download", time.Now()},
	}
	for _, a := range aa {
		err := as.Create(a)
		if err != nil {
			fmt.Println(fmt.Errorf("error on create user %v: %w", a, err))
		}
	}
}

func getActivities(as *activitystr.ActivityStore) {
	a, err := as.Read(10)
	if err != nil {
		log.Fatalf("Err to get activity, %v", err)
	}

	log.Println(a)
}
