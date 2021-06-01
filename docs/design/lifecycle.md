# Server Lifecycle

KMG(Kill Me Gently) is a library that aids in graceful shutdown of a process/application.

## Requirements

- [ ] Graceful Shutdown
- [ ] Shutdown Timeout
- [ ] workers listen for the cancellation event and cancel any in-progress translations/ NACK PubSub messages and release any resources. 
## References

- Graceful Shutdowns in Golang with [signal.NotifyContext](https://millhouse.dev/posts/graceful-shutdowns-in-golang-with-signal-notify-context)
- [How to Manage Database Timeouts and Cancellations in Go](https://www.alexedwards.net/blog/how-to-manage-database-timeouts-and-cancellations-in-go)
- [Graceful shutdown with Go http servers and Kubernetes rolling updates](https://medium.com/over-engineering/graceful-shutdown-with-go-http-servers-and-kubernetes-rolling-updates-6697e7db17cf)
- [Example of graceful shutdown with grpc healthserver * httpserver](https://gist.github.com/akhenakh/38dbfea70dc36964e23acc19777f3869)
- https://www.youtube.com/watch?v=LSzR0VEraWw
- https://www.sohamkamani.com/golang/2018-06-17-golang-using-context-cancellation/
- https://millhouse.dev/posts/graceful-shutdowns-in-golang-with-signal-notify-context
- [livenessProbe for gRPC](https://codeburst.io/kubernetes-grpc-services-and-probes-by-example-1cb611da45ab)
