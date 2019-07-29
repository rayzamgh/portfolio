package service

import (
	"gitlab.com/standard-go/project/internal/app/project"
)

/*
|--------------------------------------------------------------------------
| Services
|--------------------------------------------------------------------------
|
*/

func (s *Service) FetchIndexUser(pageRequest *project.PageRequest) ([]*project.User, int, *project.SingleResponse) {
	return s.Repo.FetchIndexUser(pageRequest)
}

func (s *Service) FetchShowUser(id string) (*project.User, *project.SingleResponse) {
	return s.Repo.FetchShowUser(id)
}

func (s *Service) FetchStoreUser(data *project.User) (*project.User, *project.SingleResponse) {
	if err := s.validateUser(data); err != nil {
		return nil, err
	}

	return s.Repo.FetchStoreUser(data)
}

func (s *Service) FetchUpdateUser(id string, data *project.User) (*project.User, *project.SingleResponse) {
	if err := s.validateUser(data); err != nil {
		return nil, err
	}
	
	return s.Repo.FetchUpdateUser(id, data)
}

func (s *Service) FetchDestroyUser(id string) *project.SingleResponse {
	return s.Repo.FetchDestroyUser(id)
}

/*
|--------------------------------------------------------------------------
| Validator
|--------------------------------------------------------------------------
|
*/

func (s *Service) validateUser(data *project.User) *project.SingleResponse {
	messages := make(project.Message)

	if data.FullName == "" {
		messages["full_name"] = []string{"This field is required."}
	}

	if len(messages) > 0 {
		response := ResponseValidation(messages)

		return response
	}

	return nil
}
