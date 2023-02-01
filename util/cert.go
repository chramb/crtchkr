package util

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
)

func GetCerts(link string) ([]*x509.Certificate, error) {
	// TODO: map of protocols
	u, err := url.Parse(link)
	if err != nil {
		return nil, err
	}
	var certs []*x509.Certificate
	if u.Host == "" {
		certificates, err := getCertsFromFile(link)
		if err != nil {
			return nil, err
		}
		certs = append(certificates)
	} else if u.Scheme == "https" {
		certificates, err := getCertFromHTTPS(u)
		if err != nil {
			return nil, err
		}
		certs = append(certificates)
	}
	// DEBUG: fmt.Printf("--- Host: %s ---\nScheme: %s\nHostname: %s\nPort: %s\n", u.Host, u.Scheme, u.Hostname(), u.Port())
	return certs, nil
}

func getCertFromHTTPS(u *url.URL) ([]*x509.Certificate, error) {
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
	return certs, nil
}

func CheckCert(cert *x509.Certificate, roots *x509.CertPool) ([][]*x509.Certificate, error) {
	/* TODO: expose in flags
	x509.VerifyOptions{
		DNSName:                   "",
		Intermediates:             nil,
		Roots:                     nil,
		CurrentTime:               time.Time{},
		KeyUsages:                 nil,
		MaxConstraintComparisons: 0,
	}*/
	chains, err := cert.Verify(x509.VerifyOptions{
		Roots: roots,
	})
	if err != nil {
		return nil, err
	}
	return chains, nil
}

func getSystemCertPool() *x509.CertPool {
	CertPool, err := x509.SystemCertPool()
	if err != nil {
		CertPool = x509.NewCertPool()
	}
	return CertPool
}

func getCertsFromFile(filePath string) ([]*x509.Certificate, error) {
	// Read the PEM-encoded certificate file
	certPEM, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Failed to read certificate file:", err)
		return nil, err
	}

	// Decode the PEM-encoded certificates
	var certificates []*x509.Certificate
	for len(certPEM) > 0 {
		var block *pem.Block
		block, certPEM = pem.Decode(certPEM)
		if block == nil {
			break
		}

		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			fmt.Println("Failed to parse certificate:", err)
			return nil, err
		}

		certificates = append(certificates, cert)
	}

	// Print the Common Name of each certificate
	return certificates, nil
}
