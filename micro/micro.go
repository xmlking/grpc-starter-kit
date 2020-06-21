package micro

// https://github.com/stevepartridge/service

import (
	"net"

	"github.com/pkg/errors"
	"github.com/sarulabs/di/v2"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type Option func(*Options)

type Service interface {
	// The service name
	Name() string
	// Init initialises options
	Init(...Option)
	// Options returns the current options
	Options() Options
	// Client is used to call services
	Client() *grpc.ClientConn
	// Server is for handling requests and events
	Server() *grpc.Server
	// Run the service
	Run() error
	// The service implementation
	String() string
}

type MicroService struct {
	mux       cmux.CMux
	lis       net.Listener
	config    Interface
	container di.Container

	MaxReceiveSize int
	MaxSendSize    int

	DialOption         []grpc.DialOption
	ServerOptions      []grpc.ServerOption
	UnaryInterceptors  []grpc.UnaryServerInterceptor
	StreamInterceptors []grpc.StreamServerInterceptor

	Server *grpc.Server
	// ceClient
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
	s.Server.GracefulStop()
}
