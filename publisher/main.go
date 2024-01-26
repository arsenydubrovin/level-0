package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	stan "github.com/nats-io/stan.go"
)

func main() {
	// Reading config
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	subject := os.Getenv("STAN_SUBJECT")
	natsURL := fmt.Sprintf("nats://%s:%s", os.Getenv("NATS_HOST"), os.Getenv("NATS_PORT"))
	clusterID := os.Getenv("STAN_CLUSTER_ID")

	// Connecting to nats-streaming
	sc, err := stan.Connect(clusterID, "nats-publisher", stan.NatsURL(natsURL))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	log.Println("Publisher started...")

	// Sending messages

	for {
		// Choose: publisher/invalid-orders or publisher/valid-orders
		file, err := os.Open("publisher/valid-orders")
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			time.Sleep(time.Second * 5)
			msg := scanner.Text()

			err = sc.Publish(subject, []byte(msg))
			log.Println("Sending message...")

			if err != nil {
				log.Fatal(err)
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		file.Close()
	}
}
