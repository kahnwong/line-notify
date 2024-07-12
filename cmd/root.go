/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

type lineNotifyResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func notify(lineToken string, message string) {
	url := "https://notify-api.line.me/api/notify"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("message", message)
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", lineToken))

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// decode
	var response lineNotifyResponse
	if err := json.Unmarshal(body, &response); err != nil {
		slog.Error("Can not unmarshal JSON")
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
		// init env
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Loading env from env var instead...")
		}
		lineToken := os.Getenv("LINE_TOKEN")

		// main
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
