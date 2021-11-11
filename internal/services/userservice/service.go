package userservice

type Service struct {
	UserStore UsersStore
}

func NewUserService(userStore UsersStore) *Service {
	return &Service{
		UserStore: userStore,
	}
}
