# lwapi
LeaseWeb Cloud API library in Golang, you can find the full API documentatoin [here](http://developer.leaseweb.com/cloudapi-docs/?shell#get-the-datatraffic-metrics)


[![Build Status](https://travis-ci.org/OpenSourceZombie/lwapi.svg?branch=master)](https://travis-ci.org/OpenSourceZombie/lwapi) [![GoDoc](https://godoc.org/github.com/OpenSourceZombie/lwapi?status.svg)](https://godoc.org/github.com/OpenSourceZombie/lwapi)   [![codecov](https://codecov.io/gh/OpenSourceZombie/lwapi/branch/master/graph/badge.svg)](https://codecov.io/gh/OpenSourceZombie/lwapi)

## Install
    go get github.com/OpenSourceZombie/lwapi

### Simple Example:
```go
package main

import (
	"fmt"
	"log"
	"github.com/OpenSourceZombie/lwapi"
)
func main() {
	    lwclient := lwapi.LW{
        AuthToken: "****-****-****-****",
    }
	vServerslList, err := lwclient.GetVirtualServersList()
	if err != nil {
		log.Println(err)
	}
	for k, vServer := range vServerslList.VirtualServers {
		fmt.Printf("%d.\t%s\t%s\n", k, vServer.Reference, vServer.Ips[0].IP)
	}
}
```
