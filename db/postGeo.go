package db

import (
	"context"
	"fmt"
	"log"
	"user-service/models"

	"github.com/jackc/pgx/v4"
	geo "github.com/paulmach/go.geo"
)

var conn *pgx.Conn

// InitializeDB initializes the PostgreSQL connection using pgx
func InitializeDB(dbConfig models.DatabaseConfig) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
}

// CreatePlace inserts a new place with geospatial data (using go.geo)
func CreatePlace(name string, lat, lon float64) error {
	// Create a geo.Point using go.geo
	point := geo.NewPoint(lon, lat)

	// Insert the point as a WKT string into PostGIS
	wkt := point.ToWKT()

	query := `INSERT INTO places (name, location) VALUES ($1, ST_SetSRID(ST_GeomFromText($2), 4326))`
	_, err := conn.Exec(context.Background(), query, name, wkt)
	if err != nil {
		return err
	}
	return nil
}

// FindPlaceByLocation queries places within a given radius from a point
func FindPlaceByLocation(lat, lon float64) ([]string, error) {
	var names []string
	query := `SELECT name FROM places WHERE ST_DWithin(location, ST_SetSRID(ST_MakePoint($1, $2), 4326), 5000)`
	rows, err := conn.Query(context.Background(), query, lon, lat)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Loop through the rows and collect place names
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		names = append(names, name)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return names, nil
}
