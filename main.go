package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"stock/driver"
	"stock/proto/proto"
	"stock/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	driver.InitGorm()
	driver.Register("stock_srv-1", "stock_srv-1", "localhost:8001", []string{"stock_srv-1"})
	lis, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())
	proto.RegisterStockServiceServer(grpcServer, service.NewStockService())
	go func() {
		err := grpcServer.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Kill, os.Interrupt)
	<-quit
	log.Println("shutdown grpc server...")
	log.Println("deregister consul...")
	driver.DeRegister("stock_srv-1")
}
