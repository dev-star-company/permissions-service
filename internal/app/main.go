package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"permissions-service/internal/app/ent"
	"permissions-service/internal/config/env"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func New(port int) {
	connString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		env.PG_HOST, env.PG_PORT, env.PG_USER, env.PG_DBNAME, env.PG_PASSWORD, env.PG_SSLMODE)

	client, err := ent.Open("postgres", connString)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	opts = append(opts,
		grpc.ChainUnaryInterceptor(),
		grpc.ChainStreamInterceptor(),
	)
	grpcServer := grpc.NewServer(opts...)

	RegisterControllers(grpcServer, client)
	reflection.Register(grpcServer)

	fmt.Printf("Server is running on port:%d\n", port)
	grpcServer.Serve(lis)
}
