package subscriber

import (
    "context"

    //pscontext "github.com/cloudevents/sdk-go/protocol/pubsub/v2/context"
    cloudevents "github.com/cloudevents/sdk-go/v2"
    "github.com/cloudevents/sdk-go/v2/protocol"
    "github.com/rs/zerolog/log"

    emailerPB "github.com/xmlking/grpc-starter-kit/service/emailer/proto/emailer"
    "github.com/xmlking/grpc-starter-kit/service/emailer/service"
)

// EmailSubscriber is Subscriber
type EmailSubscriber struct {
    emailService service.EmailService
}

// NewEmailSubscriber returns an instance of `EmailSubscriber`.
func NewEmailSubscriber(emailService service.EmailService) *EmailSubscriber {
    return &EmailSubscriber{
        emailService: emailService,
    }
}

// Handle is a method to send emails
func (s *EmailSubscriber) HandleSend(ctx context.Context, event cloudevents.Event) error {
    log.Debug().Msgf("Event Context: %+v\n", event.Context)
    log.Debug().Msgf("Event Source from Context: %+v\n",event.Context.AsV1().Source)
    //log.Debug().Msgf("Transport Context: %+v\n", pscontext.ProtocolContextFrom(ctx))

    // validate event conforms to cloudevents specification
    if err := event.Validate(); err != nil {
        log.Error().Err(err).Send()
        return err
    }

    data := &emailerPB.Message{}
    if err := event.DataAs(data); err != nil {
        log.Error().Err(err).Send()
        return err
    }

    return s.emailService.Welcome(data.Subject, data.To)
    // return cloudevents.ResultACK
}

func (s *EmailSubscriber) HandleRequest(ctx context.Context, event cloudevents.Event) (*cloudevents.Event,  cloudevents.Result) {
    log.Debug().Msgf("Event Context: %+v\n", event.Context)
    log.Debug().Msgf("Event Source from Context: %+v\n",event.Context.AsV1().Source)

    data := &emailerPB.Message{}
    if err := event.DataAs(data); err != nil {
        log.Error().Err(err).Send()
        return nil, err
    }

    responseEvent := cloudevents.NewEvent()
    responseEvent.SetSource("/mod3")
    responseEvent.SetType("samples.http.mod3")
    _ = responseEvent.SetData(cloudevents.ApplicationJSON, &emailerPB.Message{Subject: "echo" + data.Subject})
    return &responseEvent, protocol.ResultACK
    // return &responseEvent, nil
}
