<a href="https://github.com/tommy-sho/grouper/releases">
    <img
        src="https://img.shields.io/github/v/release/tommy-sho/grouper.svg"
        alt="GitHub Releases"/>
</a>
<br />
<a href="https://img.shields.io/github/license/tommy-sho/grouper">
    <img
        src="https://img.shields.io/github/license/tommy-sho/grouper"
        alt="license"/>
</a>
<br />
<a href="https://github.com/tommy-sho/grouper/actions">
    <img src="https://github.com/tommy-sho/grouper/workflows/Test/badge.svg" alt="build status" />
</a>

## Features
 - grouping/ordering import, if blank line exists before/after paths.

## Usage

- set directories or files as arguments of the command.
- if you set the prefix to -local option, the packages which beginning with that prefix put after 3rd-party packages.
- use -w option to override original files.
```shell
GLOBAL OPTIONS:
   --local value, -l value  specify imports prefix beginning with this string after 3rd-party packages. especially your own organization name. comma-separated list
   --write, -w              write result source to original file instead od stdout (default: false)
```


## Example

example command.

```shell script
$ grouper -local "github.com/tommy-sho" path_to_file.go
```

- before
```go
package main

import(
	"bytes"
	"errors"
	"github.com/tommy-sho/grouper"

	"golang.org/x/tools/go/ast/astutil"
	
	"golang.org/x/tools/imports"
)
...
```


- after grouper run...
```go
package main

import (
        "bytes"
        "errors"
        "fmt"

        "golang.org/x/tools/go/ast/astutil"
        "golang.org/x/tools/imports"

        "github.com/tommy-sho/grouper"
)
```

- in case of goimports...
```go
package main

import(
        "bytes"
        "errors"

        "github.com/tommy-sho/grouper"

        "golang.org/x/tools/go/ast/astutil"

        "golang.org/x/tools/imports"
)
```


## Install

with `go` command.

```shell script
$  go install github.com/tommy-sho/grouper
```

## License
MIT

## Author
tommy-sho