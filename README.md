# Todo CLI (Cobra)

A terminal-based todo manager written in Go. Create, list, edit, delete, and toggle tasks from the command line, with persistence to a local JSON file.

Built with [Cobra](https://github.com/spf13/cobra) for the CLI layer and a generic storage layer for JSON read/write.

## Features

- Add todos with a title
- List all todos, or filter by completed / pending status
- Edit a todo's title by index
- Delete a todo by index
- Toggle completion status (tracks `completedAt` timestamp)
- Colored terminal table output
- JSON persistence in `data/todo.json`

## Requirements

- Go 1.25+

## Installation

Clone the repository and build the binary:

```bash
git clone https://github.com/promise111/todo-cli-cobra.git
cd todo-cli-cobra
go build -o todo ./cmd/todo
```

Or run without installing:

```bash
go run ./cmd/todo <command> [args]
```

## Usage

```bash
todo <command> [arguments] [flags]
```

### Commands

| Command | Alias(es) | Description |
|---------|-----------|-------------|
| `add <title>` | `enter` | Add a new todo |
| `list` | `print` | List todos in a table |
| `edit <index>:<title>` | `update` | Update a todo's title by index |
| `delete <index>` | `remove` | Delete a todo by index |
| `toggle <index>` | — | Toggle a todo between completed and pending |

Indexes are **0-based** and match the `s/n` column shown in `list` output.

### Examples

```bash
# Add tasks
todo add "Buy bread"
todo add "Finish homework"

# List all todos
todo list

# List only completed todos
todo list --completed
todo list -c

# List only pending todos
todo list --pending
todo list -p

# Edit todo at index 1
todo edit 1:"Updated title"

# Mark todo at index 0 as complete (or mark complete back to pending)
todo toggle 0

# Delete todo at index 2
todo delete 2
```

### Help

```bash
todo --help
todo add --help
todo list --help
```

## Data Storage

Todos are stored as JSON at `data/todo.json`. The `data/` directory is created automatically on first save.

Each todo is stored with the following fields:

| Field | Type | Description |
|-------|------|-------------|
| `title` | string | Task description |
| `completed` | bool | Whether the task is done |
| `createdAt` | timestamp | When the task was created |
| `completedAt` | timestamp or null | When the task was marked complete |

On every run, the app loads todos from disk, executes the requested command, then saves any changes back to the file.

## Project Structure

```
todo-cli-cobra/
├── cmd/
│   └── todo/
│       └── main.go          # Entry point: load → execute CLI → save
├── data/
│   └── todo.json            # Persisted todos (gitignored)
├── internal/
│   ├── todo.go              # Todo model and business logic
│   ├── storage.go           # Generic JSON file storage
│   ├── commands/
│   │   ├── rootCmd.go       # Root command and Execute()
│   │   ├── addCmd.go
│   │   ├── listCmd.go
│   │   ├── editCmd.go
│   │   ├── deleteCmd.go
│   │   └── toggleCmd.go
│   └── utils/
│       └── todo.go          # Index validation and formatting helpers
├── go.mod
└── README.md
```

## Architecture

```
main (cmd/todo)
  │
  ├─ Load todos from data/todo.json
  ├─ commands.Execute()  →  Cobra CLI (add, list, edit, delete, toggle)
  └─ Save todos to data/todo.json
```

- **`cmd/todo`** — Thin `package main` entry point. Wires storage to the shared `commands.Todos` slice.
- **`internal`** — Core domain logic (`Todos` methods) and generic `Storage[T]` for JSON persistence.
- **`internal/commands`** — Cobra command definitions. All commands operate on the shared `Todos` variable.
- **`internal/utils`** — Shared helpers such as index validation.

## Dependencies

| Package | Purpose |
|---------|---------|
| [spf13/cobra](https://github.com/spf13/cobra) | CLI framework |
| [aquasecurity/table](https://github.com/aquasecurity/table) | Terminal table rendering |
| [liamg/tml](https://github.com/liamg/tml) | Colored terminal output |

## License

MIT
