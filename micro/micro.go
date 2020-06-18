package micro



import (
    "net"

    "github.com/pkg/errors"
    "github.com/soheilhy/cmux"
    "google.golang.org/grpc/grpclog"
)

type MicroServer struct {
    mux cmux.CMux
    lis net.Listener
}
type Interface interface {
    Serve(l net.Listener) error
    Shutdown()
}

func NewService(mux cmux.CMux, lis net.Listener) Interface {
    return &MicroServer{
        mux: mux,
        lis: lis,
    }
}
// Serve implements Server.Serve
func (s *MicroServer) Serve(net.Listener) error {
    grpclog.Info("mux is starting %s", s.lis.Addr())

    err := s.mux.Serve()

    grpclog.Infof("mux is closed: %v", err)

    return errors.Wrap(err, "failed to serve cmux server")
}

// Shutdown implements Server.Shutdown
func (s *MicroServer) Shutdown() {
    err := s.lis.Close()
    if err != nil {
        grpclog.Errorf("failed to close cmux's listener: %v", err)
    }
}
