package tododb

// https://www.google.com/search?q=golang+migrate+gorm
// https://stackoverflow.com/questions/64510093/gorm-migration-using-golang-migrate-migrate

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	// gorm.Model includes ID, CreatedAt, UpdatedAt, DeletedAt
	// https://gorm.io/docs/models.html#gorm-Model
	gorm.Model
	ContentCode string `gorm:"not null;unique"`
	ContentName string
	Description string
	StartDate   time.Time
	EndDate     time.Time
	Status      string
	CreatedBy   string
	UpdatedBy   string
}