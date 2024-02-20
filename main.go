package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"os"
	"time"
)

type ServerConfig struct {
	Port          int    `yaml:"port"`
	ServerPath    string `yaml:"server_path"`
	StaticContent *struct {
		RootPath  string `yaml:"root_path"`
		IndexName string `yaml:"index_name"`
	} `yaml:"static_content"`
}

func main() {
	f, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	serverConfig := ServerConfig{}
	err = yaml.Unmarshal(f, &serverConfig)
	if err != nil {
		log.Fatal(err)
	}

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", serverConfig.Port),
		Handler:        http.FileServer(http.Dir(serverConfig.StaticContent.RootPath)),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
