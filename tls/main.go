package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net/http"
	"time"
)

func generateTLSCert() (*tls.Certificate, error) {
	k, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	t := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: "localhost"},
		NotBefore:    time.Now().Add(-30 * time.Second),
		NotAfter:     time.Now().Add(86400 * time.Second),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageServerAuth,
		},
	}
	t.BasicConstraintsValid = true
	c, err := x509.CreateCertificate(rand.Reader, &t, &t, &k.PublicKey, k)
	if err != nil {
		return nil, err
	}
	certPem := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: c,
	})
	keyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(k),
	})
	tlsCert, err := tls.X509KeyPair(certPem, keyPem)
	if err != nil {
		return nil, err
	}
	return &tlsCert, err
}

func main() {
	tlsCert, err := generateTLSCert()
	if err != nil {
		panic(err)
	}
	http.Handle("/", http.FileServer(http.Dir(".")))
	s := http.Server{
		Addr: "localhost:8443",
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{*tlsCert},
		},
	}
	err = s.ListenAndServeTLS("", "")
	if err != nil {
		panic(err)
	}
}
