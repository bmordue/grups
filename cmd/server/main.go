package main

import (
    "log"
    "net/http"
    "os"
    "time"

    "github.com/go-chi/chi/v5"

    "github.com/example/grups/internal/db"
    "github.com/example/grups/internal/handlers"
)

func main() {
    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        dsn = "./data/grups.db"
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
