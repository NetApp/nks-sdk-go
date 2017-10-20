# stackpoint.io Go SDK

The project contains a set of stackpoint type definitions, a
client that interacts with a stackpoint server, and an example
commandline binary.

At the current time, the client library does not cover the full
capabilities of the stackpoint api, and so the library remains
unversioned.  At such point as coverage is roughly complete, a 
library version will be assigned to match the stackpoint api
version.  Although it has to be considered a pre-alpha library,
it is still useful for programming against the stackpoint servers.
Contributions and pull requests are welcome.

In order to use the client from the stackpoint go sdk, you must 
provide a stackpoint api token and an url.

## Library

Incorporate the library into your application as you would normally, 
with the dependency management tool of your choice.

## Command-line interface

### Build

The commandline client, `spcctl`, uses _dep_ to manage dependencies; see https://github.com/golang/dep for more information. To build and install 
the commandline client,
```
go get github.com/StackPointCloud/stackpoint-sdk-go/cmd
cd $GOPATH/src/github.com/StackPointCloud/stackpoint-sdk-go/
dep ensure
GOBIN=$GOPATH/bin go install -v cmd/spcctl.go
```

### Use

The commandline client requires a token, a url, an organization id and a 
cluster id. These values may be set from environment variables or from 
the commandline.  The `get` subcommand simply retrieves stackpoint objects
in a serialized json form to stdout.

In this example, the CLUSTER_API_TOKEN environment variable has been set, and
other values are specified as arguments. The json result can be processed with 
`jq`

```
$ spcctl --url=https://api-staging.stackpoint.io --org=4 --cluster=2904  get nodes | jq '.[] | select(.role=="master")'
{
  "pk": 9409,
  "name": "",
  "cluster": 2904,
  "instance_id": "spcx31lvn3-master-1",
  "role": "master",
  "private_ip": "10.136.46.40",
  "public_ip": "198.211.101.121",
  "platform": "coreos",
  "image": "coreos-stable",
  "location": "nyc1",
  "size": "2gb",
  "state": "running",
  "created": "2017-11-03T15:43:56.695056Z",
  "updated": "2017-11-03T15:43:56.695072Z"
}
```

The CLI is quite limited in functionality. Future extensions will include

- more get-able types to inspect
- creation of clusters and nodes
- interaction with existing clusters