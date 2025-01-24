package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

// Intervalo para el monitoreo continuo
const interval = 10 * time.Second

// Manejo de errores con colores
type SystemStats struct {
	CPUUsage  float64 `json:"cpu_usage"`
	RAMUsage  float64 `json:"ram_usage"`
	DISKUsage float64 `json:"disk_usage"`
	TIMEStamp string  `json:"timestammp"`
}

// Manejo de errores con colores
func handleError(err error, sucessMessage string) {
	if err != nil {
		color.Red("Error: %s. Verifica permisos o rutas", err)
	} else {
		color.Green("Success: %s", sucessMessage)
	}
}

// Función para obtener estadísticas del sistema
func getSystemStats() (*SystemStats, error) {
	cpuUsage, err := cpu.Percent(0, false)
	if err != nil {
		return nil, err
	}

	memStats, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	diskStats, err := disk.Usage("/")
	if err != nil {
		return nil, err
	}

	stats := &SystemStats{
		CPUUsage:  cpuUsage[0],
		RAMUsage:  memStats.UsedPercent,
		DISKUsage: diskStats.UsedPercent,
		TIMEStamp: time.Now().Format("2006-01-02 15:04:05"),
	}

	return stats, nil

}

// Servidor Web
func startWebServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		stats, err := getSystemStats()
		if err != nil {
			http.Error(w, "Error obteniendo estadísticas del sistema", http.StatusInternalServerError)
			return
		}

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
				<h1>Monitoreo del Sistema</h1>
				<p><strong>CPU Usage:</strong> %.2f%%</p>
				<p><strong>RAM Usage:</strong> %.2f%%</p>
				<p><strong>Disk Usage:</strong> %.2f%%</p>
				<p><strong>Timestamp:</strong> %s</p>
			</body>
			</html>
		`, stats.CPUUsage, stats.RAMUsage, stats.DISKUsage, stats.TIMEStamp)
	})

	fmt.Println("Servidor corriendo en http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// Generación de Reporte
func generationReport() {
	stats, err := getSystemStats()
	if err != nil {
		handleError(err, "No se pudo crear el archivo del reporte")
		return
	}

	file, err := os.Create("system_report.json")
	if err != nil {
		handleError(err, "No se pudo crear el archivo del reporte")
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(stats)
	handleError(err, "Reporte generado correctamente")
}

// Función Principal
func main() {
	web := flag.Bool("web", false, "Inicia el servidor web")
	report := flag.Bool("report", false, "Genera un reporte JSON del sistema")
	flag.Parse()

	if *web {
		startWebServer()
	} else if *report {
		generationReport()
	} else {
		fmt.Println("Uso:")
		fmt.Println(" --web")
		fmt.Println(" --report")
	}
}
