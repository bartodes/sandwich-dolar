package store

import (
	"database/sql"
	"fmt"
	"log"
	"sandwich-dolar/menu"

	_ "github.com/lib/pq"
)

/*

- ID (autoinctemental)
- Nombre
- Cantidad sandwich
- Total venta (en dolares)

*/

// Mejora: Implementar interfaces y structs

func OpenDB(secret, dbName string) *sql.DB {
	conStr := fmt.Sprintf("postgres://postgres:%s@localhost:5432/%s?sslmode=disable", secret, dbName)
	db, err := sql.Open("postgres", conStr)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func TestDb(db *sql.DB) {

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func InstertVenta(db *sql.DB, orden menu.Orden) int {
	query := `INSERT INTO ventas (nombre, cantidad_sandwich, total) VALUES ($1, $2, $3) RETURNING id`

	var id int

	if err := db.QueryRow(query, orden.Nombre, orden.Cantidad, orden.Total).Scan(&id); err != nil {
		log.Fatal(err)
	}

	return id
}
