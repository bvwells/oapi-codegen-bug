# Bug in deepmap/oapi-codegen

This repo illustrates a bug in github.com/deepmap/oapi-codegen. The bug appears
when a colon appears in the path in an API.

The example located in /petstore/petstore.yaml contains an API with the path
'/pets:validate:'. The way in which the request URL is constructed from the 
server URL and the path of the API leads to a bug in which the path is completely
wrong. The generate code in question is:

```Go
func NewValidatePetsRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/pets:validate")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}
```

Pulling this code out into a small example:

```Go
package main

import (
	"fmt"
	"net/url"
)

func main() {
	server := "http://petstore.swagger.io/api"
	queryUrl, err := url.Parse(server)
	if err != nil {
		panic(err)
	}

	basePath := fmt.Sprintf("/pets:validate")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("host=%s\n", queryUrl.Host)
	fmt.Printf("scheme=%s\n", queryUrl.Scheme)
	fmt.Printf("path=%s\n", queryUrl.Path)
}
```

Running this program results in 

```bash
host=
scheme=pets
path=
```

See https://play.golang.org/p/JBkhGVWIFzV
