# TODO-List

A simple command-line TODO list application written in Go with PostgreSQL as the backend. This tool helps you manage your tasks efficiently with a terminal-based menu system.

## Features

- Add new tasks 📝
- List all tasks 📋
- Mark tasks as completed ✅
- Undo completed tasks 🔄
- Delete tasks ❌
- Uses PostgreSQL for persistent task storage 🗄️

## Installation

### Prerequisites

- Go 1.23 or later
- PostgreSQL installed and running

### Clone the Repository

```sh
git clone git@github.com:robinho46/TODO-List.git
cd TODO-List
```

### Set Up Database

Ensure you have a PostgreSQL database running and create a database named `todo_list`.

```sh
psql -U your_user -c "CREATE DATABASE todo_list;"
```

Run the setup script to initialize the database:

```sh
./setup.sh
```

## Usage

### Build the Project

```sh
go build -o todo
```

### Run the TODO List CLI

```sh
./todo
```

Follow the interactive menu to add, list, complete, undo, or delete tasks.
