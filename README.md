uafaker is a Go module that generates fake user-agent string

## Installation

```
$ go get github.com/subsan/uafaker
```

## Usage

``` go
package main

import (
	"log"

	"github.com/subsan/uafaker"
)

func main() {
    // random from all browsers
    userAgent := uafaker.Random()
    log.Printf("Random browser user-agent: %s", userAgent)

    // random only mobile browsers
    mobileUserAgent := uafaker.Mobile().Random()
    log.Printf("Mobile browser user-agent: %s", mobileUserAgent)

    // random only mobile chrome browsers
    mobileUserAgent := uafaker.Mobile().Chrome().Random()
    log.Printf("Mobile browser user-agent: %s", mobileUserAgent)
}
```

## Example
``` go
httpClient := &http.Client{}

req, _ := http.NewRequest("GET", "https://google.com", nil)

req.Header.Set("User-Agent", uafaker.Desktop().Random())

res, _ := httpClient.Do(req)
```

### Allowed filters
- Desktop()
- Mobile()
- OS()
- Windows()
- MacOS()
- Linux()
- IOS()
- Android()
- Chrome()
- Firefox()
- Safari()
- IE()
- Edge()
- Opera()
- Yandex()
