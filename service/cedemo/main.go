package main

import (
    "context"
    "log"

    cloudevents "github.com/cloudevents/sdk-go/v2"

    "github.com/xmlking/grpc-starter-kit/service/cedemo/subscriber"
)


func main() {

    client, err := cloudevents.NewDefaultClient()
    if err != nil {
        log.Fatal(err.Error())
    }

    r := subscriber.Receiver{Client: client, Target: "http://localhost:8081"}

    // Depending on whether targeting data has been supplied,
    // we will either reply with our response or send it on to
    // an event sink.
    var receiver interface{} // the SDK reflects on the signature.
    if r.Target == "" {
        receiver = r.ReceiveAndReply
    } else {
        receiver = r.ReceiveAndSend
    }

    if err := client.StartReceiver(context.Background(), receiver); err != nil {
        log.Fatal(err)
    }
}
