# drill
A Go HTTP client library for creating and sending API requests 

## Install
```
go get -u github.com/yadon-infra/drill
```

## Usage
```go
package main

import "github.com/yadon-infra/drill"

func main(){
	drill.Get("https://hogehoge.com")
}
```
