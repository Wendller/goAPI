package main

import (
	"os"
	"path/filepath"

	"github.com/Wendller/goexpert/apis/configs"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	serverDir := filepath.Join(currentDir, "cmd", "server")

	err = os.Chdir(serverDir)
	if err != nil {
		panic(err)
	}

	_, err = configs.LoadConfig(serverDir)
	if err != nil {
		panic(err)
	}

	defer os.Chdir(currentDir)
}
