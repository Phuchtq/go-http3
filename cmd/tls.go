package cmd

import "crypto/tls"

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
