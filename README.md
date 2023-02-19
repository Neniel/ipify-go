# ipify-go

An unofficial and very simple Go library for using [ipify API](https://www.ipify.org/).

## Features
Get IP v4 and IP v6 in different formats such as:

### Text
```
127.0.0.1
```

### JSON
  ```json
    {"ip": "127.0.0.1"}
  ```

### JSONP
```jsonp
callback({"ip": "127.0.0.1"})
```

### JSONP with custom callback

```jsonp
my_function({"ip": "127.0.0.1"})
```


## How to get started

Please find below a quick example on how to get started with this library:

```go
package main

import (
	"fmt"
	"github.com/Neniel/ipify-go"
)

func main() {
	ip, err := ipify.GetIPv4()
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Println(ip)
}
```