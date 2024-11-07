package main

import (
	"context"
	"github.com/ascenmmo/websocket-server/env"
	"github.com/ascenmmo/websocket-server/pkg/start"
	"github.com/rs/zerolog"
	"os"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	ctx := context.Background()

	err := start.StartWebSocket(
		ctx,
		env.ServerAddress,
		env.TCPPort,
		env.WebsocketPort,
		env.TokenKey,
		env.MaxRequestPerSecond,
		10,
		logger,
		true,
	)

	if err != nil {
		logger.Fatal().Err(err).Msg("failed to start websocket server")
	}
}
