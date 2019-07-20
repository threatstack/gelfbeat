package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/threatstack/gelfbeat/config"
)

// Gelfbeat configuration.
type Gelfbeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// New creates an instance of gelfbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Gelfbeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts gelfbeat.
func (bt *Gelfbeat) Run(b *beat.Beat) error {
	logp.Info("gelfbeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	var packet = make(chan string)
	go listen(packet, bt.config)

	for {
		var jsonMessage string
		select {
		case <-bt.done:
			return nil
		case jsonMessage = <-packet:
		}

		if err != nil {
			return err
		}

		event := beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"message": jsonMessage,
			},
		}
		bt.client.Publish(event)
	}

}

// Stop stops gelfbeat.
func (bt *Gelfbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}

