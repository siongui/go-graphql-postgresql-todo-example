package tododb

import (
	"testing"
	"time"

	"github.com/siongui/go-graphql-postgresql-todo-example/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	// Load config
	if err := config.LoadConfig(); err != nil {
		t.Fatal(err)
	}

	gormdsn := "host=" + config.Config.Database.Postgres.Host +
		" user=" + config.Config.Database.Postgres.User +
		" password=" + config.Config.Database.Postgres.Password +
		" dbname=" + config.Config.Database.Postgres.Dbname +
		" port=" + config.Config.Database.Postgres.Port +
		" sslmode=disable TimeZone=" + config.Config.App.Timezone
	db, err := gorm.Open(postgres.Open(gormdsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	ted, _ := time.Parse(time.RFC3339, "2022-01-23T11:45:26.371Z")
	td := Todo{
		ContentCode: "TD001",
		StartDate:   ted,
		EndDate:     time.Now(),
		CreatedBy:   "me"}
	result := db.Create(&td)

	if result.Error != nil {
		t.Fatal(result.Error)
	}

	t.Log(td.ID)
	t.Log(result.RowsAffected)
}
