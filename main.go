package main

import (
	"github.com/BruggerPedro/Gopportunities.git/config"
	"github.com/BruggerPedro/Gopportunities.git/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Inicialize()
}
