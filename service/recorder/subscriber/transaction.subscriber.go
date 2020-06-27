package subscriber

import (
	"context"
	"fmt"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	transactionv1 "github.com/xmlking/grpc-starter-kit/mkit/service/transaction/v1"

	"github.com/xmlking/grpc-starter-kit/shared/constants"
)

type TransactionSubscriber struct {
}

// NewTransactionSubscriber returns an instance of `TransactionSubscriber`.
func NewTransactionSubscriber() *TransactionSubscriber {
	return &TransactionSubscriber{}
}

// Handle is a method to send emails
func (s *TransactionSubscriber) HandleSend(ctx context.Context, event cloudevents.Event) cloudevents.Result {
	log.Debug().Msgf("Event Context: %+v\n", event.Context)

	from := event.Source()
	tranId := event.ID()
	if len(tranId) == 0 {
		log.Error().Msg("TransactionSubscriber: missing  TranID")
		return errors.New("TransactionSubscriber: missing  TranID")
	}

	transaction := &transactionv1.TransactionEvent{}
	if err := event.DataAs(transaction); err != nil {
		log.Error().Err(err).Send()
		// return err
		return cloudevents.ResultNACK
	}

	switch from {
	case constants.ACCOUNT_SERVICE:
		log.Debug().Msgf("%s#%s, transaction: %#v", tranId, from, transaction)
	case constants.EMAILER_SERVICE:
		log.Debug().Msgf("%s#%s, transaction: %#v", tranId, from, transaction)
	case constants.GREETER_SERVICE:
		log.Debug().Msgf("%s#%s, transaction: %#v", tranId, from, transaction)
	case constants.ACCOUNT_CLIENT:
		log.Debug().Msgf("%s#%s, transaction: %#v", tranId, from, transaction)
	default:
		log.Error().Msgf("TransactionSubscriber: unknown  from: %s", from)
		return fmt.Errorf("TransactionSubscriber: unknown  from: %s", from)
	}
	return cloudevents.ResultACK
}
