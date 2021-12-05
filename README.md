# Golang universal captcha resolve service

[![GoDoc](https://godoc.org/github.com/xta/okrun?status.svg)](https://pkg.go.dev/github.com/Flagon00/CaptchaSolvingServiceClient)

One package for all services that support /createTask and /getTaskResult api endpoints. Tested on cptch.net, anti-captcha.com, 2captcha.com and XEvil

Setup:
```go get -u github.com/Flagon00/CaptchaSolvingServiceClient```

Example usage with reCaptchaV2 and cptch.net:
```go
package main

import (
	"log"

	"github.com/Flagon00/CaptchaSolvingServiceClient"
)

func main() {
	client, err := captcha.Client(true, "cptch.net", "api-key")
	if err != nil{
		log.Fatal(err)
	}

	resolve, err := client.ReCaptchaV2("https://www.google.com/recaptcha/api2/demo",  "6Le-wvkSAAAAAPBMRTvw0Q4Muexq9bi0DJwx_mJ-", "", false, 60)
	if err != nil{
		log.Fatal(err)
	}

	log.Println(resolve)
}
```

Also usage with image captcha example:
```go
client, err := captcha.Client(true, "cptch.net", "api-key")
if err != nil{
	log.Fatal(err)
}

resolve, err := client.RegularCaptcha("base64-string", 60)
if err != nil{
	log.Fatal(err)
}

log.Println(resolve)
```

Example client for 2captcha:
```go
captcha.Client(true, "2captcha.com", "api-key")
```

Or if you want, you can use this package with XEvil:
```go
captcha.Client(false, "localhost", "api-key")
```