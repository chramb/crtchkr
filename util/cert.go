package util

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
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
		// TODO: not prettiest but less typing: handle error without overwriting err
		cert, err := getCertFromFile(link)
		if err != nil {
			return nil, err
		}
		certs = append(certs, cert)
	} else if u.Scheme == "https" {
		certs, err = getCertFromHTTPS(u)
		if err != nil {
			return nil, err
		}
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

func CheckCert(cert *x509.Certificate) ([][]*x509.Certificate, error) {
	/* TODO: expose in flags
	x509.VerifyOptions{
		DNSName:                   "",
		Intermediates:             nil,
		Roots:                     nil,
		CurrentTime:               time.Time{},
		KeyUsages:                 nil,
		MaxConstraintComparisons: 0,
	}*/
	cerPool, _ := x509.SystemCertPool()
	chains, err := cert.Verify(x509.VerifyOptions{
		Roots: cerPool,
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
