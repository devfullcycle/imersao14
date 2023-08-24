// create unit test for the repository using testify mocks
package repository

import (
	"database/sql"
	"testing"

	"github.com/devfullcycle/imersao14/go/internal/routes/entity"
	_ "github.com/mattn/go-sqlite3"

	"github.com/stretchr/testify/require"
)

func TestRouteRepositoryMysqlCreate(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	require.Nil(t, err)
	defer db.Close()
	// create table
	sql := `CREATE TABLE routes (id VARCHAR(36) PRIMARY KEY,name VARCHAR(255) NOT NULL,distance FLOAT NOT NULL,status VARCHAR(255) NOT NULL,freight_price FLOAT,started_at DATETIME,finished_at DATETIME);`
	_, err = db.Exec(sql)
	repository := NewRouteRepositoryMysql(db)
	route := entity.NewRoute("1", "Route 1", 100)
	err = repository.Create(route)
	require.Nil(t, err)
	routeFinded, err := repository.FindById("1")
	require.Nil(t, err)
	require.Equal(t, route.ID, routeFinded.ID)
	require.Equal(t, route.Name, routeFinded.Name)
	require.Equal(t, route.Distance, routeFinded.Distance)
	require.Equal(t, route.Status, routeFinded.Status)
	require.Equal(t, route.FreightPrice, routeFinded.FreightPrice)
	require.Equal(t, route.StartedAt, routeFinded.StartedAt)
	require.Equal(t, route.FinishedAt, routeFinded.FinishedAt)
}
