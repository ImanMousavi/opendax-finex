package main

import (
	"context"
	"log"

	"opendax-clean/finex/api/proto/apipb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {
	// Open a connection to the server.
	conn, err := grpc.Dial(":5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed connecting to server: %s", err)
	}
	defer conn.Close()

	// Create a Asset service Client on the connection.
	client := apipb.NewAssetServiceClient(conn)

	// Get Asset.
	get, err := client.Get(context.Background(), &apipb.AssetGetRequest{
		Id: "WBTC",
	})
	if err != nil {
		se, _ := status.FromError(err)
		log.Fatalf("failed retrieving asset: status=%s message=%s", se.Code(), se.Message())
	}
	log.Printf("retrieved asset: %v", get)

	// List assets.
	list, err := client.List(context.Background(), &apipb.AssetListRequest{})
	if err != nil {
		se, _ := status.FromError(err)
		log.Fatalf("failed retrieving assets: status=%s message=%s", se.Code(), se.Message())
	}
	log.Printf("retrieved assets: %v", list)

	// List assets by ids.
	list, err = client.List(context.Background(), &apipb.AssetListRequest{Ids: []string{"ETH", "USDT", "unknown"}})
	if err != nil {
		se, _ := status.FromError(err)
		log.Fatalf("failed retrieving assets: status=%s message=%s", se.Code(), se.Message())
	}
	log.Printf("retrieved assets: %v", list)
}
