package util

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"net/url"
)

func GetCert(link string) (*x509.Certificate, error) {
	// TODO: map of protocols
	u, err := url.Parse(link)
	if err != nil {
		return nil, err
	}
	var cert *x509.Certificate
	if u.Host == "" {
		// TODO: not prettiest but less typing: handle error without overwriting err
		cert, err = getCertFromFile(link)
	} else if u.Scheme == "https" {
		cert, err = getCertFromHTTPS(u)
	}
	if err != nil {
		return nil, err
	}
	// DEBUG: fmt.Printf("--- Host: %s ---\nScheme: %s\nHostname: %s\nPort: %s\n", u.Host, u.Scheme, u.Hostname(), u.Port())
	return cert, nil

}

func getCertFromHTTPS(u *url.URL) (*x509.Certificate, error) {
	host := u.Host
	var port string
	if u.Port() == "" {
		port = "443"
	} else {
		port = u.Port()
	}
	// DEBUG: fmt.Printf("Dialing: %s:%s\n", host, port)
	conn, err := tls.Dial("tcp", host+":"+port, &tls.Config{})
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	state := conn.ConnectionState()
	certs := state.PeerCertificates
	if len(certs) == 0 {
		return nil, errors.New("no certificates found")
	}

	// TODO: handle more than one certificate
	return certs[0], nil
}

func getCertFromFile(filePath string) (*x509.Certificate, error) {
	certPEM, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(certPEM)
	if block == nil {
		return nil, errors.New("failed to parse certificate PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}
	return cert, nil
}
