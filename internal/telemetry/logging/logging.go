package logging

// https://cloud.google.com/trace/docs/setup/go-ot#gke
// https://medium.com/google-cloud/integrating-tracing-and-logging-with-opentelemetry-and-stackdriver-a5396fbc3e78

// Adding new snap https://github.com/open-telemetry/opentelemetry-go
//import (
//    "cloud.google.com/go/logging"
//)
//
//func EnableLogging(projectID string, sampling float64) {
//    ctx := context.Background()
//    var err error
//    loggingClient, err = logging.NewClient(ctx, projectID)
//    if err != nil {
//        fmt.Printf("Failed to create logging client: %v", err)
//        return
//    }
//    fmt.Printf("Stackdriver Logging initialized with project id %s, see Cloud "+
//        " Console under GCE VM instance > all instance_id\n", projectID)
//}
