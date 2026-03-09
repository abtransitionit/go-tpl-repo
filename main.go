package main

import (
	"log"
)

func main() {
	if err := run_save_yaml_v7(); err != nil {
		log.Fatalf("loading yaml: %v", err)
	}
}
