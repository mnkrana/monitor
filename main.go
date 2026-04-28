package main

import (
	"log"

	"github.com/mnkrana/monitor/internal/display"
)

func main() {
	dashboard := display.NewDashboard()
	if err := dashboard.Run(); err != nil {
		log.Fatalf("Error running dashboard: %v", err)
	}
}
