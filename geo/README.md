# ipify-go/geo

An unofficial and very simple Go library for using [ipify Geolocation API](https://geo.ipify.org/).

## Features
Get geolocation information based on data such as:

* IP address (IP v4 and IP v6)
* Domain
* Email

### Geolocation information
Based on how you use the API you can get:
* Country data
* City data
* VPN data

### Full request example response
This is an example extracted from the [official documentation](https://geo.ipify.org/docs):

```json
{
    "ip": "8.8.8.8",
    "location": {
        "country": "US",
        "region": "California",
        "city": "Mountain View",
        "lat": 37.40599,
        "lng": -122.078514,
        "postalCode": "94043",
        "timezone": "-07:00",
        "geonameId": 5375481
    },
    "domains": [
        "0d2.net",
        "003725.com",
        "0f6.b0094c.cn",
        "007515.com",
        "0guhi.jocose.cn"
    ],
    "as": {
        "asn": 15169,
        "name": "Google LLC",
        "route": "8.8.8.0/24",
        "domain": "https://about.google/intl/en/",
        "type": "Content"
    },
    "isp": "Google LLC",
    "proxy": {
        "proxy": false,
        "vpn": false,
        "tor": false
    },
}
```

## How to get started

Please find below a quick example on how to get started with this library:

```go
package main

import (
	"fmt"
	"github.com/Neniel/ipify-go/geo"
)

func main() {
	geoData, err := geo.GetGeo("YOUR_API_KEY", string[]{"country"}, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Println(geoData)
}
```