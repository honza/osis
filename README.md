osis
====

A Golang library for parsing OSIS values into English

Usage
-----

This library exposes two functions, `Format`, and `FormatMany`.

### Format

`Format` takes a string, and returns a formatted string and an error

### FormatMany

`FormatMany` takes a string, and returns a slice of formatted strings and an
error.  It expects a string like `"John.3.16,Ps.1.1"`.

```golang
package main

import (
	"fmt"
	"github.com/honza/osis"
)

func main() {
	english, err := osis.Format("John.3.16")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(english) // "John 3:16"
}
```

License
-------

Apache 2.0
