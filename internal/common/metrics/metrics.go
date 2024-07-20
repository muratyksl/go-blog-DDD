package metrics

import (
	"encoding/json"
	"log"
	"time"
)

// MetricLogger is a simple logger for metrics
type MetricLogger struct {
	logger *log.Logger
}

// NewMetricLogger creates a new MetricLogger
func NewMetricLogger(logger *log.Logger) *MetricLogger {
	return &MetricLogger{logger: logger}
}

// LogRequest logs HTTP request metrics
func (m *MetricLogger) LogRequest(method, endpoint, status string, duration time.Duration) {
	metric := map[string]interface{}{
		"type":      "http_request",
		"method":    method,
		"endpoint":  endpoint,
		"status":    status,
		"duration":  duration.Seconds(),
		"timestamp": time.Now().Format(time.RFC3339),
	}

	jsonMetric, err := json.Marshal(metric)
	if err != nil {
		m.logger.Printf("Error marshaling metric: %v", err)
		return
	}

	m.logger.Println(string(jsonMetric))
}

// Global metric logger
var GlobalMetricLogger *MetricLogger

// InitMetricLogger initializes the global metric logger
func InitMetricLogger(logger *log.Logger) {
	GlobalMetricLogger = NewMetricLogger(logger)
}

// LogHTTPRequest logs an HTTP request metric using the global logger
func LogHTTPRequest(method, endpoint, status string, duration time.Duration) {
	if GlobalMetricLogger != nil {
		GlobalMetricLogger.LogRequest(method, endpoint, status, duration)
	}
}
