package main

import (
	"context"

	cloudevents "github.com/cloudevents/sdk-go/v2"
    "github.com/rs/zerolog/log"

    "github.com/xmlking/grpc-starter-kit/service/cedemo/subscriber"
    "github.com/xmlking/grpc-starter-kit/shared/config"
)

func main() {

	ceClient, err := cloudevents.NewDefaultClient()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	r := subscriber.Receiver{Client: ceClient, Target: "http://localhost:8081"}

	// Depending on whether targeting data has been supplied,
	// we will either reply with our response or send it on to
	// an event sink.
	var receiver interface{} // the SDK reflects on the signature.
	if r.Target == "" {
		receiver = r.ReceiveAndReply
	} else {
		receiver = r.ReceiveAndSend
	}

    // Start server!
    println(config.GetBuildInfo())
    log.Info().Msgf("Server (%s) started at: %s, secure: %t", "cedemo", "http://localhost:8080", false)
	if err := ceClient.StartReceiver(context.Background(), receiver); err != nil {
		log.Fatal().Err(err).Send()
	}
}
