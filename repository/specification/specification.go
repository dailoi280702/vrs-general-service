package specification

import "gorm.io/gorm"

type I interface {
	GormQuery(db *gorm.DB) *gorm.DB
}

type CompositeQuery []I

func (c CompositeQuery) GormQuery(db *gorm.DB) *gorm.DB {
	for i := range c {
		db = c[i].GormQuery(db)
	}

	return db
}
