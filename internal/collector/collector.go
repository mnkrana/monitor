package collector

import (
	"fmt"
	"net"
	"strings"
	"time"

	gopsnet "github.com/shirou/gopsutil/v3/net"
	gopsprocess "github.com/shirou/gopsutil/v3/process"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type SystemStats struct {
	CPUPercent    float64
	RAMTotal      uint64
	RAMUsed       uint64
	RAMPercent    float64
	OpenPorts     []PortInfo
	SSHSessions   []SSHInfo
	NetUpload     float64
	NetDownload   float64
}

type PortInfo struct {
	Port     uint32
	Protocol string
	Program  string
}

type SSHInfo struct {
	LocalAddr  string
	RemoteAddr string
	User       string
	State      string
}

var prevNetCounters map[string]gopsnet.IOCountersStat
var lastNetUpdate time.Time

func Collect() (*SystemStats, error) {
	stats := &SystemStats{}

	// CPU
	cpuPercent, err := cpu.Percent(200*time.Millisecond, false)
	if err == nil && len(cpuPercent) > 0 {
		stats.CPUPercent = cpuPercent[0]
	}

	// RAM
	vm, err := mem.VirtualMemory()
	if err == nil {
		stats.RAMTotal = vm.Total
		stats.RAMUsed = vm.Used
		stats.RAMPercent = vm.UsedPercent
	}

	// Open ports
	stats.OpenPorts = getOpenPorts()

	// SSH sessions
	stats.SSHSessions = getSSHSessions()

	// Network speeds
	stats.NetUpload, stats.NetDownload = getNetworkSpeed()

	return stats, nil
}

func getOpenPorts() []PortInfo {
	ports := []PortInfo{}

	connections, err := gopsnet.Connections("all")
	if err != nil {
		return ports
	}

	seen := make(map[uint32]bool)
	for _, conn := range connections {
		if conn.Status == "LISTEN" && conn.Laddr.Port > 0 && !seen[conn.Laddr.Port] {
			seen[conn.Laddr.Port] = true
			program := "unknown"
			if conn.Pid > 0 {
				if p, err := gopsprocess.NewProcess(conn.Pid); err == nil {
					if name, err := p.Name(); err == nil {
						program = name
					}
				}
			}
			ports = append(ports, PortInfo{
				Port:     conn.Laddr.Port,
				Protocol: fmt.Sprintf("%d", conn.Type),
				Program:  program,
			})
		}
	}
	return ports
}

func getSSHSessions() []SSHInfo {
	sessions := []SSHInfo{}

	connections, err := gopsnet.Connections("all")
	if err != nil {
		return sessions
	}

	for _, conn := range connections {
		// Check for SSH connections (port 22) that are established
		if conn.Status == "ESTABLISHED" && (conn.Laddr.Port == 22 || conn.Raddr.Port == 22) {
			sessions = append(sessions, SSHInfo{
				LocalAddr:  fmt.Sprintf("%s:%d", conn.Laddr.IP, conn.Laddr.Port),
				RemoteAddr: fmt.Sprintf("%s:%d", conn.Raddr.IP, conn.Raddr.Port),
				State:      conn.Status,
			})
		}
	}

	// Also check for SSH processes
	processes, _ := gopsprocess.Processes()
	for _, p := range processes {
		name, err := p.Name()
		if err != nil {
			continue
		}
		if strings.Contains(strings.ToLower(name), "sshd") || strings.Contains(strings.ToLower(name), "ssh") {
			connections, _ := p.Connections()
			for _, conn := range connections {
				if conn.Status == "ESTABLISHED" && conn.Raddr.Port > 0 {
					sessions = append(sessions, SSHInfo{
						LocalAddr:  fmt.Sprintf("%s:%d", conn.Laddr.IP, conn.Laddr.Port),
						RemoteAddr: fmt.Sprintf("%s:%d", conn.Raddr.IP, conn.Raddr.Port),
						User:       name,
						State:      conn.Status,
					})
				}
			}
		}
	}

	return sessions
}

func getNetworkSpeed() (upload float64, download float64) {
	counters, err := gopsnet.IOCounters(false)
	if err != nil || len(counters) == 0 {
		return 0, 0
	}

	now := time.Now()
	if prevNetCounters == nil {
		prevNetCounters = make(map[string]gopsnet.IOCountersStat)
		for _, c := range counters {
			prevNetCounters[c.Name] = c
		}
		lastNetUpdate = now
		return 0, 0
	}

	elapsed := now.Sub(lastNetUpdate).Seconds()
	if elapsed <= 0 {
		return 0, 0
	}

	var totalSent, totalRecv uint64
	var prevSent, prevRecv uint64

	for _, c := range counters {
		totalSent += c.BytesSent
		totalRecv += c.BytesRecv
		if prev, ok := prevNetCounters[c.Name]; ok {
			prevSent += prev.BytesSent
			prevRecv += prev.BytesRecv
		}
	}

	upload = float64(totalSent-prevSent) / elapsed
	download = float64(totalRecv-prevRecv) / elapsed

	prevNetCounters = make(map[string]gopsnet.IOCountersStat)
	for _, c := range counters {
		prevNetCounters[c.Name] = c
	}
	lastNetUpdate = now

	return upload, download
}

func FormatBytes(bytes uint64) string {
	if bytes >= 1073741824 {
		return fmt.Sprintf("%.2f GB", float64(bytes)/1073741824)
	} else if bytes >= 1048576 {
		return fmt.Sprintf("%.2f MB", float64(bytes)/1048576)
	} else if bytes >= 1024 {
		return fmt.Sprintf("%.2f KB", float64(bytes)/1024)
	}
	return fmt.Sprintf("%d B", bytes)
}

func FormatSpeed(bytesPerSec float64) string {
	if bytesPerSec >= 1048576 {
		return fmt.Sprintf("%.2f MB/s", bytesPerSec/1048576)
	} else if bytesPerSec >= 1024 {
		return fmt.Sprintf("%.2f KB/s", bytesPerSec/1024)
	}
	return fmt.Sprintf("%.0f B/s", bytesPerSec)
}

func GetLocalIPs() []string {
	var ips []string
	interfaces, err := net.Interfaces()
	if err != nil {
		return ips
	}
	for _, i := range interfaces {
		if i.Flags&net.FlagUp == 0 || i.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ips = append(ips, ipnet.IP.String())
				}
			}
		}
	}
	return ips
}
