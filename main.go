package main

import (
	"github.com/salignatmoandal/terralambda/cmd"
	"github.com/salignatmoandal/terralambda/internal/logger"
)

func main() {
	// Initialiser le logger
	logger.Init()
	cmd.Execute()
}
