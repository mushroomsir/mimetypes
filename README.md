# mimetypes
[![Build Status](https://img.shields.io/travis/mushroomsir/	.svg?style=flat-square)](https://travis-ci.org/mushroomsir/mimetypes)
[![Coverage Status](http://img.shields.io/coveralls/mushroomsir/mimetypes.svg?style=flat-square)](https://coveralls.io/github/mushroomsir/mimetypes?branch=master)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://github.com/mushroomsir/mimetypes/blob/master/LICENSE)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/mushroomsir/mimetypes)

## Installation

```sh
go get github.com/mushroomsir/mimetypes
```

## Usage
```go
package main

import "github.com/mushroomsir/mimetypes"
import "fmt"

func main() {

	fmt.Println(mimetypes.Lookup(".jpg"))
	// Output:  image/jpeg
}

```

## Licenses

All source code is licensed under the [MIT License](https://github.com/mushroomsir/mimetypes/blob/master/LICENSE).
