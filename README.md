# watch-go-scripts

A simple Go script to monitor and restart Go processes defined in a YAML configuration file.

## Overview

This script checks if specified Go processes are running, and if not, restarts them. The process names and corresponding restart commands are defined in a YAML configuration file.

## Usage

1. Clone the repository:

   ```bash
   git clone https://github.com/siqueiraa/watch-go-scripts.git
   cd watch-go-scripts
    ```
2. Create a YAML configuration file named commands.yaml with the processes and restart commands:

```bash 
processes:
  go_datafeeder: "./go_datafeeder"
  go_triggers: "./go_triggers"
```

3. Run the script:
```bash 
go run main.go
```
This will check and restart the specified processes.

##Configuration
You can customize the configuration by editing the commands.yaml file.

##Requirements
Go installed on your system.
