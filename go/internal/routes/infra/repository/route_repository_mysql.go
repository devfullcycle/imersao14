package repository

import (
	"database/sql"
	"fmt"

	"github.com/devfullcycle/imersao14/go/internal/routes/entity"
)

type RouteRepositoryMysql struct {
	db *sql.DB
}

func NewRouteRepositoryMysql(db *sql.DB) *RouteRepositoryMysql {
	return &RouteRepositoryMysql{
		db: db,
	}
}

func (r *RouteRepositoryMysql) Create(route *entity.Route) error {
	// insert new route on database
	sql := "INSERT INTO routes (id, name, distance, status, freight_price) VALUES (?, ?, ?, ?,?)"
	_, err := r.db.Exec(sql, route.ID, route.Name, route.Distance, route.Status, route.FreightPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *RouteRepositoryMysql) FindById(id string) (*entity.Route, error) {
	// find route on database
	sqlStmt := "SELECT id, name, distance, status, freight_price, started_at, finished_at FROM routes WHERE id = ?"
	row := r.db.QueryRow(sqlStmt, id)

	var route entity.Route
	var startedAt, finishedAt sql.NullTime // Temporary variables to handle nullable columns

	err := row.Scan(&route.ID, &route.Name, &route.Distance, &route.Status, &route.FreightPrice, &startedAt, &finishedAt)
	if err != nil {
		return nil, err
	}

	// Check if startedAt and finishedAt have valid values and assign them to the entity.Route struct
	if startedAt.Valid {
		route.StartedAt = startedAt.Time
	}
	if finishedAt.Valid {
		route.FinishedAt = finishedAt.Time
	}

	return &route, nil
}

func (r *RouteRepositoryMysql) Update(route *entity.Route) error {
	// update route on database
	// format started at to database
	startedAt := route.StartedAt.Format("2006-01-02 15:04:05")
	fmt.Println(startedAt)
	finishedAt := route.FinishedAt.Format("2006-01-02 15:04:05")
	sql := "UPDATE routes SET status = ?, freight_price = ?, started_at=?, finished_at = ? WHERE id = ?"
	_, err := r.db.Exec(sql, route.Status, route.FreightPrice, startedAt, finishedAt, route.ID)
	if err != nil {
		return err
	}
	return nil
}
