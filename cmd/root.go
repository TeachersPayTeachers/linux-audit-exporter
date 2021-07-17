// Package cmd parses CLI arguments and starts Prometheus HTTP server.
package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"

	"github.com/TeachersPayTeachers/linux-audit-exporter/internal/audit"
	"github.com/TeachersPayTeachers/linux-audit-exporter/internal/exporter"
)

var (
	// nolint:gochecknoglobals
	healthPath string
	// nolint:gochecknoglobals
	listenAddress string
	// nolint:gochecknoglobals
	metricsPath string
	// nolint:exhaustivestruct,gochecknoglobals
	rootCmd = &cobra.Command{
		Use:   "linux-audit-exporter",
		Short: "Export Linux audit status as Prometheus metrics",
		Run:   run,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// nolint:gochecknoinits
func init() {
	rootCmd.Flags().StringVar(&healthPath, "health-path", "/healthz", "health path")
	rootCmd.Flags().StringVar(&listenAddress, "listen-address", "0.0.0.0:9090", "listen address")
	rootCmd.Flags().StringVar(&metricsPath, "metrics-path", "/metrics", "metrics path")
}

func run(cmd *cobra.Command, args []string) {
	a, err := audit.New()
	if err != nil {
		log.Fatalf("Error initializing audit: %s", err)
	}

	e := exporter.New(a)

	err = prometheus.Register(e)
	if err != nil {
		log.Fatalf("Error registering exporter: %s", err)
	}

	http.Handle(metricsPath, promhttp.Handler())
	http.HandleFunc(healthPath, func(w http.ResponseWriter, r *http.Request) {
		if _, err = w.Write([]byte(`{"status":"healthy"}`)); err != nil {
			log.Printf("Error sending response body: %s", err)
		}
	})

	err = http.ListenAndServe(listenAddress, nil)
	if err != nil {
		log.Fatalf("Error listening on %s: %s", listenAddress, err)
	}
}
