# smallstep CA

[Zero Trust](https://smallstep.com/docs/practical-zero-trust) or [BeyondProd](https://cloud.google.com/security/beyondprod) approaches require authenticated and encrypted communications everywhere.

## Usage 

Using my hosted CA: https://smallstep.com/app/chinthagunta/cm/authorities

### Install

```bash
brew install step
```

### Configure `step` to use this authority

```bash
step ca bootstrap --ca-url https://sumanth.chinthagunta.ca.smallstep.com --fingerprint 8604fd19cade85cb10021b0912f4ac666aefaee68dc136cc8846fad254150afe
```

### Issue a certificate

Request the Common Name (eg. myservice) and SANs (eg. myservice.internal.mycompany.net) you'd like to include in your certificate. <br/>
The last two parameters specify the resulting certificate and private key filenames.

```bash
step ca certificate myservice --san "myservice.default.svc.cluster.local" myservice.crt myservice.key --not-after 24h
```

Inspect your new certificate to see all the details.

```bash
step certificate inspect --short myservice.crt
```


Request a copy of your CA root certificate, which will be used to make sure each application can trust certificates presented by other applications.

```bash
step ca root ca.crt
```
