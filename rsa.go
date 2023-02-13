package client
import (
  "bytes"
  "crypto/rand"
  "crypto/rsa"
  "crypto/x509"
  "crypto/x509/pkix"
  "encoding/ans1"
  "encoding/pem"
  "fmt"
  "math/big"
  "time"
)
// from https://gist.github.com/gambol99/d55afd69217b8e2dd727be99f0a20e7d
// createCertificateAuthority generates a certificate authority request ready to be signed
//
func (r *secretStore) createCertificateAuthority(username string, expireDays uintt64) (*caCertificate, error) {
	expiration := time.Duration(expireDays) * 24 * time.Hour
  names := pkix.Name{
    Country:      []string{"CN"},
    Organization: []string{"SeePlant"},
    Locality:     []string{"BJ"},
    CommonName:   username,
  }
  // step: generate a keypair
	keys, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		return nil, fmt.Errorf("unable to genarate private keys, error: %s", err)
	}

	// step: generate a csr template
	var csrTemplate = x509.CertificateRequest{
		Subject:            names,
		SignatureAlgorithm: x509.SHA512WithRSA,
		ExtraExtensions: []pkix.Extension{
			{
				Id:       asn1.ObjectIdentifier{2, 5, 29, 19},
				Value:    []byte("ca:true"),
				Critical: true,
			},
		},
	}
	// step: generate the csr request
	csrCertificate, err := x509.CreateCertificateRequest(rand.Reader, &csrTemplate, keys)
	if err != nil {
		return nil, err
	}
	csr := pem.EncodeToMemory(&pem.Block{
		Type: "CERTIFICATE REQUEST", Bytes: csrCertificate,
	})

	// step: generate a serial number
	serial, err := rand.Int(rand.Reader, (&big.Int{}).Exp(big.NewInt(2), big.NewInt(159), nil))
	if err != nil {
		return nil, err
	}

	now := time.Now()
	// step: create the request template
	template := x509.Certificate{
		SerialNumber:          serial,
		Subject:               names,
		NotBefore:             now.Add(-10 * time.Minute).UTC(),
		NotAfter:              now.Add(expiration).UTC(),
		BasicConstraintsValid: true,
		IsCA:        true,
		KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
	}

	// step: sign the certificate authority
	certificate, err := x509.CreateCertificate(rand.Reader, &template, &template, &keys.PublicKey, keys)
	if err != nil {
		return nil, fmt.Errorf("failed to generate certificate, error: %s", err)
	}

	var request bytes.Buffer
	var privateKey bytes.Buffer
	if err := pem.Encode(&request, &pem.Block{Type: "CERTIFICATE", Bytes: certificate}); err != nil {
		return nil, err
	}
	if err := pem.Encode(&privateKey, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(keys)}); err != nil {
		return nil, err
	}

	return &caCertificate{
		privateKey: privateKey.String(),
		publicKey:  request.String(),
		csr:        string(csr),
	}, nil
}

type caCertificate struct {
  privateKey string
  publicKey  string
  csr        string
}
