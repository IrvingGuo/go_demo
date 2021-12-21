package main

import (
	"context"
	"resource-plan-improvement/server"
)

func main() {
	server.NewServer(context.Background())
}
