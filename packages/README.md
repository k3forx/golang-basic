## Initialization

```bash
go mod init github.com/k3forx/<package name>
```

## Add a custom (local) packages

```golang
module github.com/k3forx/golang-basic

go 1.16

replace local.packages/helpers => ./helpers # Need to specify the location of local package as relative path
```

## Use the custom package

```golang
package main

import (
        "log"

        "github.com/k3forx/golang-basic/helpers"
)

func main() {
        log.Println("Hello World")

        var myVar helpers.SomeType
        myVar.TypeName = "string"
        myVar.TypeNumber = 123
        log.Println(myVar)
}
```
