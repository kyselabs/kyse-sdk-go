# kyse-sdk (go)

The SDK library for Kyse.

### Install

```bash
go get github.com/kyselabs/kyse-sdk-go
```

### Authentication Types

```go
	var authentication auth.AuthMethod

	// Based on User Credentials
	authentication = auth.UserBased{
		Username: "admin",
		Password: "admin",
	}

	// Based on Kyse Keys
	authentication = auth.Key{
		Key: "YOUR_KEY",
	}

	// Based on OAuth2 Device Auth Flow
	authentication = auth.DeviceBased{}

	kyse := sdk.NewKyse(authentication, nil)
```

### Quick Start

```go
package main

import (
	"fmt"

	sdk "github.com/kyselabs/kyse-sdk-go"
	"github.com/kyselabs/kyse-sdk-go/auth"
	"github.com/kyselabs/kyse-sdk-go/sca"
)

func main() {
	// Set the KYSE-API and KYSE-AUTH Addresses

	// KYSE_API_ADDRESS=http://localhost:1337
	// AUTH_API_ADDRESS=http://localhost:8443

	// Create a Kyse Instance
	kyse := sdk.NewKyse(&auth.KeyBased{Key: ""}, nil)

	client := sca.NewSCA(kyse)

	resources := client.Audit([]sca.Resource{
		{
			Vendor: "PyPI",
			Assets: []sca.Asset{
				{
					Package: "requests",
					Version: "1.2.3",
				},
			},
		},
	})

	for _, resource := range resources {
		fmt.Printf("%+v\n", resource)
	}
}
```
