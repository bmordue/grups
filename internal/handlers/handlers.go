package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"

    "github.com/go-chi/chi/v5"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func Register(r chi.Router, db *sql.DB) {
    r.Get("/health", health)
    r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
        listUsers(w, r, db)
    })
}

func health(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("ok"))
}

func listUsers(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    rows, err := db.Query("SELECT id, name FROM users LIMIT 100")
    if err != nil {
        http.Error(w, "query error", http.StatusInternalServerError)
        return
    }
    defer rows.Close()
    users := []User{}
    for rows.Next() {
        var u User
        if err := rows.Scan(&u.ID, &u.Name); err != nil {
            http.Error(w, "scan error", http.StatusInternalServerError)
            return
        }
        users = append(users, u)
    }
    json.NewEncoder(w).Encode(users)
}
