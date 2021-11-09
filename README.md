# My zap logger config

## Install

```
go get github.com/Elren44/elog
```

## Example

```
package main

import "github.com/Elren44/elog"

func main() {
	log := elog.InitLogger()
	log.Info("test logs")
}
```
