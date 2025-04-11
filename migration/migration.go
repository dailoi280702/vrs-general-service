package migration

import (
	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	migrations := migrate.FileMigrationSource{
		Dir: "migration/migrations",
	}

	_, err = migrate.Exec(sqlDB, "mysql", migrations, migrate.Up)

	return err
}
