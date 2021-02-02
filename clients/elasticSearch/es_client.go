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
	Index(string, interface{}) (*elastic.IndexResponse, error)
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

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

func (esClient *esClient) Index(index string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := esClient.client.Index().Index(index).BodyJson(doc).Do(ctx)
	if err != nil {
		println("error when trying to index document in es", err)
		return nil, err
	}
	return result, nil
}
