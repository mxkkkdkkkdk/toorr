package main

import (
	"log"

	"github.com/yohix/torcloud/server"
)

var VERSION = "1.0.0"

func main() {
	s := server.Server{
		ConfigPath: "torcloud.json",
		Port:       80,
	}

	if err := s.Run(VERSION); err != nil {
		log.Fatal(err)
	}
}
