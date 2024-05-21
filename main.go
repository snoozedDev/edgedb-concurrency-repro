package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/edgedb/edgedb-go"
)

var ctx = context.Background()

func main() {
	client, _ := edgedb.CreateClientDSN(ctx, "edgedb://edgedb:password@localhost/main", edgedb.Options{
		TLSOptions: edgedb.TLSOptions{
			SecurityMode: edgedb.TLSModeInsecure,
		},
		Concurrency: 4,
	})

	defer client.Close()

	startTime := time.Now()

	err := client.EnsureConnected(ctx)
	if err != nil {
		log.Print("Failed to connect to the database: ", err)
		return
	}

	connectedElapsed := time.Since(startTime)

	fmt.Printf("Connected in %v\n", connectedElapsed)

	time.Sleep(time.Second)

	var wg sync.WaitGroup
	wg.Add(8)

	for i := 0; i < 4; i++ {
		for i := 0; i < 2; i++ {
			go testDB(client, &wg)
		}
		time.Sleep(time.Second)
	}

	wg.Wait()
}

func testDB(client *edgedb.Client, wg *sync.WaitGroup) {
	defer wg.Done()

	startTime := time.Now()
	var result int64
	err := client.QuerySingle(ctx, "SELECT 1", &result)
	connectedElapsed := time.Since(startTime)
	if err != nil {
		log.Fatal(err)
	}

	if connectedElapsed > 1*time.Millisecond {
		fmt.Printf("Connected had to reconnect (%v)\n", connectedElapsed)
	} else {
		fmt.Printf("Client didn't have to reconnect (%v)\n", connectedElapsed)
	}
}
