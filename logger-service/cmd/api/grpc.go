package main

import (
	"context"
	"fmt"
	"log"
	"logger-service/data"
	"logger-service/logs"
	"net"

	"google.golang.org/grpc"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
	Models data.Model
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error)  {
	input :=req.GetLogEntry()

	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.Insert(logEntry)

	if err!=nil{
		return nil, err
	}

	res := &logs.LogResponse{
		Result: "Log inserted",
	}

	return res, nil
}


func (app *Config) gRPCListen()  {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err!=nil {
		log.Fatalf("Failed to listen for gRpc: %v", err)
	}

	s := grpc.NewServer()

	logs.RegisterLogServiceServer(s, &LogServer{Models: app.Models})

	log.Printf("gRPC server started successfully on port %s", gRpcPort)

	err = s.Serve(lis)
	if err!=nil {
		log.Fatalf("Failed to listen for gRpc: %v", err)
	}
}