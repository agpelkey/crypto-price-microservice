package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/agpelkey/microserv-prac/client"
)

func main() {
	client := client.New("http://localhost:3000/")

	price, err := client.FetchPrice(context.Background(), "ET")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", price)
	listenAddr := flag.String("listenaddr", ":3000", "listen address the service is running on")
	flag.Parse()

	svc := NewLoggingService(NewMetricService(&priceFetcher{}))

	server := NewJSONAPIServer(*listenAddr, svc)
	server.Run()
}
