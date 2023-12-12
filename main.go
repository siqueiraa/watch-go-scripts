package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Processes map[string]string `yaml:"processes"`
}

func loadConfig(file string) (Config, error) {
	var config Config

	yamlFile, err := os.ReadFile(file)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func isProcessRunning(processName string) bool {
	cmd := exec.Command("pgrep", "-fl", processName)
	output, err := cmd.Output()

	if err != nil {
		return false
	}

	return len(output) > 0
}

func restartProcess(command string) {
	fmt.Printf("Restarting process: %s\n", command)
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()

	if err != nil {
		fmt.Printf("Error restarting process: %v\n", err)
	}
}

func main() {
	configFile := "commands.yaml" // Adjust this to your YAML file name

	for {
		config, err := loadConfig(configFile)
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			os.Exit(1)
		}

		for processName, command := range config.Processes {
			if !isProcessRunning(processName) {
				log.Println("Starting:", processName)
				restartProcess(command)
			}
		}
		time.Sleep(10 * time.Second)
	}
}
