package cmd

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"http3-integrate/constants"
	"log"
	"math/big"
	"os"
	"time"
)

// File path
const (
	certPath string = "Your certificate path.pem" // Extension file with .pem
	keyPath  string = "Your key path.pem"
)

// Cert type
const (
	certType       string = "CERTIFICATE"
	privateKeyType string = "EC PRIVATE KEY"
)

// Expiration
const (
	expiration time.Duration = time.Hour * 24 * 365 // 1 year
)

// func getCertificate() *x509.Certificate {
// 	var logger = &log.Logger{}

// 	// Read file
// 	data, err := os.ReadFile(certPath)
// 	if err != nil {
// 		logger.Println(fmt.Sprintf(constants.ReadFileErrMsg, certPath) + err.Error())
// 		return nil
// 	}

// 	// Decode data
// 	block, _ := pem.Decode(data)
// 	if block == nil || block.Type != certType {
// 		logger.Println("Decode PEM block containing certificate meets problem.")
// 		return nil
// 	}

// 	// Parse cert
// 	res, err := x509.ParseCertificate(block.Bytes)
// 	if err != nil {
// 		logger.Println("Parse certificate meets problem - " + err.Error())
// 		return nil
// 	}

// 	return res
// }

func generateCertificate() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)

	var logger = &log.Logger{}

	if err != nil {
		logger.Fatal(constants.GenerateCertificateErrMsg + err.Error())
	}

	var template = x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Country:      []string{"Your country"},
			Organization: []string{"Your organization"}, // Optional
			CommonName:   "*.your_domain",               // Multi-prefix mix
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(expiration),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	// Create certificate
	certBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		logger.Fatal(constants.GenerateCertificateErrMsg + err.Error())
	}

	// Create certificate file in the current directory
	certFile, err := os.Create(certPath)
	if err != nil {
		logger.Fatal(fmt.Sprintf(constants.CreateFileErrMsg, certPath) + err.Error())
	}
	defer certFile.Close()

	if err := pem.Encode(certFile, &pem.Block{
		Type:  certType,
		Bytes: certBytes,
	}); err != nil {
		logger.Fatal()
	}

	// Create private key file
	keyFile, err := os.Create(keyPath)
	if err != nil {
		logger.Fatal(fmt.Sprintf(constants.CreateFileErrMsg, keyPath) + err.Error())
	}
	defer keyFile.Close()

	// Convert key
	privateBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		logger.Fatal(constants.ConvertKeyErrMsg + err.Error())
	}

	// Encode key and write to file
	if err := pem.Encode(certFile, &pem.Block{
		Type:  certType,
		Bytes: privateBytes,
	}); err != nil {
		logger.Fatal()
	}
}

func isCertificateExpired(cert *x509.Certificate) bool {
	var current = time.Now()
	return current.Before(cert.NotBefore) || current.After(cert.NotAfter)
}
