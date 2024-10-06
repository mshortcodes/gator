package main

import (
	"fmt"
	"gator/internal/config"
	"log"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("%+v\n", cfg)

	err = cfg.SetUser("Michael")
	if err != nil {
		log.Fatalf("error setting user: %v", err)
	}
	fmt.Printf("%+v\n", cfg)
}
