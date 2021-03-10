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
    mobileChromeUserAgent := uafaker.Mobile().Chrome().Random()
    log.Printf("Mobile chrome browser user-agent: %s", mobileChromeUserAgent)

    // random only chrome or firefox mobile browsers
    mobileChromeOrFirefoxUserAgent := uafaker.Chrome().Firefox().Mobile().Random()
    log.Printf("Chrome or firefox mobile browser user-agent: %s", mobileChromeOrFirefoxUserAgent)

    // generate 20 desktop, windows or linux, chrome or firefox browsers
    faker := uafaker.Desktop().Windows().Linux().Chrome().Firefox()
    for i := 0; i < 20; i++ {
        fmt.Println(faker.Random())
    }
}
```

## Example
``` go
httpClient := &http.Client{}

req, _ := http.NewRequest("GET", "https://google.com", nil)

req.Header.Set("User-Agent", uafaker.Desktop().Random())

res, _ := httpClient.Do(req)
```

### Available filter groups
#### Device:
- Desktop()
- Mobile()
#### OS:
- Windows()
- MacOS()
- Linux()
- IOS()
- Android()
#### Browser:
- Chrome()
- Firefox()
- Safari()
- IE()
- Edge()
- Opera()
- Yandex()
- OS()