package activitystr

import (
	"github.com/lapitskyss/go_backend_2/model"
	"github.com/lapitskyss/go_backend_2/sharding"
)

type ActivityStore struct {
	s *sharding.Sharding
}

func NewActivityStore(s *sharding.Sharding) *ActivityStore {
	return &ActivityStore{
		s: s,
	}
}

func (as *ActivityStore) Create(a *model.Activity) error {
	c, err := as.s.Master(a.UserId)
	if err != nil {
		return err
	}

	_, err = c.Exec(`INSERT INTO activities (user_id, name, date) VALUES ($1, $2, $3)`, a.UserId, a.Name, a.Date)
	return err
}

func (as *ActivityStore) Read(id int) (*model.Activity, error) {
	c, err := as.s.Random(id)
	if err != nil {
		return nil, err
	}

	a := model.Activity{
		UserId: id,
	}

	err = c.QueryRow(`SELECT name, date FROM activities WHERE user_id = $1`, id).
		Scan(&a.Name, &a.Date)
	if err != nil {
		return nil, err
	}

	return &a, nil
}
