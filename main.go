package main

import (
	"log"
	"net/http"
	"statics/version"

	"github.com/kelseyhightower/envconfig"
)

// Config задает параметры конфигурации приложения
type Config struct {
	Port        string `envconfig:"PORT" default:"8080"`
	StaticsPath string `envconfig:"STATICS_PATH" default:"./static"`
}

func main() {
	config := new(Config)
	err := envconfig.Process("", config)
	if err != nil {
		log.Fatalf("Can't process config: %v", err)
	}

	http.HandleFunc("/__heartbeat__", heartbeatHandler)
	http.HandleFunc("/__version__", versionHandler)

	fs := http.FileServer(http.Dir(config.StaticsPath))
	http.Handle("/", fs)

	err = http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		log.Fatalf("Error while serving: %v", err)
	}
}

func heartbeatHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func versionHandler(w http.ResponseWriter, req *http.Request) {
	_, _ = w.Write([]byte(`{"version":"` + version.Version + `"}`))
}
