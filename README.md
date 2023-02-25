# *fnboot*

![Go version](https://img.shields.io/badge/go-%3E%3Dv1.18-9cf)
[![Release](https://img.shields.io/badge/release-0.14.1-green.svg)](https://github.com/fraynjo/fnboot/releases)
[![GoDoc](https://godoc.org/github.com/fraynjo/fnboot?status.svg)](https://pkg.go.dev/github.com/fraynjo/fnboot)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/fraynjo/fnboot/blob/main/LICENSE)

## Installation

```
go get -u github.com/fraynjo/fnboot
```

## Usage

```
import "github.com/fraynjo/fnboot/string"
```

## Example

```
package main

import (
    "fmt"
    "github.com/fraynjo/fnboot/string"
)

func main() {
	lowerStr := string.ToLower("FNBOOT")
	fmt.Println(lowerStr)
}
```
