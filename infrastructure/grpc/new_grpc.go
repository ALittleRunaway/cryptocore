package grpc

import (
	"cryptocore/config"
	"cryptocore/entrypoint"
	pb "cryptocore/infrastructure/grpc/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

const serviceName = "grpc"

func NewGrpcConnection(grpcCfg *config.GrpcConfig, logger *zap.SugaredLogger, fatalErrCh chan error) error {

	serviceLogger := logger.Named(serviceName)

	listener, err := net.Listen(grpcCfg.Network, grpcCfg.Addr)
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterCryptoCoreServer(grpcServer, &entrypoint.GrpcServer{})

	serviceLogger.Info("Establishing the gRPC connection...")
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			fatalErrCh <- err
		}
	}()
	serviceLogger.Info("Established the gRPC connection successfully.")

	return nil
}
