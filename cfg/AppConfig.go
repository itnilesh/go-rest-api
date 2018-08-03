package cfg

import "fmt"

// EmpServiceConfig - configuration needed to run micro-service
type EmpServiceConfig struct {
	Host     string `default:"localhost"`
	Port     int    `default:"8888"`
	LogLevel string `default:"info"`
}

// ToString - string representaion of EmpServiceConfig
func (e *EmpServiceConfig) ToString() string {
	s := fmt.Sprintf("Host=%s, Port=%d, LogLevel=%s", e.Host, e.Port, e.LogLevel)
	return s

}
