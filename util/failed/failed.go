package failed

import "crypto/x509"

type FailedCert struct {
	Cert *x509.Certificate
	Err  error
}
type FailedCerts []FailedCert

var Fc = FailedCerts{}
