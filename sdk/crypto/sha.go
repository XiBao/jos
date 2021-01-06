package crypto

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

const (
	PublicKey = `-----BEGIN CERTIFICATE-----
MIIEcDCCA1igAwIBAgIJAKCBMSvIHNiEMA0GCSqGSIb3DQEBBQUAMIGAMQswCQYD
VQQGEwJDTjEQMA4GA1UECBMHQmVpamluZzEQMA4GA1UEBxMHQmVpamluZzEPMA0G
A1UEChMGSkQuQ09NMQwwCgYDVQQLEwNKT1MxEzARBgNVBAMTCmpvcy5qZC5jb20x
GTAXBgkqhkiG9w0BCQEWCmpvc0BqZC5jb20wIBcNMTkwMzE1MDQ1NTM2WhgPMjA1
OTAzMDUwNDU1MzZaMIGAMQswCQYDVQQGEwJDTjEQMA4GA1UECBMHQmVpamluZzEQ
MA4GA1UEBxMHQmVpamluZzEPMA0GA1UEChMGSkQuQ09NMQwwCgYDVQQLEwNKT1Mx
EzARBgNVBAMTCmpvcy5qZC5jb20xGTAXBgkqhkiG9w0BCQEWCmpvc0BqZC5jb20w
ggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDf9jdFbaYJLb6H/B1EEtuO
okkjrU1taQSudZhuBlnzCiKeUjK6vYDoqGgJSzRI86slU/rkK/7o4mc8LOvmAJRv
ULWLUdM9EzI+6+M6eVLwuWnm3QMIJJl1y7dQqwnAMLl3T/P6UGP1g19R7D8LcaEw
289Y8i/qJaVdobaM822xcW4Wv+QIldlWo6YlDoE7dfY9pXTlAkTP/GzO+LOnCzp1
/VA3Q6Xl1Cl4Kvk0wFWnGiMEbVEZx9yEknwPV1Viq3QGjMPoEGEau6x9srCcEitC
lllqXHOWkIVNt//qN2ubx90wjyHKZTe3HrQ/LFSIWLTeNo738iR8tFzxSfa5hitZ
AgMBAAGjgegwgeUwHQYDVR0OBBYEFHYHDa2moq7nEccftSm3x72QBWWJMIG1BgNV
HSMEga0wgaqAFHYHDa2moq7nEccftSm3x72QBWWJoYGGpIGDMIGAMQswCQYDVQQG
EwJDTjEQMA4GA1UECBMHQmVpamluZzEQMA4GA1UEBxMHQmVpamluZzEPMA0GA1UE
ChMGSkQuQ09NMQwwCgYDVQQLEwNKT1MxEzARBgNVBAMTCmpvcy5qZC5jb20xGTAX
BgkqhkiG9w0BCQEWCmpvc0BqZC5jb22CCQCggTEryBzYhDAMBgNVHRMEBTADAQH/
MA0GCSqGSIb3DQEBBQUAA4IBAQAr9qLL6qkNJjtcOzYM5afdyt+KBF9iwIcKG8ca
NUPNXwOFnOFw/JBKR4svjafvV3rSGs7ZtVMmASLUhrtStwfJJvXV7tdyqC0p44u/
sWK6SHoTNIHX+kXbzKrkwggqeTiUlHDTw60BP/mmbrYhIwOiTNvI247iWZ4IxxyD
bpFULv0gBfTVuc/ATWrHTI2pT78lIectDgUCpTOAhQIvE0PLK9nZjrsSCvW7tRED
PC+6KCPYQAzxmKvRRMCHXkAVeqb/0M6GEXBIT0aYEBHKdQ7s4g1VSGrbMUL5mQsA
+3fYhR+QEhE8PboH5kVct1V9tiMpx7kymJQKVfNufC3FIlyr
-----END CERTIFICATE-----`
)

func Sha256(origData []byte) []byte {
	h := sha256.New()
	h.Write(origData)
	return h.Sum(nil)
}

func RSAVerifySignWithSha256(publicKey []byte, origData []byte, signData string) error {
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return errors.New("public key error")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Println(err)
		return err
	}
	pub := cert.PublicKey.(*rsa.PublicKey)
	h := sha256.New()
	h.Write(origData)
	hash := h.Sum(nil)
	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hash, sign)
}
