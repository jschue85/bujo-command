package bujo

import "github.com/jschue85/bujo-command/data"

type UserService struct {
	Store data.BujoUserStore
}

func (s *UserService) GetUser(userName string) (data.User, error) {
	user, err := s.Store.GetUserByUserName(userName)
	if err != nil {
		return data.User{}, err
	}

	return user, nil
}

func (s *UserService) AddUser(userName string, firstName string, lastName string) error {
	user := data.User{
		UserName:  userName,
		FirstName: firstName,
		LastName:  lastName,
	}
	err := s.Store.AddUser(user)
	if err != nil {
		return err
	}

	return nil
}
