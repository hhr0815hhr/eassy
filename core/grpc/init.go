package grpc

import (
	g "google.golang.org/grpc"
)

var (
	GRPC *g.Server
)

func init() {
	GRPC = g.NewServer()
}
