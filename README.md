# OBS Scene Switcher

A simple Go utility to toggle between two OBS Studio scenes quickly.

## Features

- Switch between two predefined scenes in OBS Studio
- Lightweight and easy to use
- Uses OBS WebSocket for scene management

## Prerequisites

- OBS Studio
- OBS WebSocket plugin installed
- Go programming environment

## Installation

1. Clone the repository:

```bash
git clone https://github.com/your-username/obs_switchscene.git
cd obs_switchscene
```

2. Build the project:

```bash
go build
```

## Usage

```bash
./obs_switchscene <scene1> <scene2>
```

### Example

```bash
./obs_switchscene "Gaming" "Streaming"
```

This will switch between the "Gaming" and "Streaming" scenes.

## Dependencies

- [go-obs-websocket](https://github.com/christopher-dG/go-obs-websocket)

## Configuration

- Ensure OBS WebSocket is running on `localhost:4444`
- Provide two scene names as command-line arguments

## License

MIT License

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
