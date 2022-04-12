package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	// connect
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal("db connection failure: %v", err)
	}

	// test connection
	err = db.Ping()
	if err != nil {
		log.Fatal("db ping failure: %v", err)
	}

	return &Adapter{
		db: db,
	}, nil
}

func (dbAdapter *Adapter) CloseDbConnection() {
	err := dbAdapter.db.Close()
	if err != nil {
		log.Fatal("db close failure: %v", err)
	}
}

func (dbAdapter *Adapter) AddToHistory(answer int32, operation string) error {
	// create insert query
	queryString, args, err := sq.Insert("history").
		Columns("answer", "operation", "created_at").
		Values(answer, operation, time.Now()).ToSql()
	if err != nil {
		return err
	}

	// execute query
	_, err = dbAdapter.db.Exec(queryString, args...)
	if err != nil {
		return err
	}

	return nil
}
