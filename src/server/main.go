package server

import (
	"fmt"

	"github.com/cateiru/zatsunen/src/config"
)

func RunServer(mode string, path string) {
	configPerEnv := config.GetConfig(mode, path)

	fmt.Println("Mode:", configPerEnv.GetMode())
}
