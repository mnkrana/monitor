package display

import (
	"fmt"
	"strings"
	"time"

	"github.com/mnkrana/monitor/internal/collector"
	"github.com/rivo/tview"
)

type Dashboard struct {
	app        *tview.Application
	grid       *tview.Grid
	cpuText    *tview.TextView
	ramText    *tview.TextView
	portsTable *tview.Table
	sshTable   *tview.Table
	netText    *tview.TextView
	statusBar  *tview.TextView
}

func NewDashboard() *Dashboard {
	d := &Dashboard{
		app:        tview.NewApplication(),
		cpuText:    tview.NewTextView().SetDynamicColors(true).SetTextAlign(tview.AlignLeft),
		ramText:    tview.NewTextView().SetDynamicColors(true).SetTextAlign(tview.AlignLeft),
		portsTable: tview.NewTable().SetBorders(false).SetSelectable(false, false),
		sshTable:   tview.NewTable().SetBorders(false).SetSelectable(false, false),
		netText:    tview.NewTextView().SetDynamicColors(true).SetTextAlign(tview.AlignLeft),
		statusBar:  tview.NewTextView().SetDynamicColors(true).SetTextAlign(tview.AlignCenter),
	}

	d.setupGrid()
	return d
}

func (d *Dashboard) setupGrid() {
	d.cpuText.SetBorder(true).SetTitle(" CPU Usage ")
	d.ramText.SetBorder(true).SetTitle(" RAM Usage ")
	d.portsTable.SetBorder(true).SetTitle(" Open Ports ")
	d.netText.SetBorder(true).SetTitle(" Network Speed ")
	d.sshTable.SetBorder(true).SetTitle(" SSH Sessions ")

	header := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true).
		SetText("[::b]╔═══════════════════════════════════════════════════════════════════════╗\n[::b]║                    MONITOR - System Dashboard                              ║\n[::b]╚═══════════════════════════════════════════════════════════════════════╝")

	topRow := tview.NewFlex().
		AddItem(d.cpuText, 0, 1, false).
		AddItem(d.ramText, 0, 1, false)

	bottomRight := tview.NewFlex().
		AddItem(d.netText, 0, 1, false).
		AddItem(d.sshTable, 0, 1, false)

	d.grid = tview.NewGrid().
		SetRows(3, 0, 0, 3).
		SetColumns(0, 0)
	d.grid.AddItem(header, 0, 0, 1, 2, 0, 0, false).
		AddItem(topRow, 1, 0, 1, 2, 0, 0, false).
		AddItem(d.portsTable, 2, 0, 1, 1, 0, 0, false).
		AddItem(bottomRight, 2, 1, 1, 1, 0, 0, false).
		AddItem(d.statusBar, 3, 0, 1, 2, 0, 0, false)
}

func (d *Dashboard) update(stats *collector.SystemStats) {
	d.app.QueueUpdateDraw(func() {
		d.updateCPU(stats)
		d.updateRAM(stats)
		d.updatePorts(stats)
		d.updateSSH(stats)
		d.updateNetwork(stats)
		d.updateStatus()
	})
}

func (d *Dashboard) updateCPU(stats *collector.SystemStats) {
	percent := stats.CPUPercent
	barLen := 30
	filled := int(percent / 100 * float64(barLen))
	if filled > barLen {
		filled = barLen
	}

	bar := strings.Repeat("█", filled) + strings.Repeat("░", barLen-filled)
	color := "green"
	if percent > 80 {
		color = "red"
	} else if percent > 60 {
		color = "yellow"
	}

	d.cpuText.SetText(fmt.Sprintf("[%s]%s[white]\n\n%.1f%%", color, bar, percent))
}

func (d *Dashboard) updateRAM(stats *collector.SystemStats) {
	percent := stats.RAMPercent
	barLen := 30
	filled := int(percent / 100 * float64(barLen))
	if filled > barLen {
		filled = barLen
	}

	bar := strings.Repeat("█", filled) + strings.Repeat("░", barLen-filled)
	color := "green"
	if percent > 85 {
		color = "red"
	} else if percent > 70 {
		color = "yellow"
	}

	used := collector.FormatBytes(stats.RAMUsed)
	total := collector.FormatBytes(stats.RAMTotal)

	d.ramText.SetText(fmt.Sprintf("[%s]%s[white]\n\n%.1f%% - %s / %s", color, bar, percent, used, total))
}

func (d *Dashboard) updatePorts(stats *collector.SystemStats) {
	d.portsTable.Clear()
	d.portsTable.SetCell(0, 0, tview.NewTableCell("[::b]Port").SetSelectable(false))
	d.portsTable.SetCell(0, 1, tview.NewTableCell("[::b]Protocol").SetSelectable(false))
	d.portsTable.SetCell(0, 2, tview.NewTableCell("[::b]Program").SetSelectable(false))

	for i, port := range stats.OpenPorts {
		if i >= 15 {
			break
		}
		row := i + 1
		d.portsTable.SetCell(row, 0, tview.NewTableCell(fmt.Sprintf("%d", port.Port)).SetSelectable(false))
		d.portsTable.SetCell(row, 1, tview.NewTableCell(port.Protocol).SetSelectable(false))
		d.portsTable.SetCell(row, 2, tview.NewTableCell(port.Program).SetSelectable(false))
	}
}

func (d *Dashboard) updateSSH(stats *collector.SystemStats) {
	d.sshTable.Clear()
	d.sshTable.SetCell(0, 0, tview.NewTableCell("[::b]Local").SetSelectable(false))
	d.sshTable.SetCell(0, 1, tview.NewTableCell("[::b]Remote").SetSelectable(false))
	d.sshTable.SetCell(0, 2, tview.NewTableCell("[::b]User").SetSelectable(false))
	d.sshTable.SetCell(0, 3, tview.NewTableCell("[::b]State").SetSelectable(false))

	if len(stats.SSHSessions) == 0 {
		d.sshTable.SetCell(1, 0, tview.NewTableCell("[gray]No active SSH sessions"))
	} else {
		for i, session := range stats.SSHSessions {
			if i >= 10 {
				break
			}
			row := i + 1
			d.sshTable.SetCell(row, 0, tview.NewTableCell(session.LocalAddr))
			d.sshTable.SetCell(row, 1, tview.NewTableCell(session.RemoteAddr))
			d.sshTable.SetCell(row, 2, tview.NewTableCell(session.User))
			d.sshTable.SetCell(row, 3, tview.NewTableCell(session.State))
		}
	}
}

func (d *Dashboard) updateNetwork(stats *collector.SystemStats) {
	upload := collector.FormatSpeed(stats.NetUpload)
	download := collector.FormatSpeed(stats.NetDownload)

	d.netText.SetText(fmt.Sprintf("[green]↓ Download: [white]%s\n\n[blue]↑ Upload: [white]%s", download, upload))
}

func (d *Dashboard) updateStatus() {
	ips := collector.GetLocalIPs()
	ipStr := strings.Join(ips, ", ")
	if ipStr == "" {
		ipStr = "127.0.0.1"
	}
	d.statusBar.SetText(fmt.Sprintf("[gray]Press Ctrl+C to exit | Local IPs: [white]%s[gray] | Updated: [white]%s", ipStr, time.Now().Format("15:04:05")))
}

func (d *Dashboard) Run() error {
	go func() {
		for {
			stats, err := collector.Collect()
			if err == nil {
				d.update(stats)
			}
			time.Sleep(2 * time.Second)
		}
	}()

	return d.app.SetRoot(d.grid, true).EnableMouse(true).Run()
}
