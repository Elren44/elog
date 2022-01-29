# My zap logger config

## Install

```
go get github.com/Elren44/elog
```

## Example  json log

```
package main

import "github.com/Elren44/elog"

func main() {
	log := elog.InitLogger(elog.JsonOutput)
	log.Info("test logs")
}
```

## Example console log

```
package main

import "github.com/Elren44/elog"

func main() {
	log := elog.InitLogger(elog.ConsoleOutput)
	log.Info("test logs")
}
