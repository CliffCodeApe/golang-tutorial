package contract

import "golang-tutorial/entity"

type Repository struct {
	Intro IntroRepository
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
