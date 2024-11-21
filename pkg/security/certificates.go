// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
package security

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"time"
)

const (
	SEAGATE_CFM_SERVICE_CRT_FILEPATH = "/usr/local/share/ca-certificates/github_com_seagate_cfm-self-signed.crt"
)

// GenerateSelfSignedCert - Generates the self-signed SSL/TLS certificate and private key at runtime.
func GenerateSelfSignedCert() (*tls.Certificate, []byte, error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	notBefore := time.Now()
	notAfter := notBefore.Add(365 * 24 * time.Hour)

	serialNumber, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return nil, nil, err
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Country:            []string{"US"},
			Organization:       []string{"SEAGATE TECHNOLOGY LLC"},
			OrganizationalUnit: []string{"MAG"},
			Locality:           []string{"Longmont"},
			Province:           []string{"Colorado"},
			CommonName:         "localhost", // Set CommonName to a valid hostname
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IPAddresses:           []net.IP{net.ParseIP("127.0.0.1")}, // Add IP SAN
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return nil, nil, err
	}

	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	keyPEM, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		return nil, nil, err
	}

	keyPEMBytes := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: keyPEM})

	cert, err := tls.X509KeyPair(certPEM, keyPEMBytes)
	if err != nil {
		return nil, nil, err
	}

	return &cert, certPEM, nil
}
