package translog

import (
	"context"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/rs/zerolog/log"
	broker "github.com/xmlking/toolkit/broker/cloudevents"
	"google.golang.org/grpc"

	"github.com/xmlking/grpc-starter-kit/internal/constants"
	transactionv1 "github.com/xmlking/grpc-starter-kit/mkit/service/transaction/v1"
)

func publish(ctx context.Context, publisher broker.Publisher, source string, req, rsp proto.Message) (err error) {
	var anyReq, anyRsp *any.Any
	var event cloudevents.Event
	anyReq, err = ptypes.MarshalAny(req)
	if err != nil {
		goto End
	}
	anyRsp, err = ptypes.MarshalAny(rsp)
	if err != nil {
		goto End
	}

	// Create an Event.
	event = cloudevents.NewEvent()
	event.SetSource(source)
	event.SetType("translog.transaction.event")
	if traceId := metautils.ExtractIncoming(ctx).Get(constants.TraceIDKey); traceId != "" {
		event.SetID(traceId)
	}

	// err = event.SetData(cloudevents.ApplicationJSON, &transactionv1.TransactionEvent{Req:&structpb.Value{}, Rsp: &structpb.Value{}})
	err = event.SetData(cloudevents.ApplicationJSON, &transactionv1.TransactionEvent{Req: anyReq, Rsp: anyRsp})
	if err != nil {
		goto End
	}

	if result := publisher.Publish(ctx, event); cloudevents.IsUndelivered(result) {
		log.Error().Err(result).Str("component", "translog").Msg("Publish: Failed to send")
		err = result
	} else if cloudevents.IsNACK(result) {
		log.Error().Err(result).Str("component", "translog").Msg("Publish: event not accepted")
		err = result
	}

End:
	if err != nil {
		log.Error().Err(err).Msg("Publisher: Failed publishing translation")
	}
	return
}

// UnaryServerInterceptor publish request and response to cloudevents
func UnaryServerInterceptor(publisher broker.Publisher, source string) grpc.UnaryServerInterceptor {

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)

		// we already logged error in Publish. lets ignore error here. # Note: this is blocking call..
		_ = publish(ctx, publisher, source, req.(proto.Message), resp.(proto.Message))
		// go publish(ctx, p, req.Body().(proto.Message), rsp.(proto.Message))
		return resp, err
	}
}
