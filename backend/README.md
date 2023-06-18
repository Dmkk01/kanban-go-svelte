# Kanban Backend


## Commands

https://fly.io/docs/languages-and-frameworks/golang/
`flyctl deploy`
`flyctl status`
`flyctl ips list`


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

