package user

import (
	"context"

	spec "github.com/dailoi280702/vrs-general-service/repository/specification"
	"github.com/dailoi280702/vrs-general-service/type/model"
	"gorm.io/gorm"
)

func NewRepo(getDB func() *gorm.DB) I {
	return &repo{getDB()}
}

type repo struct {
	db *gorm.DB
}

func (r *repo) Count(ctx context.Context, spec spec.I) (int64, error) {
	var data int64

	err := spec.GormQuery(r.db.WithContext(ctx)).Count(&data).Error
	if err != nil {
		return data, err
	}

	return data, err
}

func (r *repo) Find(ctx context.Context, spec spec.I) ([]model.User, error) {
	var data []model.User

	err := spec.GormQuery(r.db.WithContext(ctx)).Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, err
}

func (p *repo) Get(ctx context.Context, spec spec.I) (model.User, error) {
	var data model.User

	err := spec.GormQuery(p.db.WithContext(ctx)).First(&data).Error
	if err != nil {
		return data, err
	}

	return data, err
}

func (p *repo) Create(ctx context.Context, data model.User) error {
	return p.db.WithContext(ctx).Create(&data).Error
}
