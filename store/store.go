package store

type DB interface {
	DbConnect()
}

type PostgresDB struct{}
