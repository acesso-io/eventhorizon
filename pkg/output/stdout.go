package output

import (
	"context"

	"github.com/rs/zerolog/log"
)

type Stdout struct {
}

func (o *Stdout) Send(ctx context.Context, event interface{}) error {
	log.Debug().
		Interface("event", event).
		Msg("Stdout event")

	return nil
}

func (o *Stdout) Close() error {
	return nil
}

func NewStdout() (*Stdout, error) {
	o := Stdout{}

	return &o, nil
}
