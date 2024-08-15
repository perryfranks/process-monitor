# Process Monitor
Web server + CLI tool to monitor long running processes on multiple machines from a single interface. 
Run commands/processes via the CLI tool and their status, running information and ultimately output will be sent to the web server. 

Use this to validate long running processes are running as intended. Such as server workloads, renders, etc.

## Features

## Running the Server: 
Compile with a command like 
```
go build -o monitorServer  ./cmd/web/
```

Then simply run it and visit http://localhost:4000. Processes with connect via the webserver and automatically be monitored

## Running the CLI: 

Compile with a command like: 
```
go build -o monitorwrap ./cmd/cli/
```

Then for the most part you can invoke commands through the cli such as 
``` 
./monitorwrap sleep 20 // useful for checking connections to the web server 
./monitorwrap "sleep 20" // allows for -(-) to be correctly piped to the monitored command not the cli 
./monitorwrap --help // get some real usage instructions 

```

A command is invoked by the process monitor and then status messages are sent to the webserver to handle monitoring. The output from the invoked command can either be wholly captured or dumped to stdout/stderr as normal. Default operation is captured

## Monitoring API 

The API for starting and ending monitors is quite simple and is not tied to the CLI by anything. See monitor-api.md for more information.

## TODO 
- x Fix polling with when inspecting finished procs 
- copy output button 
- Poll rate buttons should change text based on what is going on
- X Add icons for exit code success/fail
- X fix output display. Limit width. Styling more in line with logs. 
- Search
- x Cli check health / shappy exit on no connection
- cli download page? 
- x clear/delete for finished processes
