# NetApp Kubernetes Service Go SDK

Version: nks-sdk-go **2.0.7**

The NetApp Kubernetes Service software development kit for [Go](https://www.golang.org/) provides you with access to the NetApp Kubernetes Service API. It is designed for developers who are building applications in Go.

In order to use the client from the NetApp Kubernetes Service Go SDK, you must provide a NetApp Kubernetes Service API token and endpoint url.

#### Installation

Install the Go language from from the official [Go installation](https://golang.org/doc/install) guide or using your Linux distribution package management system.

The `GOPATH` environment variable specifies the location of your Go workspace. It is likely the only environment variable you will need to set when developing Go code. This is an example of pointing to a workspace configured under your home directory:

```
mkdir -p ~/go/bin
export GOPATH=~/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

The following `go` command will download `nks-sdk-go` to your configured `GOPATH`:

```go
go get "github.com/NetApp/nks-sdk-go"
```

The source code of the package will be located here:

    $GOPATH/src/github.com/NetApp/nks-sdk-go

## Library

Include the NetApp Kubernetes Service SDK for Go like any other Go library. For example, create main package file *example.go*:

```go
package main

import (
	"fmt"
)

func main() {
}
```

Include the NetApp Kubernetes Service SDK for Go under the list of imports.

```go
import(
	"fmt"    
	nks "github.com/NetApp/nks-sdk-go"
)
```

#### Authentication

Add your NetApp Kubernetes Service API token and endpoint URL to the client connection.

```go
client := nks.NewClient("token", "endpoint")
```

It might be necessary to accept credentials through environment variables in a containerized environment.

Set your environment variables in your shell.

```
export NKS_API_TOKEN="YOUR TOKEN HERE"
export NKS_API_URL="YOUR ENDPOINT URL HERE"
```

Now you can use a helper function to get a client instance with environment variables.

```go
import (
	"fmt"
	"os"
	nks "github.com/NetApp/nks-sdk-go"
)

func main() {
	client, err := nks.NewClientFromEnv()
...
```

**Caution**: You will want to ensure you follow security best practices when using credentials within your code or stored in a file.
-----------------

## Examples

The NetApp Kubernetes Service SDK for Go comes with several example programs to demonstrate how most major operations can be performed, from listing organizations and nodes, to building clusters in various cloud ecosystems.  The examples will be located in:

github.com/NetApp/nks-sdk-go/example
