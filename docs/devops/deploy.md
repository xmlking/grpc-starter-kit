# Deploy

## Setup

follow this __GCP__ setup [guide](https://spring-gcp.saturnism.me/getting-started/gcloud-cli) to install `gcloud` CLI and `BigTable`, `PubSub` emulators

## Deploying to Cloud Run

Cloud Run needs to pull our application from a container registry, so the first
step is to push the image we just built.

Make sure that [you can use `gcloud`](https://cloud.google.com/sdk/gcloud/reference/auth/login)
and are [able to push to `gcr.io`.](https://cloud.google.com/container-registry/docs/pushing-and-pulling)

```bash
gcloud auth login
gcloud auth configure-docker
```

Now we can push our image.

```bash
docker push gcr.io/$GCP_PROJECT/account-service:latest
```

Finally, we deploy our application to Cloud Run:

```bash
GCP_REGION="us-west1" # Or ...
gcloud run deploy account-service \
--image=gcr.io/$GCP_PROJECT/account-service:latest \
--platform=managed \
--allow-unauthenticated \
--project=${GCP_PROJECT} \
--region=${GCP_REGION}
```

This command will give you a message like
```
Service [account-service] revision [account-service-00001-baw] has been deployed and is serving 100 percent of traffic at https://account-service-xyspwhk3xq-uc.a.run.app
```

We can programmatically determine the gRPC service's endpoint:

```bash
ENDPOINT=$(\
  gcloud run services list \
  --project=${GCP_PROJECT} \
  --region=${GCP_REGION} \
  --platform=managed \
  --format="value(status.address.url)" \
  --filter="metadata.name=account-service") 
ENDPOINT=${ENDPOINT#https://} && echo ${ENDPOINT}
```

Notice that this endpoint is secured with TLS even though the server we wrote 
uses a plaintext connection. Cloud Run provides a proxy that provides TLS for us.

We'll account for that in our `grpcurl` invocation by omitting the `-plaintext` flag:

```bash
grpcurl \
    -proto proto/demoapi/gkit/service/greeter/v1/greeter.proto \
    -d '{"firstName": "sumo", "lastName": "demo", "email": "sumo@demo.com"}' \
    ${ENDPOINT}:443 \
    gkit.service.account.v1.AccountService
```

There's a simple Golang client too:

```bash
go run github.com/xmlking/grpc-starter-kit/cmd/account/main.go \
--gprc_endpoint=${ENDPOINT}:443
```

You have an auto-scaling gRPC-based account service!

## Reference
- [gRPC in Google Cloud Run](https://github.com/grpc-ecosystem/grpc-cloud-run-example)
