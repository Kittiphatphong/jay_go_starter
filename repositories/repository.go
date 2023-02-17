package repositories

import "gorm.io/gorm"

type Repository interface {
	//Insert your function interface
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	// db.Migrator().DropTable(Table{})
	// db.AutoMigrate(Table{})
	return &repository{db: db}
}
