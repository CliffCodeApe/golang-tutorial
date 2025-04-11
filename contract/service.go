package contract

import "golang-tutorial/dto"

type Service struct {
	Intro IntroService
	User  UserService
}

// type exampleService interface {
// Code here
// }

type IntroService interface {
	GetIntro(introID int) (*dto.IntroResponse, error)
	CreateIntro(payload *dto.IntroRequest) (*dto.IntroResponse, error)
	UpdateIntro(id int, payload *dto.IntroRequest) (*dto.IntroResponse, error)
	DeleteIntro(id int) (*dto.IntroResponse, error)
}

type UserService interface {
	GetUser(userID int) (*dto.UserResponse, error)
	Register(payload *dto.UserRequest) (*dto.UserResponse, error)
	Login(payload *dto.UserRequest) (*dto.UserResponse, error)
}
