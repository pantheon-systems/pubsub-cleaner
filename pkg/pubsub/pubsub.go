package pscleaner

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/cloud"
	"google.golang.org/cloud/pubsub"
)

// Cleaner is our local struct for holding config, context and pubsub client to GCE
type Cleaner struct {
	ctx    context.Context
	client *pubsub.Client
	Config Config
}

// Config is what you pass to the various cleaners
type Config struct {
	NoOP    bool
	Keep    string
	Topic   string
	Project string
}

func cloudContext(projectID string) (context.Context, error) {
	ctx := context.Background()
	httpClient, err := google.DefaultClient(ctx, pubsub.ScopePubSub)
	if err != nil {
		return nil, err
	}
	return cloud.WithContext(ctx, projectID, httpClient), nil
}

// NewCleaner creates a publisher using GCE PubSub backend
func NewCleaner(config Config) (Cleaner, error) {
	pubSubCtx, err := cloudContext(config.Project)
	fatalIf(err)

	client, err := pubsub.NewClient(pubSubCtx, config.Project)
	fatalIf(err)

	c := Cleaner{
		ctx:    pubSubCtx,
		client: client,
		Config: config,
	}

	return c, nil
}

func fatalIf(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// CleanTopicSubscriptions will clean subscrbers per the CleanerConfig
func (c *Cleaner) CleanTopicSubscriptions() error {
	topic := c.client.Topic(c.Config.Topic)
	subs := topic.Subscriptions(c.ctx)

	fmt.Println("Subscriptions for Topic: ", c.Config.Topic)
	for {
		sub, err := subs.Next()
		if err != nil {
			if err.Error() == "no more messages" {
				break
			} else {
				fmt.Println("Error getting subscription:", err.Error())
				return err
			}
		}

		fmt.Println("Sub: ", sub.Name())
		if !strings.Contains(sub.Name(), c.Config.Keep) {
			fmt.Println("\tDeleting subscription doesn't have keep string")
			if !c.Config.NoOP {
				fatalIf(sub.Delete(c.ctx))
			}
		} else {
			fmt.Println("\tKeeping matches: ", c.Config.Keep)
		}
	}

	return nil
}
