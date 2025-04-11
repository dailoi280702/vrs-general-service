package specification

import (
	"github.com/dailoi280702/vrs-general-service/type/model"
	"gorm.io/gorm"
)

type videoByIds struct {
	ids []int64
}

func VideoByIds(id ...int64) I {
	return &videoByIds{
		ids: id,
	}
}

func (v *videoByIds) GormQuery(db *gorm.DB) *gorm.DB {
	if len(v.ids) == 0 {
		return db
	}

	if len(v.ids) == 1 {
		return db.Model(&model.Video{}).Where("id = ?", v.ids[0])
	}

	return db.Model(&model.Video{}).Where("id IN ?", v.ids)
}
