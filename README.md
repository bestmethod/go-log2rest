# go-log2rest
A Go Remote logging module, to be used with go-rest2log (github.com/bestmethod/go-rest2log)

## Usage
#### Get it
```
go get github.com/bestmethod/go-log2rest
```

#### Use it
```go
package main

import "github.com/bestmethod/go-log2rest"
import "fmt"

func main() {
	//initialize
	logger := log2rest.Log2Rest{Endpoint:"http://some.logging.location.example.com:80"}
	
	//log and wait for response
	logger.Debug("My log line that will end up on the remote side")
	
	//log in background instead
	go logger.Info("Some log line that will end up getting remotely there in the background")
	
	//actually check if it succeeds
	err := logger.Warn("Something went weirdly wrong...")
	if err != nil {
		fmt.Printf("Could not log lines remotely. Details: %s",err)
	}
}
```

#### Use it - and make it log locally as well as remotely
```go
package main

import "github.com/bestmethod/go-log2rest"
import "github.com/bestmethod/go-logger"
import "fmt"

func main() {
	// initialize local logger
	locallogger := new(Logger.Logger)
   	locallogger.Init("SUBNAME", "SERVICENAME", Logger.LEVEL_DEBUG | Logger.LEVEL_INFO | Logger.LEVEL_WARN, Logger.LEVEL_ERROR | Logger.LEVEL_CRITICAL, Logger.LEVEL_NONE)

	//initialize
	logger := log2rest.Log2Rest{Endpoint:"http://some.logging.location.example.com:80",LocalLogger:locallogger}
	
	//log and wait for response
	logger.Debug("My log line that will end up on the remote side")
	
	//log in background instead
	go logger.Info("Some log line that will end up getting remotely there in the background")
	
	//actually check if it succeeds
	err := logger.Warn("Something went weirdly wrong...")
	if err != nil {
		fmt.Printf("Could not log lines remotely. Details: %s",err)
	}
}
```

#### Functions
```go
log2rest.Log2Rest{Endpoint:string, LocalLogger:Logger.Logger}

Log2Rest.Debug(m string) error {}
Log2Rest.Info(m string) error {}
Log2Rest.Warn(m string) error {}
Log2Rest.Error(m string) error {}
Log2Rest.Critical(m string) error {}
```