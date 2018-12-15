package roku

import (
	ssdp "github.com/koron/go-ssdp"
	"net/url"
)

// Find sends out a SSDP request on the network,
// looking for Roku devices that can be connected to.
func Find(seconds int) (endpoints Endpoints, err error) {
	services, err := ssdp.Search("roku:ecp", seconds, "")

	if err != nil {
		return nil, err
	}

	for _, srv := range services {
		url, err := url.Parse(srv.Location)

		if err != nil {
			return nil, err
		}

		endpoints = append(endpoints, &Endpoint{url.String()})
	}

	return
}
