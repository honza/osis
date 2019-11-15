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

Supported values
----------------

* `Ps` (whole book)
* `Rom.8` (chapter)
* `John.3.16` (single verse)
* `Jude.1.1` (single verse in a single chapter book)
* `Gen.1-Gen.2` (multiple chapters in one book)
* `Gen.1.1-Gen.1.2` (verse range in the same book)
* `Ps.1.1` (psalms are special)
* `Ps.1-Ps.2` (pluralize if multiple)
* `John.7.53-John.8.11` (range over multiple chapters ina single book)
* `Luke.12-Acts.1` (chapter range across books)
* `Luke.12.1-Acts.1.1` (verse range over multiple books)

License
-------

Apache 2.0
