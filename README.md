# Geocube GoLang Client Library

Geocube GoLang Client Library is delivered as an example of Geocube Client.
It can be used as a module in a Golang application or as a Command-Line-Interface.

- [Geocube GoLang Client Library](#geocube-golang-client-library)
  - [Requirements](#requirements)
  - [Installation](#installation)
    - [Geocube-Client-Go](#geocube-client-go)
    - [Command-Line-Interface](#command-line-interface)
    - [Connect to the Geocube](#connect-to-the-geocube)
  - [Documentation](#documentation)
    - [CLI](#cli)
    - [Client](#client)
  - [Status](#status)
  - [Update GRPC Interface](#update-grpc-interface)
  - [Contributing](#contributing)
  - [Licensing](#licensing)
  - [Credits](#credits)

## Requirements
- GoLang 1.16
- An instance of the Geocube Server, its url and, depending on the configuration, its ApiKey

## Installation

### Geocube-Client-Go

```bash
$ GO111MODULE=on go get github.com/airbusgeo/geocube-client-go
```

```golang
...
import (
  "github.com/airbusgeo/geocube-client-go/client"
)
...
```

### Command-Line-Interface

Build the command-line-interface

```bash
$ cd cli && go build
$ ./cli --help
NAME:
   cli - Command-Line-Interface Client for the Geocube

USAGE:
   cli [global options] command [command options] [arguments...]

VERSION:
   0.2.0

DESCRIPTION:
   Command-Line-Interface Client to connect to a Geocube Server (github.com/airbusgeo/geocube)

COMMANDS:
   records, r    manage records
   layouts, l    manage layout
   catalog, c    access catalog
   variables, v  manage variables
   palettes, p   manage palettes
   operation, o  manage datasets
   help, h       Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --srv value     geocube server endpoint (default: "127.0.0.1:8080")
   --apikey value  distribute api key
   --insecure      allow insecure grpc
   --help, -h      show help
   --version, -v   print the version
```


### Connect to the Geocube

The connection is configured using the global options `srv`, `apikey`, `insecure`:
```bash
./cli --srv 127.0.0.1:8080 --insecure
```

## Documentation

### CLI

Commands are grouped by kind of operations. The entities accepts the [CRUD(L)](https://en.wikipedia.org/wiki/Create,_read,_update_and_delete) operations (all operations are not implemented yet):
- `records` : commands to handle records
- `layouts` : commands to handle layouts
- `variables` : commands to handle variables
- `palettes` : commands to handle palettes (create)
- `operation` : commands to handle datasets (indexation, delation, consolidation...)
- `catalog` : commands to access catalog (get cube of data)
Please follow the [Data Access Jupyter notebook](https://github.com/airbusgeo/geocube-client-python/Jupyter/Geocube-Client-DataAccess.ipynb)


Documentation of each command is available with the following command:

```bash
$ ./cli command --help
```

### Client
Documentation of the client will be available soon. Please refer to the GRPC documentation.

## Status

Geocube-Client-Go is under development. The API might evolve in backwards incompatible ways until essential functionality is covered.

## Update GRPC Interface

Geocube-Client-Go uses the protobuf GRPC interface, automatically generated from the protobuf files provided by [the Geocube](https://github.com/airbusgeo/geocube).

The `pb` files can be generated using the module `protoc-gen-go` (`go get -u github.com/golang/protobuf/protoc-gen-go`), the geocube [protobuf folder](https://github.com/airbusgeo/geocube/api/v1/) and the following command:

```bash
protoc -I <geocube_folder>/api/v1/ --go_out=plugins=grpc:. pb/geocube.proto pb/catalog.proto pb/records.proto pb/dataformat.proto pb/variables.proto pb/layouts.proto pb/operations.proto
```

## Contributing

Contributions are welcome. Please read the [contributing guidelines](https://github.com/airbusgeo/geocube-client-go/CONTRIBUTING.MD) before submitting fixes or enhancements.

## Licensing

Geocube-Client-Go is licensed under the Apache License, Version 2.0. See [LICENSE](https://github.com/airbusgeo/geocube-client-go/LICENSE) for the full license text.

## Credits

Geocube is a project under development by [Airbus DS Geo SA](http://www.intelligence-airbusds.com) with the support of [CNES](http://www.cnes.fr).
