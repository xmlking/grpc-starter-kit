package protoutil

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
)

// AnyToMessage helper
func AnyToMessage(a *any.Any) (proto.Message, error) {
	var x ptypes.DynamicAny
	err := ptypes.UnmarshalAny(a, &x)
	return x.Message, err
}

// MessageToAny helper
func MessageToAny(pb proto.Message) (any *any.Any, err error) {
	if pb == nil {
		return
	}
	any, err = ptypes.MarshalAny(pb)
	if err != nil {
		return nil, err
	}
	return
}

// MessageToAny2 helper
func MessageToAny2(pb proto.Message) (*any.Any, error) {
	value, err := proto.Marshal(pb)
	if err != nil {
		return nil, err
	}
	// fullName := msg.(*transactionv1.TransactionEvent).ProtoReflect().Descriptor().FullName()
	return &any.Any{TypeUrl: "/" + proto.MessageName(pb), Value: value}, nil
}
