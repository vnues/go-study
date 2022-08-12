package main

import (
	"context"
	trippb "coolcar/proto/gen/go"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	//连接gateway
	con, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}

	tsClient := trippb.NewTripServiceClient(con)
	res, err := tsClient.GetTrip(context.Background(), &trippb.GetTripRequest{
		Id: "trips01",
	})
	if err != nil {
		log.Fatalf("未获取到trips: %v", err)
	}
	fmt.Println(res)
}
