package userstr

import (
	"github.com/lapitskyss/go_backend_2/model"
	"github.com/lapitskyss/go_backend_2/sharding"
)

type UserStore struct {
	s *sharding.Sharding
}

func NewUserStore(s *sharding.Sharding) *UserStore {
	return &UserStore{
		s: s,
	}
}

func (us *UserStore) Create(u *model.User) error {
	c, err := us.s.Master(u.UserId)
	if err != nil {
		return err
	}

	_, err = c.Exec(`INSERT INTO "users" VALUES ($1, $2, $3, $4)`, u.UserId, u.Name, u.Age, u.Spouse)
	return err
}

func (us *UserStore) Read(id int) (*model.User, error) {
	c, err := us.s.Slave(id)
	if err != nil {
		return nil, err
	}

	u := model.User{
		UserId: id,
	}

	err = c.QueryRow(`SELECT "name", "age", "spouse" FROM "users" WHERE "user_id" = $1`, id).
		Scan(&u.Name, &u.Age, &u.Spouse)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
