package main

import (
    "flag"
    "fmt"
    "os"
    "text/template"
   _ "time"

     // The following import is intended for use in generated tests
    _ "github.com/shirou/gopsutil/v3/process" // Import the package

 )
// PrometheusMonitorConfig holds common configuration for process monitoring
type PrometheusMonitorConfig struct {
	ProcessName string
	Interval    string // Change the Interval type to string
	Port        int
}

// PrometheusNginxMonitorConfig holds configuration for monitoring Nginx process
type PrometheusNginxMonitorConfig struct {
	PrometheusMonitorConfig
	MetricName string
	MetricHelp string
}

// processMonitorTemplate is the Go template for the Prometheus listener code
const processMonitorTemplate = `package main

import (
	"net/http"
	"time"
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/v3/process"
)

var (
	{{ .MetricName }} = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "{{ .MetricName }}",
		Help: "{{ .MetricHelp }}",
	})
)

func init() {
	prometheus.MustRegister({{ .MetricName }})
}

func processMonitor() {
	for {
		// Use the gopsutil package to list processes and check if the specified process is running
		processes, err := process.Processes()
		if err != nil {
			log.Printf("Error fetching processes: %s\n", err)
			continue
		}

		var found bool
		for _, p := range processes {
			name, err := p.Name()
			if err != nil {
				continue
			}
			if name == "{{ .ProcessName }}" {
				found = true
				break
			}
		}

		if found {
			{{ .MetricName }}.Set(1)
		} else {
			{{ .MetricName }}.Set(0)
		}

		time.Sleep({{ .Interval }} * time.Second) // Corrected sleep duration format
	}
}

func main() {
	go processMonitor()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":{{ .Port }}", nil))
}
`

func main() {
	processName := flag.String("process-name", "nginx", "Name of the process to monitor")
	interval := flag.Int("interval", 60, "Check interval in seconds") // Change the interval type to int
	port := flag.Int("port", 9090, "Port for Prometheus listener")

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Error: No output file specified")
		os.Exit(1)
	}
	outputFileName := args[len(args)-1]

	config := PrometheusNginxMonitorConfig{
		PrometheusMonitorConfig: PrometheusMonitorConfig{
			ProcessName: *processName,
			Interval:    fmt.Sprintf("%d", *interval), // Format interval as string
			Port:        *port,
		},
		MetricName: "nginx_up",
		MetricHelp: "Indicates if Nginx is running (1) or not (0).",
	}

	tmpl, err := template.New("processMonitor").Parse(processMonitorTemplate)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		os.Exit(1)
	}

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	if err := tmpl.Execute(outputFile, config); err != nil {
		fmt.Println("Error executing template:", err)
		os.Exit(1)
	}

	fmt.Printf("Nginx monitor code generated in %s\n", outputFileName)
}
