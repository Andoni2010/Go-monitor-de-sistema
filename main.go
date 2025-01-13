package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

const interval = 10 * time.Second

func handleError(err error, sucessMessage string) {
	if err != nil {
		color.Red("Error: %s. Verifica permisos o rutas", err)
	} else {
		color.Green("Success: %s", sucessMessage)
	}
}

func getSystemStats() string {
	usage, err := cpu.Percent(0, false)
	cpuData := fmt.Sprintf("CPU Usage: %.2f%%", usage[0])
	handleError(err, cpuData)

	memStats, err := mem.VirtualMemory()
	ramData := fmt.Sprintf("RAM Usage: %.2f%% (%d MB / %d MB)",
		memStats.UsedPercent, memStats.Used/1024/1024, memStats.Total/1024/1024)
	handleError(err, ramData)

	diskStats, err := disk.Usage("/")
	diskData := fmt.Sprintf("Disk Usage: %.2f%% (%d MB / %d MB)",
		diskStats.UsedPercent, diskStats.Used/1024/1024/1024, diskStats.Total/1024/1024/1024)
	handleError(err, diskData)

	return fmt.Sprintf(`
		<h1>Monitoreo del Sistema</h1>
		<p><strong>%s</strong></p>
		<p><strong>%s</strong></p>
		<p><strong>%s</strong></p>
	`, cpuData, ramData, diskData)
}

func systemStatsHandler(w http.ResponseWriter, r *http.Request) {
	statsHTML := getSystemStats()

	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html lang="es">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Monitor del Sistema</title>
			<meta http-equiv="refresh" content="5">
		</head>
		<body>
			%s
		</body>
		</html>
	`, statsHTML)
}
func main() {

	go func() {
		http.HandleFunc("/", systemStatsHandler)
		fmt.Println("Servidor corriendo en http://localhost:8000")
		if err := http.ListenAndServe(":8000", nil); err != nil {
			color.Red("Error al iniciar el servidor: %s", err)
		}
	}()

	for {
		fmt.Printf("Cargando datos... \n")
		monitorCPU()
		monitorRAM()
		monitorDISK()
		time.Sleep(interval)
	}
}

func monitorCPU() {
	usage, err := cpu.Percent(0, false)
	if len(usage) > 0 {
		handleError(err, fmt.Sprintf("CPU Usage:  %.2f%%", usage[0]))
	} else {
		handleError(fmt.Errorf("no cpu usage data"), "")
	}
}
func monitorRAM() {
	memStats, err := mem.VirtualMemory()

	handleError(err, fmt.Sprintf("RAM Usage:  %.2f%% (%d MB / %d MB)",
		memStats.UsedPercent, memStats.Used/1024/1024, memStats.Total/1024/1024))
}
func monitorDISK() {
	diskStats, err := disk.Usage("/")

	handleError(err, fmt.Sprintf("DISK Usage:  %.2f%% (%d GB / %d GB)\n",
		diskStats.UsedPercent, diskStats.Used/1024/1024/1024, diskStats.Total/1024/1024/1024))
}
