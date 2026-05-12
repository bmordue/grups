# grups — minimal Go API with Postgres (dev flake)

Quick start:

1. Enter the Nix dev shell (requires Nix with flakes enabled):

```bash
nix develop
```


2. Setup database (SQLite):

```bash
nix develop
make migrate
```

3. Copy example env (optional):

```bash
cp .env.example .env
export DATABASE_URL="./data/grups.db"
```

4. Run the server:

```bash
make run
```

Endpoints:
- `GET /health` — health check
- `GET /users` — list users (reads `users` table)

Files added:
- [cmd/server/main.go](cmd/server/main.go)
- [internal/db/db.go](internal/db/db.go)
- [internal/handlers/handlers.go](internal/handlers/handlers.go)
- [db/migrations/001_init.sql](db/migrations/001_init.sql)
- [docker-compose.yml](docker-compose.yml)
- [flake.nix](flake.nix)
- [Makefile](Makefile)
