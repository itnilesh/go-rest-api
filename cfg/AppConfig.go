package cfg

import "github.com/kelseyhightower/envconfig"

// AppConfig - configuration needed to run micro-service
type AppConfig struct {
	Host string `default:"localhost"`
	Post int   `default:"8888"`
}