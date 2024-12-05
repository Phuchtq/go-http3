package cmd

import (
	"crypto/tls"
	"crypto/x509"
)

var protos = []string{
	"http/1.1", // HTTP/1.1
	"h2",       // HTTP/2 - gRPC
	"h3",       // HTTP3
}

func generateTlsConfig() *tls.Config {
	return &tls.Config{
		Certificates: []tls.Certificate{generateTLSCert()},
		NextProtos:   protos,
	}
}

func generateTLSCert() tls.Certificate {
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		panic("")
	}

	// Parse the certificate
	certData, err := x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		panic("Error while parsing TLS certificate: " + err.Error())
	}

	// Certificate expired
	if isCertificateExpired(certData) {
		generateCertificate() // Regenerate

		res, err := tls.LoadX509KeyPair(certPath, keyPath) // Reload
		if err != nil {
			panic("")
		}

		return res
	}

	return cert
}
