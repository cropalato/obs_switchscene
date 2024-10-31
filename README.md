i# OBS Scene Switcher

A robust Go utility to toggle between two OBS Studio scenes with configurable options and improved error handling.

## Features

- Quick switching between two predefined scenes in OBS Studio
- Configurable connection settings (host, port, timeout)
- Scene validation to prevent errors
- Verbose logging option for debugging
- Connection timeout protection
- Lightweight and easy to use
- Uses OBS WebSocket for secure scene management

## Prerequisites

- OBS Studio
- OBS WebSocket plugin installed and configured
- Go programming environment (1.13 or higher recommended)

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

### Basic Usage

```bash
./obs_switchscene <scene1> <scene2>
```

### Advanced Usage with Options

```bash
./obs_switchscene [options] <scene1> <scene2>
```

### Available Options

```
-host string
      OBS WebSocket host (default "localhost")
-port int
      OBS WebSocket port (default 4444)
-timeout duration
      Connection timeout (default 5s)
-verbose
      Enable verbose logging
```

### Examples

Basic scene switching:

```bash
./obs_switchscene "Gaming" "Streaming"
```

Custom host and port:

```bash
./obs_switchscene -host=192.168.1.100 -port=4445 "Gaming" "Streaming"
```

With longer timeout and verbose logging:

```bash
./obs_switchscene -timeout=10s -verbose "Gaming" "Streaming"
```

## Configuration

The program supports various configuration options through command-line flags:

| Flag    | Description                 | Default   |
| ------- | --------------------------- | --------- |
| host    | OBS WebSocket host address  | localhost |
| port    | OBS WebSocket port number   | 4444      |
| timeout | Connection timeout duration | 5s        |
| verbose | Enable detailed logging     | false     |

## Error Handling

The program includes robust error handling for common scenarios:

- Connection timeouts
- Invalid scene names
- OBS connection failures
- Scene switching failures

When running with `-verbose`, detailed error information and operational status will be logged.

## Dependencies

- [go-obs-websocket](https://github.com/christopher-dG/go-obs-websocket)

## Contributing

Contributions are welcome! Here are some ways you can contribute:

- Report bugs
- Suggest new features
- Submit pull requests
- Improve documentation

### Development Setup

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Planned Features

- Configuration file support
- WebSocket authentication
- Scene transition effects
- Connection health monitoring
- Multiple scene group support

## License

MIT License

## Troubleshooting

### Common Issues

1. **Connection Timeout**

   - Verify OBS is running
   - Check if OBS WebSocket plugin is installed
   - Confirm the correct host and port settings

2. **Scene Not Found**

   - Verify scene names match exactly (case-sensitive)
   - Check for extra spaces in scene names
   - Use `-verbose` flag to see available scenes

3. **Port Already in Use**
   - Verify no other instance is running
   - Check if the specified port is available

## Support

For support, please:

1. Check the troubleshooting section
2. Enable verbose logging with `-verbose` flag
3. Open an issue on GitHub with the error logs
