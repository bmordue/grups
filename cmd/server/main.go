package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/go-chi/chi/v5"
    _ "github.com/jackc/pgx/v5/stdlib"

    "github.com/example/grups/internal/db"
    "github.com/example/grups/internal/handlers"
)

func main() {
    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        host := os.Getenv("PG_HOST")
        port := os.Getenv("PG_PORT")
        user := os.Getenv("PG_USER")
        pass := os.Getenv("PG_PASSWORD")
        name := os.Getenv("PG_DB")
        if port == "" {
            port = "5432"
        }
        dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pass, host, port, name)
    }

    conn, err := db.Connect(dsn)
    if err != nil {
        log.Fatalf("db connect: %v", err)
    }
    defer conn.Close()

    r := chi.NewRouter()
    handlers.Register(r, conn)

    srv := &http.Server{
        Addr:         ":8080",
        Handler:      r,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
    }

    log.Printf("listening on %s", srv.Addr)
    log.Fatal(srv.ListenAndServe())
}
