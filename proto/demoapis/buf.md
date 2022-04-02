# grpc-starter-kit

**gRPC Services Kit** contains _schema_ and _service_ definitions for this demo app

## Schema

### Entities

`gkit/schema/<domain>/v1` contains:
1. Command schema for **data models** that are passed via Google **PubSub** messages or schema of **BigQuery** table
2. Wrapper entities such as [CloudEvents](https://github.com/cloudevents/sdk-go/blob/main/binding/format/protobuf/v2/internal/pb/cloudevent.proto)  encryption wrappers etc.
3. schemas under `gkit/schema/<domain>/v1` are synchronized with [Schema Registry](https://cloud.google.com/pubsub/docs/schemas#go_3)

## Services

`gkit/service/<domain>/v1` contains:
1. gRPC service definitions for respective domain  

