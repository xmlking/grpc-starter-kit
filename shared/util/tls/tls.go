package tls

import (
    "crypto/tls"
    "crypto/x509"

    "github.com/xmlking/grpc-starter-kit/shared/util/ioutil"
)

func GetTLSConfig(certFile string, keyFile string, caFile string, address string) (tlsConfig *tls.Config, err error) {
    var cert tls.Certificate
    var certPEMBlock, keyPEMBlock, caPEMBlock []byte
    // cert, err = tls.LoadX509KeyPair(certFile, keyFile)
    certPEMBlock, err = ioutil.ReadFile(certFile)
    if err != nil {
        return
    }
    keyPEMBlock, err = ioutil.ReadFile(keyFile)
    if err != nil {
        return
    }
    cert, err = tls.X509KeyPair(certPEMBlock, keyPEMBlock)
    if err != nil {
        return
    }
    caPEMBlock, err = ioutil.ReadFile(caFile)
    if err != nil {
        return
    }
    caCertPool := x509.NewCertPool()
    caCertPool.AppendCertsFromPEM(caPEMBlock)

    tlsConfig = &tls.Config{
        Certificates: []tls.Certificate{cert},
        ServerName:   address,
        RootCAs:      caCertPool,
    }
    return
}
