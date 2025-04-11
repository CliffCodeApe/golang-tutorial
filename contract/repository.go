package contract

import "golang-tutorial/entity"

type Repository struct {
	Intro IntroRepository
	User  UserRepository
}

// type exampleRepository interface {
// Code here
// }

type IntroRepository interface {
	CreateIntro(intro *entity.Intro) error
	GetIntro(id int) (*entity.Intro, error)
	UpdateIntro(id int, intro *entity.Intro) error
	DeleteIntro(id int) error
}

type UserRepository interface {
	GetUser(id int) (*entity.User, error)
	CreateUser(user *entity.User) error
	CheckEmail(email string) (bool, error)
	GetUserByEmail(email string) (*entity.User, error)
}
