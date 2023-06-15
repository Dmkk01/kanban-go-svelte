FROM golang:latest

WORKDIR /kanban-backend

COPY go.mod go.sum ./
RUN go mod download

COPY *.go .

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -polling -build="go build main.go" -command="./main"