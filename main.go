package main

import (
	"course-go/config"
	"fmt"
)
func main() {
	config := config.LoadConfig()
	fmt.Printf("Server running on port %s\n", config.Port)
}