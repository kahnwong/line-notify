/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/carlmjohnson/requests"
	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
)

func notify(message string) {
	var lineChannelAccessToken = os.Getenv("LINE_CHANNEL_ACCESS_TOKEN")

	bodyBytes := []byte(fmt.Sprintf("{\n    \"messages\": [\n        {\n            \"type\": \"text\",\n            \"text\": \"%s\"\n        }\n    ]\n}", message))

	err := requests.
		URL("https://api.line.me/v2/bot/message/narrowcast").
		Method(http.MethodPost).
		BodyBytes(bodyBytes).
		Header("Authorization", fmt.Sprintf("Bearer %s", lineChannelAccessToken)).
		Header("Content-Type", "application/json").
		Fetch(context.Background())

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to notify")
	}
}

var notifyCmd = &cobra.Command{
	Use: "notify",
	Run: func(cmd *cobra.Command, args []string) {
		notify(args[0])
	},
}

func init() {
	rootCmd.AddCommand(notifyCmd)
}
