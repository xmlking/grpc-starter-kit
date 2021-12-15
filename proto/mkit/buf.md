# mKit

**Microservices Kit (µKit)** contains _schema_ and _service_ definitions for this sample app

## Schema

### Entities

`mkit/schema/<domain>/v1` contains:
1. Command schema for **data models** that are passed via Google **PubSub** messages or schema of **BigQuery** table
2. Wrapper entities such as [CloudEvents](https://github.com/cloudevents/sdk-go/blob/main/binding/format/protobuf/v2/internal/pb/cloudevent.proto)  encryption wrappers etc.
3. schemas under `mkit/schema/<domain>/v1` are synchronized with [Schema Registry](https://cloud.google.com/pubsub/docs/schemas#go_3)

## Services

`mkit/service/<domain>/v1` contains:
1. gRPC service definitions for respective domain  

