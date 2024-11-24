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

func notifyTest(message string) {
	var lineChannelAccessToken = os.Getenv("LINE_CHANNEL_ACCESS_TOKEN")
	var linePersonalUserID = os.Getenv("LINE_PERSONAL_USER_ID")

	bodyBytes := []byte(fmt.Sprintf("{ \"to\": \"%s\",\n    \"messages\": [\n        {\n            \"type\": \"text\",\n            \"text\": \"%s\"\n        }\n    ]\n}", linePersonalUserID, message))

	err := requests.
		URL("https://api.line.me/v2/bot/message/push").
		Method(http.MethodPost).
		BodyBytes(bodyBytes).
		Header("Authorization", fmt.Sprintf("Bearer %s", lineChannelAccessToken)).
		Header("Content-Type", "application/json").
		Fetch(context.Background())

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to notify")
	}
}

var testCmd = &cobra.Command{
	Use: "test",
	Run: func(cmd *cobra.Command, args []string) {
		notifyTest(args[0])
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
