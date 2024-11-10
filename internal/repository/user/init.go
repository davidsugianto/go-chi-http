package user

import "gorm.io/gorm"

type Repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) (*Repo, error) {
	return &Repo{
		db: db,
	}, nil
}
