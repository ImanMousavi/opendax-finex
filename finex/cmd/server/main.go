package main

import (
	"context"
	"log"
	"net"

	"opendax-clean/finex/api"
	"opendax-clean/finex/api/proto/apipb"
	"opendax-clean/finex/ent"

	"ariga.io/entcache"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

func main() {
	// Open the database connection.
	db, err := sql.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	// Decorates the sql.Driver with entcache.Driver.
	drv := entcache.NewDriver(db)
	// Create an ent.Client.
	client := ent.NewClient(ent.Driver(drv))
	defer client.Close()

	// Run the migration tool (creating tables, etc).
	// Tell the entcache.Driver to skip the caching layer
	// when running the schema migration.
	if err := client.Schema.Create(entcache.Skip(context.Background())); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Create some assets.
	client.Asset.CreateBulk(
		client.Asset.Create().SetID("ETH").SetName("Ethereum").SetIndex(0),
		client.Asset.Create().SetID("USDT").SetName("Tether").SetIndex(1),
		client.Asset.Create().SetID("USDC").SetName("USD Coin").SetIndex(2),
		client.Asset.Create().SetID("WBTC").SetName("Wrapped Bitcoin").SetIndex(3),
		client.Asset.Create().SetID("BUSD").SetName("Binance USD").SetIndex(4),
	).SaveX(context.Background())

	// Create a new gRPC server (you can wire multiple services to a single server).
	server := grpc.NewServer()

	// Register the Asset service with the server.
	apipb.RegisterAssetServiceServer(server, api.NewAssetService(client))

	// Open port 5000 for listening to traffic.
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed listening: %s", err)
	}

	// Listen for traffic indefinitely.
	if err := server.Serve(lis); err != nil {
		log.Fatalf("server ended: %s", err)
	}
}
