# Monitor

A lightweight, terminal-based system monitoring CLI tool built in Go. Monitor provides a real-time, live-updating dashboard that displays critical system metrics all in one view.

## Demo

```
╔═══════════════════════════════════════════════════════════════════╗
║                    MONITOR - System Dashboard                              ║
╚═══════════════════════════════════════════════════════════════════╝
┌─────────────────┬─────────────────┬─────────────────┐
│   CPU Usage    │   RAM Usage     │   Disk Usage    │
│   [█████░]    │   [██████░]    │   [████░]      │
│   78.5%        │   82% - 12 GB  │   45% - 250GB  │
├─────────────────┴─────────────────┴─────────────────┤
│ ┌──────────────┬───────────────┐  ┌──────────────┬───────────────┐│
│ │ Open Ports   │ Top 5 CPU    │  │ Network Speed│ SSH Sessions  ││
│ │ 22 sshd      │ PID  Name  %  │  │ ↓ 1.2 MB/s  │ Local: ...   ││
│ │ 80 nginx     │ 123 chrome 5% │  │ ↑ 0.8 MB/s  │ Remote: ...  ││
│ └──────────────┴───────────────┘  └──────────────┴───────────────┘│
├─────────────────────────────────────────────────────────────────────────┤
│ Top 5 RAM Processes                   │
│ PID  Name     RAM%  RSS             │
│ 456 chrome   12.3% 2.1 GB         │
│ 789 node     8.5%  1.4 GB          │
└─────────────────────────────────────────────────────────────────────────┘
```

## Features

- **CPU Usage** - Real-time CPU utilization with visual progress bar and color-coded thresholds
- **RAM Usage** - Memory consumption with used/total display and percentage
- **Disk Usage** - Storage monitoring with used/total and progress bar
- **Open Ports** - Lists all listening ports with protocol and associated programs
- **Network Speeds** - Live upload/download speed monitoring
- **SSH Sessions** - Active SSH connections tracking
- **Top 5 CPU Processes** - Live view of processes consuming the most CPU
- **Top 5 RAM Processes** - Live view of processes consuming the most memory (with PID, name, %RAM, RSS)
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
╔═══════════════════════════════════════════════════════════════════╗
║                    MONITOR - System Dashboard                              ║
╚═══════════════════════════════════════════════════════════════════╝
┌─────────────────┬─────────────────┬─────────────────┐
│   CPU Usage    │   RAM Usage     │   Disk Usage    │
├─────────────────┴─────────────────┴─────────────────┤
│ ┌──────────────┬───────────────┐  ┌──────────────┬───────────────┐│
│ │ Open Ports   │ Top 5 CPU    │  │ Network Speed│ SSH Sessions  ││
│ └──────────────┴───────────────┘  └──────────────┴───────────────┘│
├─────────────────────────────────────────────────────────────────────────┤
│ Top 5 RAM Processes (full width)                                    │
└─────────────────────────────────────────────────────────────────────────┘
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
