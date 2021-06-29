package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	var (
		conn *tls.Conn
		err  error
	)

	tlsConfig := http.DefaultTransport.(*http.Transport).TLSClientConfig

	c := &http.Client{
		Transport: &http.Transport{
			DialTLS: func(network, addr string) (net.Conn, error) {
				conn, err = tls.Dial(network, addr, tlsConfig)
				return conn, err
			},
		},
	}

	res, err := c.Get(os.Getenv("ENDPOINT")) //"https://api.sandbox.ebay.com/ws/api.dll"
	if err != nil {
		log.Fatal(err)
	}

	versions := map[uint16]string{
		tls.VersionSSL30: "SSL",
		tls.VersionTLS10: "TLS 1.0",
		tls.VersionTLS11: "TLS 1.1",
		tls.VersionTLS12: "TLS 1.2",
	}

	fmt.Println("Endpoint:", res.Request.URL)
	fmt.Println("HTTP Status:", res.Status)
	v := conn.ConnectionState().Version
	fmt.Println("CipherSuite:", conn.ConnectionState().CipherSuite)
	fmt.Println("TLS Version:", versions[v])
	fmt.Println("ConnectionState Version:", v)

	res.Body.Close()
}
