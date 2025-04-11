package video

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

func (r *repo) Find(ctx context.Context, spec spec.I) ([]model.Video, error) {
	var data []model.Video

	err := spec.GormQuery(r.db.WithContext(ctx)).Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, err
}

func (p *repo) Get(ctx context.Context, spec spec.I) (model.Video, error) {
	var data model.Video

	err := spec.GormQuery(p.db.WithContext(ctx)).First(&data).Error
	if err != nil {
		return data, err
	}

	return data, err
}

func (r *repo) Update(ctx context.Context, data model.Video) error {
	return r.db.WithContext(ctx).Save(data).Error
}
