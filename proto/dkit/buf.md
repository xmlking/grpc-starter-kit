# DataKit (dKit) 

**DataKit** module contains **tools** to safeguard **data** at rest and in transit.

custom protobuf options defined for this project<br/>
example: BugQuery, privacy/encryption

### Options

`dkit/<domain>/v1` contains:

1. bigquery
2. privacy

```protobuf
message MyMessage {
    string name = 1 [(dkit.privacy.v1.sensitive) = true];
}
```

## Reference

- Schema evolution in streaming Dataflow jobs and BigQuery tables, [part 1](https://robertsahlin.com/schema-evolution-in-streaming-dataflow-jobs-and-bigquery-tables-part-1/)
- Schema evolution in streaming Dataflow jobs and BigQuery tables, [part 2](https://robertsahlin.com/schema-evolution-in-streaming-dataflow-jobs-and-bigquery-tables-part-2/)
- Schema evolution in streaming Dataflow jobs and BigQuery tables, [part 3](https://robertsahlin.com/schema-evolution-in-streaming-dataflow-jobs-and-bigquery-tables-part-3/)
- DynamicMessage and FieldDescriptor https://go.dev/blog/protobuf-apiv2
