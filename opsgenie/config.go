package opsgenie

import (
	"log"
  "time"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/sirupsen/logrus"
)

type OpsgenieClient struct {
	client *client.OpsGenieClient
}

type Config struct {
	ApiKey string
	ApiUrl string
}

func (c *Config) Client() (*OpsgenieClient, error) {
	myLogger := logrus.New()
	myLogger.SetLevel(logrus.WarnLevel)
	myLogger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat:  time.RFC3339Nano,
			PrettyPrint:      true,
	},)

	config := &client.Config{
		ApiKey:         c.ApiKey,
		RetryCount:     10,
		OpsGenieAPIURL: client.ApiUrl(c.ApiUrl),
		Backoff:        retryablehttp.DefaultBackoff,
		Logger: myLogger,
	}
	ogCli, err := client.NewOpsGenieClient(config)
	if err != nil {
		return nil, err
	}
	ogClient := OpsgenieClient{}
	ogClient.client = ogCli
	log.Printf("[INFO] OpsGenie client configured")
	return &ogClient, nil
}
