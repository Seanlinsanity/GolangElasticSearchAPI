package elasticSearch

import (
	"context"
	"time"

	"github.com/olivere/elastic"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(*elastic.Client)
	Index(interface{}) (*elastic.IndexResponse, error)
}

type esClient struct {
	client *elastic.Client
}

func Init() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		// elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	)

	if err != nil {
		panic(err)
	}

	Client.setClient(client)
}

func (c *esClient) setClient(cleint *elastic.Client) {
	c.client = cleint
}

func (esClient *esClient) Index(interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	return esClient.client.Index().Do(ctx)
}