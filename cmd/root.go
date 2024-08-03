/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/carlmjohnson/requests"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cobra"
)

var lineToken = os.Getenv("LINE_TOKEN")

type lineNotifyResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func notify(lineToken string, message string) {
	body := url.Values{"message": {message}}

	var response lineNotifyResponse
	err := requests.
		URL("https://notify-api.line.me").
		Method(http.MethodPost).
		Path("api/notify").
		BodyForm(body).
		Header("Authorization", fmt.Sprintf("Bearer %s", lineToken)).
		Header("Content-Type", "application/x-www-form-urlencoded").
		ToJSON(&response).
		Fetch(context.Background())

	if err != nil {
		fmt.Println(err)
	}

	if response.Status == 200 {
		fmt.Println("Successfully notified line")
	} else {
		fmt.Println(response.Message)
	}
}

var rootCmd = &cobra.Command{
	Use: "line-notify",
	Run: func(cmd *cobra.Command, args []string) {
		notify(lineToken, args[0])
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
