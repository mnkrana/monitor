# Monitor

A lightweight, terminal-based system monitoring CLI tool built in Go. Monitor provides a real-time, live-updating dashboard that displays critical system metrics all in one view.

## Demo

```
╔══════════════════════════════════════════════════════════════════════╗
║                    MONITOR - System Dashboard                              ║
╚══════════════════════════════════════════════════════════════════════╝
┌─────────────────────────┬─────────────────────────────────────┐
│   CPU Usage            │   RAM Usage                         │
│   [████████░░░]      │   [████████████░]                  │
│   78.5%               │   82% - 12.3 GB / 16 GB           │
├─────────────────────────┼─────────────────────────────────────┤
│   Open Ports           │   Network Speed                     │
│   22   (tcp)  sshd   │   ↓ Download: 1.2 MB/s              │
│   80   (tcp)  nginx  │   ↑ Upload:   0.8 MB/s              │
├─────────────────────────┴─────────────────────────────────────┤
│   SSH Sessions                                                     │
│   Local: 192.168.1.5:22 -> Remote: 10.0.0.1:54321          │
└─────────────────────────────────────────────────────────────────────┘
```

## Features

- **CPU Usage** - Real-time CPU utilization with visual progress bar and color-coded thresholds
- **RAM Usage** - Memory consumption with used/total display and percentage
- **Open Ports** - Lists all listening ports with protocol and associated programs
- **Network Speeds** - Live upload/download speed monitoring
- **SSH Sessions** - Active SSH connections tracking
- **Auto-refresh** - Updates every 2 seconds for real-time monitoring

## Installation

### Using Go (Recommended)
```bash
go install github.com/mnkrana/monitor@latest
```

### From Source
```bash
git clone https://github.com/mnkrana/monitor.git
cd monitor
go build -o monitor .
sudo mv monitor /usr/local/bin/
```

### Homebrew (Coming Soon)
```bash
brew install mnkrana/tap/monitor
```

## Usage

Simply run:
```bash
monitor
```

Press `Ctrl+C` to exit.

## Dashboard Layout

```
╔════════════════════════════════════════════════════════════════════╗
║                    MONITOR - System Dashboard                              ║
╚════════════════════════════════════════════════════════════════════╝
┌─────────────────────────┬─────────────────────────────────────┐
│   CPU Usage            │   RAM Usage                         │
│   [████████░░░]      │   [████████████░]                  │
│   78.5%               │   82% - 12.3 GB / 16 GB           │
├─────────────────────────┼─────────────────────────────────────┤
│   Open Ports           │   Network Speed                     │
│   22   (tcp)  sshd   │   ↓ Download: 1.2 MB/s              │
│   80   (tcp)  nginx  │   ↑ Upload:   0.8 MB/s              │
├─────────────────────────┴─────────────────────────────────────┤
│   SSH Sessions                                                     │
│   Local: 192.168.1.5:22 -> Remote: 10.0.0.1:54321          │
└─────────────────────────────────────────────────────────────────────┘
```

## Dependencies

- [gopsutil](https://github.com/shirou/gopsutil) - Cross-platform system metrics
- [tview](https://github.com/rivo/tview) - Terminal UI toolkit

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see LICENSE file for details.

## Author

**Mayank** - [@mnkrana](https://github.com/mnkrana)

---

Made with ❤️ and Go
