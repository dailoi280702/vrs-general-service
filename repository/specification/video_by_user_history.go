package specification

import (
	"github.com/dailoi280702/vrs-general-service/type/model"
	"gorm.io/gorm"
)

type videoByUserHisotry struct {
	userId int64
}

func VideoByUserHisotry(userId int64) I {
	return &videoByUserHisotry{
		userId: userId,
	}
}

func (v *videoByUserHisotry) GormQuery(db *gorm.DB) *gorm.DB {
	return db.Model(&model.Video{}).
		Select("videos.*").
		Joins("JOIN watch_history on videos.id = watch_history.video_id").
		Where("watch_history.user_id = ?", v.userId).
		Order("watch_history.created_at DESC")
}
