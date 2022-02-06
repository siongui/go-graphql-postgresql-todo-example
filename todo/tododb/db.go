package tododb

// https://www.google.com/search?q=golang+migrate+gorm
// https://stackoverflow.com/questions/64510093/gorm-migration-using-golang-migrate-migrate

import (
	"time"

	"gorm.io/driver/postgres"
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

type TodoStore interface {
	GetTodo(string) (Todo, error)
	Create(Todo) (Todo, error)
	Save(Todo) error
	Pages(count, page int) ([]Todo, int64, error)
	Search(int, int, map[string]interface{}) ([]Todo, int64, error)
}

type todoStore struct {
	db *gorm.DB
}

func (s *todoStore) GetTodo(id string) (t Todo, err error) {
	result := s.db.First(&t, id)
	err = result.Error
	return
}

func (s *todoStore) Create(t Todo) (Todo, error) {
	result := s.db.Omit("UpdatedBy").Create(&t)
	return t, result.Error
}

func (s *todoStore) Save(t Todo) error {
	result := s.db.Save(&t)
	return result.Error
}

// Pages returns the Todo records in the database, given the record count per
// page and n-th page.
func (s *todoStore) Pages(count, page int) (todos []Todo, totalCount int64, err error) {
	// Get total count
	result := s.db.Model(&Todo{}).Count(&totalCount)
	if result.Error != nil {
		err = result.Error
		return
	}

	// Get records in the given count and page.
	result = s.db.Limit(count).Offset((page - 1) * count).Order("created_at").Find(&todos)
	err = result.Error
	return
}

func (s *todoStore) Search(count, page int, condition map[string]interface{}) (todos []Todo, totalCount int64, err error) {
	query := s.db.Model(&Todo{})
	queryTotalCount := s.db.Model(&Todo{})
	for k, v := range condition {
		query.Where(k, v)
		queryTotalCount.Where(k, v)
	}

	queryResult := query.Limit(count).Offset((page - 1) * count).Find(&todos)
	if queryResult.Error != nil {
		err = queryResult.Error
		return
	}

	queryTotalCountResult := queryTotalCount.Count(&totalCount)
	err = queryTotalCountResult.Error
	return
}

func NewTodoStore(gormdsn string) (TodoStore, error) {
	db, err := gorm.Open(postgres.Open(gormdsn), &gorm.Config{})
	if err != nil {
		return &todoStore{}, err
	}
	return &todoStore{db: db}, nil
}
