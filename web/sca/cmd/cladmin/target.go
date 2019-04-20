package main

import (
	"context"

	"github.com/IBM/cloudland/web/sca/clients"
	"github.com/IBM/cloudland/web/sca/targets"
	"google.golang.org/grpc"
)

func init() {
	RootCmd.AddCommand(targets.Commands(context.Background,
		func() *grpc.ClientConn {
			return clients.GetClientConn("admin")
		})...)
}
