package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/jimmysawczuk/try"
	"github.com/olivere/elastic"
	"github.com/pkg/errors"
)

var timeout = 60 * time.Second
var interval = 1 * time.Second

func init() {

	flag.DurationVar(&timeout, "timeout", 60*time.Second, "total amount of time to try")
	flag.DurationVar(&interval, "interval", 1*time.Second, "amount of time to wait between tries")

	flag.Parse()
}

func main() {
	if len(flag.Args()) < 1 {
		log.Fatalf("missing required argument: elasticsearch connection string")
	}

	connectionString := flag.Arg(0)
	start := time.Now()

	log.Printf("attempting to connect to elasticsearch (will try for %s, %s between attempts)", timeout, interval)

	if err := try.Try(connectToElasticsearch(connectionString), timeout, interval); err != nil {
		log.Fatal(err)
	}

	log.Printf("connected in %s", time.Now().Sub(start).Truncate(time.Millisecond))

}

func connectToElasticsearch(connectionString string) func() error {
	return func() error {
		c, err := elastic.NewClient(
			elastic.SetURL(connectionString),
			elastic.SetHealthcheck(false),
			elastic.SetSniff(false),
		)

		if err != nil {
			return errors.Wrap(err, "failed to create elasticsearch client")
		}

		// Ping the Elasticsearch server to get e.g. the version number
		_, _, err = c.Ping(connectionString).Do(context.Background())
		if err != nil {
			// Handle error
			return errors.Wrap(err, "connect to elasticSearch")
		}

		return nil
	}
}
