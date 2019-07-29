package project

import (
	// 
)

// Service contains services to manage data in the repository.
type Service interface {
	UserService

	Close() error
}

type UserService interface {
	FetchIndexUser(*PageRequest) ([]*User, int, *SingleResponse)
	FetchShowUser(string) (*User, *SingleResponse)
	FetchStoreUser(*User) (*User, *SingleResponse)
	FetchUpdateUser(string, *User) (*User, *SingleResponse)
	FetchDestroyUser(string) *SingleResponse
}
