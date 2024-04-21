package main

import (
	"fmt"
	"log"

	"github.com/shivamvijaywargi/go-commerce/configs"
	"github.com/shivamvijaywargi/go-commerce/internal/api"
)

func main() {
	cfg, err := configs.SetupEnv()

	fmt.Printf("config: %v\n", cfg)

	if err != nil {
		log.Fatalf("config file is not loaded properly %v\n", err)
	}

	api.StartServer(cfg)
}
