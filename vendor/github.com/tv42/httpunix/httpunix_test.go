package httpunix_test

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/tv42/httpunix"
)

func Example_clientStandalone() {
	// This example shows using a customized http.Client.
	u := &httpunix.Transport{
		DialTimeout:           100 * time.Millisecond,
		RequestTimeout:        1 * time.Second,
		ResponseHeaderTimeout: 1 * time.Second,
	}
	u.RegisterLocation("myservice", "/path/to/socket")

	var client = http.Client{
		Transport: u,
	}

	resp, err := client.Get("http+unix://myservice/urlpath/as/seen/by/server")
	if err != nil {
		log.Fatal(err)
	}
	buf, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", buf)
	resp.Body.Close()
}

func Example_clientIntegrated() {
	// This example shows handling all net/http requests for the
	// http+unix URL scheme.
	u := &httpunix.Transport{
		DialTimeout:           100 * time.Millisecond,
		RequestTimeout:        1 * time.Second,
		ResponseHeaderTimeout: 1 * time.Second,
	}
	u.RegisterLocation("myservice", "/path/to/socket")

	// If you want to use http: with the same client:
	t := &http.Transport{}
	t.RegisterProtocol(httpunix.Scheme, u)
	var client = http.Client{
		Transport: t,
	}

	resp, err := client.Get("http+unix://myservice/urlpath/as/seen/by/server")
	if err != nil {
		log.Fatal(err)
	}
	buf, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", buf)
	resp.Body.Close()
}

func Example_server() {
	l, err := net.Listen("unix", "/path/to/socket")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	if err := http.Serve(l, nil); err != nil {
		log.Fatal(err)
	}
}
