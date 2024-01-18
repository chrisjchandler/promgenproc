# Prometheus Process Monitor Generator

This Go program generates Prometheus listeners for monitoring the status of various processes on your system. It simplifies the process of creating Prometheus listeners by providing a command-line interface to customize the monitoring parameters.

## Getting Started

### Prerequisites

Before using the Prometheus Process Monitor Generator, ensure you have Go installed on your system.

### Installation

Clone the repository to your local machine:

```bash
git clone <repository-url>
cd <repository-folder>

Build the generator:
go build promgen.go

./promgen -process-name <process-name> -interval <interval> -port <port> <output-file>

<process-name>: Name of the process to monitor.
<interval>: Check interval in seconds.
<port>: Port for the Prometheus listener.
<output-file>: Output file for the generated Prometheus listener code.


Examples
Monitoring Nginx
Generate Prometheus listener for monitoring Nginx with a check interval of 30 seconds and on port 9090:

./promgen -process-name nginx -interval 30 -port 9090 nginx_monitor.go

Monitoring Apache
Generate Prometheus listener for monitoring Apache with a check interval of 60 seconds and on port 9100:

./promgen -process-name apache2 -interval 60 -port 9100 apache_monitor.go

Monitoring Custom Process
Generate Prometheus listener for monitoring a custom process named "myapp" with a check interval of 15 seconds and on port 9200:

./promgen -process-name myapp -interval 15 -port 9200 custom_monitor.go

Running the Generated Prometheus Listener
Once you have generated the Prometheus listener code using the generator, you can run it using the following command:

go run <output-file>

Replace <output-file> with the generated listener file (e.g., nginx_monitor.go, apache_monitor.go, or custom_monitor.go).

Prometheus Configuration
Configure Prometheus to scrape the generated listeners by adding the following to your prometheus.yml:

scrape_configs:
  - job_name: 'process_monitor'
    static_configs:
      - targets: ['localhost:<port>']

eplace <port> with the port specified when generating the listener.

Contributing
Contributions are welcome! If you have suggestions or want to improve the generator, please create an issue or submit a pull request.

License
This project is licensed under the MIT License


