package main

import (
	"crypto/tls"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Enter domain (e.g., google.com:443):")
	var domain string
	fmt.Scanln(&domain)

	conn, err := tls.Dial("tcp", domain, nil)
	if err != nil {
		fmt.Printf("Error connecting: %v\n", err)
		return
	}
	defer conn.Close()

	cert := conn.ConnectionState().PeerCertificates[0]
	fmt.Printf("Issuer: %s\n", cert.Issuer)
	fmt.Printf("Expiration: %s\n", cert.NotAfter.Format(time.RFC1123))
	fmt.Printf("Valid From: %s\n", cert.NotBefore.Format(time.RFC1123))

	if time.Now().After(cert.NotAfter) {
		fmt.Println("Warning: Certificate has expired!")
	} else if time.Until(cert.NotAfter) < 30*24*time.Hour {
		fmt.Println("Warning: Certificate is expiring soon!")
	} else {
		fmt.Println("Certificate is valid.")
	}
}
