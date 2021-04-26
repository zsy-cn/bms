package service

import (
	"context"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// LoraCredential 自定义认证
type LoraCredential struct{}

// GetRequestMetadata ...
func (c LoraCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": viper.GetString("appserver-token"),
	}, nil
}

// RequireTransportSecurity ...
func (c LoraCredential) RequireTransportSecurity() bool {
	return true
}

func getGRPCOpts() []grpc.DialOption {
	var opts []grpc.DialOption
	certPath := viper.GetString("appserver-cert")
	creds, err := credentials.NewClientTLSFromFile(certPath, "")
	if err != nil {
		panic(err)
	}
	opts = append(opts, grpc.WithTransportCredentials(creds))
	opts = append(opts, grpc.WithPerRPCCredentials(new(LoraCredential)))
	return opts
}
