[![GoDoc](https://godoc.org/github.com/FrozenKP/gotor?status.svg)](https://godoc.org/github.com/FrozenKP/gotor)
[![MIT Licence](https://badges.frapsoft.com/os/mit/mit.svg?v=103)](https://opensource.org/licenses/mit-license.php)

# gotor
a tor client package using golang

## Install
```go
go get github.com/FrozenKP/gotor
```

## Example
```go
package main

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/FrozenKP/gotor"
)

// URL to fetch
var webUrl string = "http://www.google.com"

func main() {
        // get tor client
	client := gotor.New(time.Second * 30)
	
	// Make request
	resp, err := client.Get(webUrl)
	if err != nil {
		log.Fatal("Error making GET request.", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body of response.", err)
	}
	log.Println(string(body))
	log.Println("Return status code:", resp.StatusCode)
}
```
