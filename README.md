# Portman

A terminal UI for viewing and managing network ports in use on your system.

## Overview

Portman is a lightweight Go application that provides an interactive terminal interface to display all open network ports on your machine. It uses `lsof` to gather port information and presents it in a sortable, navigable table.

## Features

- **Real-time port listing**: View all processes currently using network ports on your system
- **Interactive UI**: Built with [tview](https://github.com/rivo/tview) for a responsive terminal experience
- **Port management**: Kill processes directly from the UI
- **Live refresh**: Reload port data to see changes
- **Minimal dependencies**: Leverages standard Go libraries and two focused TUI dependencies

## Data Displayed

For each open port, Portman displays:
- **Process**: Name of the process using the port
- **PID**: Process ID
- **Address**: IP address and port number
- **Protocol**: Transport protocol (TCP/UDP)
- **User**: User account running the process

## Keyboard Controls

- **q**: Quit the application
- **d**: Kill the selected process (sends SIGINT)
- **r**: Reload the port list

## Requirements

- macOS, Linux, or BSD (requires `lsof` command)
- Go 1.25.5 or later (for building from source)

## Installation

```bash
go install github.com/zerenadam/portman/cmd/portman@latest
```

## Running

```bash
# Using make
make run

# Or directly with Go
go run cmd/portman/main.go
```

## How It Works

1. **Port Discovery**: Uses `lsof -i -P -n` to list all internet-connected processes
2. **Parsing**: Parses the output into a structured format with process, PID, user, protocol, and address
3. **UI Rendering**: Displays data in an interactive table with header and footer information
4. **Process Management**: Allows terminating processes via SIGINT signal

## Architecture

```
portman/
├── cmd/portman/main.go      # Entry point
└── internal/
    ├── ui/                  # Terminal UI components
    │   ├── app.go          # Main app struct and lifecycle
    │   ├── header.go       # Header display
    │   ├── footer.go       # Footer with keyboard shortcuts
    │   └── table.go        # Port table and interactions
    └── harbor/
        └── ports.go        # Port discovery via lsof
```

## Dependencies

- [tview](https://github.com/rivo/tview) - Terminal UI library
- [tcell](https://github.com/gdamore/tcell) - Low-level terminal handling (used by tview)

## License

See LICENSE file for details.
