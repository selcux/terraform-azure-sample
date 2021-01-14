package client

import (
	"context"
	"fmt"

	"github.com/selcux/terraform-azure-sample/pkg/grpc/greet"
	greetgrpc "github.com/selcux/terraform-azure-sample/pkg/grpc/greet"
	"google.golang.org/grpc"
)

type Greet struct {
	conn   *grpc.ClientConn
	client greetgrpc.GreetServiceClient
}

func NewGreetClient() *Greet {
	return &Greet{}
}

func (g *Greet) Connect() error {
	target := fmt.Sprintf("%s:%d", "localhost", 9010)

	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return err
	}

	g.conn = conn
	g.client = greet.NewGreetServiceClient(conn)

	return nil
}

func (g *Greet) Close() error {
	return g.conn.Close()
}

func (g *Greet) SayHello(name string) (*greetgrpc.GreetResponse, error) {
	req := &greetgrpc.GreetRequest{Name: name}
	return g.client.SayHello(context.Background(), req)
}
