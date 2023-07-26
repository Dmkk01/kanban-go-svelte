# Kanban Backend

The backend is written in Go and uses the Echo framework. It is responsible for serving the frontend, providing the API and managing the database.


## Database

Install migrate CLI

- Windows: `scoop install migrate`
- Mac: `brew install golang-migrate`


- Create migration files
`migrate create -ext sql -dir backend/cmd/db/migration/ -seq init`

- Apply migrations (local)
`migrate -path backend/cmd/db/migration/ -database "postgresql://postgres:password@localhost:6500/mydb?sslmode=disable" -verbose up`

- Rollback migrations (local)
`migrate -path backend/cmd/db/migration/ -database "postgresql://postgres:password@localhost:6500/mydb?sslmode=disable" -verbose down`

- Resolve migration errors (local)
`migrate -path backend/cmd/db/migration/ -database "postgresql://postgres:password@localhost:6500/mydb?sslmode=disable" force <VERSION>`

