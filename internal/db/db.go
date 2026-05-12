package db

import (
    "database/sql"
    _ "modernc.org/sqlite"
    "fmt"
    "net/url"
    "os"
    "path/filepath"
    "strings"
    "time"
)

func ensureDir(path string) error {
    dir := "./data"
    if _, err := os.Stat(dir); os.IsNotExist(err) {
        return os.MkdirAll(dir, 0o755)
    }
    return nil
}

func ensureFileWritable(path string) error {
    f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0o644)
    if err != nil {
        return err
    }
    return f.Close()
}

func Connect(dsn string) (*sql.DB, error) {
    // For sqlite we treat dsn as a file path. Ensure data dir exists.
    if err := ensureDir("./data"); err != nil {
        return nil, err
    }

    // Normalize DSN to a file: URI with mode=rwc when appropriate.
    if dsn == ":memory:" {
        // use in-memory
    } else {
        if !strings.HasPrefix(dsn, "file:") {
            // treat as relative path
            // allow query params if provided
            base := dsn
            // create file if missing and ensure writable
            abs, err := filepath.Abs(base)
            if err != nil {
                return nil, fmt.Errorf("abs path: %w", err)
            }
            if err := ensureFileWritable(abs); err != nil {
                return nil, fmt.Errorf("db file not writable: %w", err)
            }
            // rebuild DSN as file:abs?mode=rwc
            dsn = "file:" + abs + "?mode=rwc"
        } else {
            // file: URI — verify path portion is writable
            u, err := url.Parse(dsn)
            if err == nil {
                if u.Path != "" {
                    abs, err := filepath.Abs(u.Path)
                    if err == nil {
                        if err := ensureFileWritable(abs); err != nil {
                            return nil, fmt.Errorf("db file not writable: %w", err)
                        }
                        // preserve query
                        if u.RawQuery == "" {
                            dsn = "file:" + abs + "?mode=rwc"
                        } else {
                            dsn = "file:" + abs + "?" + u.RawQuery
                        }
                    }
                }
            }
        }
    }

    db, err := sql.Open("sqlite", dsn)
    if err != nil {
        return nil, err
    }
    db.SetMaxOpenConns(1)
    db.SetConnMaxLifetime(30 * time.Minute)
    db.SetConnMaxIdleTime(5 * time.Minute)
    if err := db.Ping(); err != nil {
        return nil, err
    }
    return db, nil
}
