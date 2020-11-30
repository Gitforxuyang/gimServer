package main

import (
	"fmt"
	"gimServer/app/service"
	"gimServer/conf"
	"gimServer/domain/repo"
	"gimServer/handler"
	"gimServer/infra/mongo"
	"gimServer/infra/redis"
	"gimServer/infra/wrapper"
	"gimServer/proto"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			wrapper.NewServerWrapper(),
		)),
	)
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05.999"})
	//资源初始化
	config := conf.InitConfig()
	redis.InitClient(config)
	mongoClient := mongo.InitMongo(config)
	repo := repo.NewDomainRepo(mongoClient)
	svc := service.NewService(repo)
	im.RegisterImServer(grpcServer, handler.NewHandler(svc))
	log.Printf("server run")
	grpcServer.Serve(lis)

}
