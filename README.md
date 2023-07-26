<p align="center">
  <img src="/images/header.png" />
</p>

# Kanban Go Svelte

A sophisticated Kanban board built with Svelte and Go!

## Features

- JWT Authentication
- Personalized Boards, Columns, Tasks and Tags
- Customizable task management and creation
- Drag and Drop for ordering and moving cards between columns
- Persistent data storage
- Personal settings
- Mobile friendly
- Emojis!

## Structure

The project is split into three parts:
- `frontend` - The Svelte frontend
- `backend` - The Go backend
- `cli` - The Go CLI for scraping additional data from the web


## Technologies

- Frontend: Svelte, Typescript and TailwindCSS
- Backend: Go, Echo and PostgreSQL
- CLI: Go and Colly

All wrapped up in Docker containers!


## Installation

### Prerequisites

- Docker
- Docker Compose

### Steps

1. Clone the repository
2. Run `docker-compose up` in the root directory
3. Migrate database schemas [Instructions](https://github.com/Dmkk01/kanban-go-svelte/tree/main/backend#database)
4. Open `localhost:8001` in your browser (frontend application)
5. Done!


## Screenshots

Login Page
<img src="/images/login.png"/>

Settings Page
<img src="/images/settings.png"/>

Board
<img src="/images/board.png"/>

Edit board
<img src="/images/edit-board.png"/>

Edit task
<img src="/images/edit-task.png"/>

Tags
<img src="/images/tags.png"/>
