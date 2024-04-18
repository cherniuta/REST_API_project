package main

import (
	"fmt"
	"rest_api_project/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Print(cfg)
}
