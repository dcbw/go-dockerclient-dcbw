# go-dockerclient

[![Drone](https://drone.io/github.com/dcbw/go-dockerclient-dcbw/status.png)](https://drone.io/github.com/dcbw/go-dockerclient-dcbw/latest)
[![Travis](https://img.shields.io/travis/dcbw/go-dockerclient-dcbw.svg?style=flat-square)](https://travis-ci.org/dcbw/go-dockerclient-dcbw)
[![GoDoc](https://img.shields.io/badge/api-Godoc-blue.svg?style=flat-square)](https://godoc.org/github.com/dcbw/go-dockerclient-dcbw)

This package presents a client for the Docker remote API. It also provides
support for the extensions in the [Swarm API](https://docs.docker.com/swarm/API/).

For more details, check the [remote API documentation](http://docs.docker.com/en/latest/reference/api/docker_remote_api/).

## Example

```go
package main

import (
	"fmt"

	"github.com/dcbw/go-dockerclient-dcbw"
)

func main() {
	endpoint := "unix:///var/run/docker.sock"
	client, _ := docker.NewClient(endpoint)
	imgs, _ := client.ListImages(docker.ListImagesOptions{All: false})
	for _, img := range imgs {
		fmt.Println("ID: ", img.ID)
		fmt.Println("RepoTags: ", img.RepoTags)
		fmt.Println("Created: ", img.Created)
		fmt.Println("Size: ", img.Size)
		fmt.Println("VirtualSize: ", img.VirtualSize)
		fmt.Println("ParentId: ", img.ParentID)
	}
}
```

## Using with TLS

In order to instantiate the client for a TLS-enabled daemon, you should use NewTLSClient, passing the endpoint and path for key and certificates as parameters.

```go
package main

import (
	"fmt"

	"github.com/dcbw/go-dockerclient-dcbw"
)

func main() {
	endpoint := "tcp://[ip]:[port]"
	path := os.Getenv("DOCKER_CERT_PATH")
	ca := fmt.Sprintf("%s/ca.pem", path)
	cert := fmt.Sprintf("%s/cert.pem", path)
	key := fmt.Sprintf("%s/key.pem", path)
	client, _ := docker.NewTLSClient(endpoint, cert, key, ca)
	// use client
}
```

If using [docker-machine](https://docs.docker.com/machine/), or another application that exports environment variables
`DOCKER_HOST, DOCKER_TLS_VERIFY, DOCKER_CERT_PATH`, you can use NewClientFromEnv.


```go
package main

import (
	"fmt"

	"github.com/dcbw/go-dockerclient-dcbw"
)

func main() {
	client, _ := docker.NewClientFromEnv()
	// use client
}
```

See the documentation for more details.

## Developing

All development commands can be seen in the [Makefile](Makefile).

Commited code must pass:

* [golint](https://github.com/golang/lint)
* [go vet](https://godoc.org/golang.org/x/tools/cmd/vet)
* [gofmt](https://golang.org/cmd/gofmt)
* [go test](https://golang.org/cmd/go/#hdr-Test_packages)

Running `make test` will check all of these. If your editor does not automatically call gofmt, `make fmt` will format all go files in this repository.
