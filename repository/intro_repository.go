package repository

import (
	"golang-tutorial/contract"
	"golang-tutorial/entity"

	"gorm.io/gorm"
)

type IntroRepo struct {
	db *gorm.DB
}

func implIntroRepository(db *gorm.DB) contract.IntroRepository {
	return &IntroRepo{
		db: db,
	}
}

func (r *IntroRepo) GetIntro(id int) (*entity.Intro, error) {
	var intro entity.Intro
	err := r.db.Table("intro").Where("id = ?", id).First(&intro).Error
	if err != nil {
		return nil, err
	}
	return &intro, err
}

func (r *IntroRepo) CreateIntro(intro *entity.Intro) error {
	return r.db.Table("intro").Create(intro).Error
}

func (r *IntroRepo) UpdateIntro(id int, intro *entity.Intro) error {
	return r.db.Where("id = ?", id).Updates(&intro).Error
}

func (r *IntroRepo) DeleteIntro(id int) error {
	return r.db.Where("id = ?", id).Delete(&entity.Intro{}).Error
}
