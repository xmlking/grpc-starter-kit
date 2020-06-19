package micro


// https://github.com/stevepartridge/service

import (
    "net"

    "github.com/pkg/errors"
    "github.com/soheilhy/cmux"
    "google.golang.org/grpc"
    "google.golang.org/grpc/grpclog"
)

type MicroService struct {
    mux cmux.CMux
    lis net.Listener

    MaxReceiveSize int
    MaxSendSize    int

    DialOption      []grpc.DialOption
    ServerOptions      []grpc.ServerOption
    UnaryInterceptors  []grpc.UnaryServerInterceptor
    StreamInterceptors []grpc.StreamServerInterceptor

    Server *grpc.Server
}
type Interface interface {
    Serve(l net.Listener) error
    Shutdown()
}

func NewService(mux cmux.CMux, lis net.Listener) Interface {
    return &MicroService{
        mux: mux,
        lis: lis,
    }
}
// Serve implements Server.Serve
func (s *MicroService) Serve(net.Listener) error {
    grpclog.Info("mux is starting %s", s.lis.Addr())

    //options := []grpc.ServerOption{
    //    grpc.MaxRecvMsgSize(s.MaxReceiveSize),
    //    grpc.MaxSendMsgSize(s.MaxSendSize),
    //}

    err := s.mux.Serve()

    grpclog.Infof("mux is closed: %v", err)

    return errors.Wrap(err, "failed to serve cmux server")
}

// Shutdown implements Server.Shutdown
func (s *MicroService) Shutdown() {
    err := s.lis.Close()
    if err != nil {
        grpclog.Errorf("failed to close cmux's listener: %v", err)
    }
}
