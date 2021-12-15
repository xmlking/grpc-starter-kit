# certificates

identities, certificates and keys

There are 3 identities:

- **CA**: Certificate Authority for **Client** and **Proxy**. It has the
  self-signed certificate *ca-cert.pem*. *ca-key.pem* is its private key.
- **Client**: It has the certificate *client-cert.pem*, which is signed by the
  **CA** using the config *client-cert.cfg*. *client-key.pem* is its private key.
- **Proxy**: It has the certificate *proxy-cert.pem*, which is signed by the
  **CA** using the config *proxy-cert.cfg*. *proxy-key.pem* is its private key.
- **Upstream CA**: Certificate Authority for **Upstream**. It has the self-signed
  certificate *upstream-ca-cert.pem*. *upstream-ca-key.pem* is its private key.
- **Upstream**: It has the certificate *upstream-cert.pem*, which is signed by
  the **Upstream CA** using the config *upstream-cert.cfg*. *upstream-key.pem* is
  its private key.
- **Upstream localhost**: It has the certificate *upstream-localhost-cert.pem*, which is signed by
  the **Upstream CA** using the config *upstream-localhost-cert.cfg*. *upstream-localhost-key.pem* is
  its private key. The different between this certificate and **Upstream** is that this certificate
  has a SAN for "localhost".

## How to update certificates

**certs.sh** has the commands to generate all files. Running `certs.sh` directly
will cause all files to be regenerated. So if you want to regenerate a
particular file, please copy the corresponding commands from `certs.sh` and
execute them in command line.

```bash
# at project root, run:
./config/certs/certs.sh
```

## mTLS

For Kubernetes environment, use [AutoCert](https://github.com/smallstep/autocert) or [ALTS](https://github.com/salrashid123/grpc_alts) based on configuration 

1. Setup  **AutoCert** operator in Kubernetes environment.
2. Add  annotations e.g., `autocert.step.sm/name: hello-mtls.default.svc.cluster.local` to gRPC server/client deployment YAML. 
3. Configure gRPC Server/Client use certs mounted at `/var/run/autocert.step.sm/` automatically by **AutoCert**

## Implementation

use `GetCertificate` and `GetClientCertificate` to reload renewed certs when certs are rotated dynamically by [AutoCert](https://github.com/smallstep/autocert)

Ref: [autocert](https://github.com/smallstep/autocert/tree/master/examples/hello-mtls/go-grpc)

```go
	tlsConfig := &tls.Config{
		RootCAs:          roots,
		MinVersion:       tls.VersionTLS12,
		CurvePreferences: []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		},
		// GetClientCertificate is called when a server requests a
		// certificate from a client.
		//
		// In this example keep alives will cause the certificate to
		// only be called once, but if we disable them,
		// GetClientCertificate will be called on every request.
		GetClientCertificate: r.getClientCertificate,
	}
```
```go
	tlsConfig := &tls.Config{
		ClientAuth:               tls.RequireAndVerifyClientCert,
		ClientCAs:                roots,
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		},
		GetCertificate: r.getCertificate,
	}
```

Or use **ACME** go client: https://github.com/smallstep/go-grpc-example/tree/main/server-acme

Example: [Go gRPC TLS — Practical Zero Trust](https://smallstep.com/practical-zero-trust/go-grpc-tls)

## Reference

- [RSA vs DSA vs ECDSA](https://www.misterpki.com/rsa-dsa-ecdsa/)
- [Certs from step-ca via ACME protocol](https://smallstep.com/practical-zero-trust/go-grpc-tls)
- [Securing gRPC connections with TLS](https://itnext.io/practical-guide-to-securing-grpc-connections-with-go-and-tls-part-2-994ef93b8ea9) via [certify](https://github.com/johanbrandhorst/certify)
- [Go gRPC TLS — Practical Zero Trust](https://smallstep.com/practical-zero-trust/go-grpc-tls)
