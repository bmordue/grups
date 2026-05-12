package db

import (
    "database/sql"
    _ "github.com/jackc/pgx/v5/stdlib"
    "time"
)

func Connect(dsn string) (*sql.DB, error) {
    db, err := sql.Open("pgx", dsn)
    if err != nil {
        return nil, err
    }
    db.SetMaxOpenConns(10)
    db.SetConnMaxLifetime(30 * time.Minute)
    db.SetConnMaxIdleTime(5 * time.Minute)
    if err := db.Ping(); err != nil {
        return nil, err
    }
    return db, nil
}
